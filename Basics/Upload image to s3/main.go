package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"time"
	"image"
	_ "image"
	"image/png"
	_ "image/png"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type JSONStruct struct {
	Img string `json:"Image"`
}

type Response struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

const (
	S3_REGION = "" //Your Bucket region
	S3_BUCKET = "" //Your S3 Bucket name
)

func getPicture(context echo.Context) JSONStruct {

	var jsonstruct JSONStruct

	errDecode := json.NewDecoder(context.Request().Body).Decode(&jsonstruct)
	if errDecode != nil {
		log.Println("error in decoding image", errDecode)
	}
	return jsonstruct
}

func PostPicture(context echo.Context) error {

	var response Response
	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		log.Fatal(err)
		response.Code = 400
		response.Type = http.StatusText(response.Code)
		response.Message = "AWS Session error"
		return json.NewEncoder(context.Response()).Encode(response)
	}

	var jsonStruct JSONStruct

	jsonStruct = getPicture(context)

	buff, errPng := Base64ToPNG(jsonStruct.Img)
	if errPng != nil {
		log.Println("Base64 to PNG Conversion error: ", errPng)
		response.Code = 400
		response.Type = http.StatusText(response.Code)
		response.Message = "Unable to convert Base64 image to PNG"
		return json.NewEncoder(context.Response()).Encode(response)
	}

	// Upload
	err = AddFileToS3(s, buff)
	if err != nil {
		log.Fatal(err)
		response.Code = 400
		response.Type = http.StatusText(response.Code)
		response.Message = "Image cannot be uploaded"
		return json.NewEncoder(context.Response()).Encode(response)
	}

	response.Code = 200
	response.Type = http.StatusText(response.Code)
	response.Message = "Success"
	context.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(context.Response()).Encode(response)
}

// Base64ToPNG converts pre-existing base64 data to png image
func Base64ToPNG(base64Image string) (*bytes.Buffer, error) {

	buff := new(bytes.Buffer)
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Image))
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Println(err)
		return buff, err
	}
	// bounds := m.Bounds()

	// log.Println(bounds, formatString)
	if err = png.Encode(buff, img); err != nil {
		log.Println(err)
		return buff, err
	}
	return buff, err
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, image *bytes.Buffer) error {

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	
	timeDet := time.Now()
	timeStr := timeDet.String()
	KeyName := "POC"+ timeStr
	cfg := &aws.Config{
		Region: aws.String(S3_REGION)}

	bodyDtls := bytes.NewReader(image.Bytes())
	fileType := http.DetectContentType(image.Bytes())

	params := &s3.PutObjectInput{
		Bucket:      aws.String(S3_BUCKET),
		Key:         aws.String(KeyName),
		Body:        bodyDtls,
		ContentType: aws.String(fileType),
	}
	svc := s3.New(session.New(), cfg)
	resp, err := svc.PutObject(params)

	log.Println("response ", awsutil.StringValue(resp))
	return err
}

func main() {

	e := echo.New()
	e.POST("/postpicturetos3", PostPicture)
	e.Logger.Fatal(e.Start(":7000"))
}
