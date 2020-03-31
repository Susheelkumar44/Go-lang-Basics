package main

import (
	"fmt"
	"os"
)

type song struct {
	name, artist string
	next         *song
}

type playlist struct {
	name       string
	head       *song
	nowPlaying *song
}

func createPlaylist(name string) *playlist {
	return &playlist{
		name: name,
	}
}

func (p *playlist) addSong(name, artist string) error {
	s := &song{
		name:   name,
		artist: artist,
	}

	if p.head == nil {
		p.head = s
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = s
	}
	return nil
}

func (p *playlist) showAllSongs() error {
	currentNode := p.head
	if p.head == nil {
		fmt.Println("Playlist is empty")
	} else {
		fmt.Println("List of all songs in playlist are: ")
		fmt.Println(currentNode.name, currentNode.artist)
		for currentNode.next != nil {
			currentNode = currentNode.next
			fmt.Println(currentNode.name, currentNode.artist)
		}
	}
	return nil
}

func (p *playlist) startPlaying() *song {
	p.nowPlaying = p.head
	if p.head == nil {
		return nil
	}
	return p.nowPlaying
}

func (p *playlist) nextSong() *song {
	if p.nowPlaying.next == nil {
		return nil
	} 
	p.nowPlaying = p.nowPlaying.next
	return p.nowPlaying
}

func main() {

	var i int
	playlistName := "myPlayList"

	MyPlayList := createPlaylist(playlistName)
	fmt.Println("Created Playlist")

	for {
		fmt.Println("\nEnter 1 to add songs to playlist")
		fmt.Println("Enter 2 to show all songs in playlist")
		fmt.Println("Enter 3 to Start Playing songs")
		fmt.Println("Enter 4 to play next song")
		fmt.Println("Enter 0 to exit!")

		fmt.Scanf("%d", &i)

		switch i {
		case 0:
			os.Exit(3)
		case 1:
			fmt.Println("Adding songs into playlist")
			MyPlayList.addSong("name1", "artist1")
			MyPlayList.addSong("name2", "artist2")
			MyPlayList.addSong("name3", "artist3")
			MyPlayList.addSong("name4", "artist1")
			MyPlayList.addSong("name5", "artist4")
			MyPlayList.addSong("name6", "artist2")
		case 2:
			fmt.Println("Showing All songs in Playlist.....")
			MyPlayList.showAllSongs()
		case 3:
			MyPlayList.startPlaying()
			if MyPlayList.startPlaying() == nil {
				fmt.Println("No songs in playlist")
			} else {
				fmt.Printf("Now playing: %s by %s ", MyPlayList.nowPlaying.name, MyPlayList.nowPlaying.artist)
				fmt.Println()
			}
		case 4:
			if MyPlayList.nextSong() == nil {
				fmt.Println("No songs in playlist")
			} else {
				MyPlayList.nextSong()
				fmt.Println("Changing to next song...")
				fmt.Printf("Now playing: %s by %s ", MyPlayList.nowPlaying.name, MyPlayList.nowPlaying.artist)
				fmt.Println()
			}
		}
	}
}
