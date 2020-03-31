package main

import "fmt"

type song struct {
	name, artist   string
	next, previous *song
}

type playlist struct {
	name       string
	head       *song
	tail       *song
	nowPlaying *song
}

func createPlayList(name string) *playlist {
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
		p.tail.next = s
		s.previous = p.tail
	}
	p.tail = s
	return nil
}

func (p *playlist) ShowAllSongs() error {
	currentNode := p.head
	if p.head == nil {
		fmt.Println("Playlist is empty")
	} else {
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
	return p.nowPlaying
}

func (p *playlist) playNextSong() *song {
	p.nowPlaying = p.nowPlaying.next
	return p.nowPlaying
}

func (p *playlist) playPreviousSong() *song {
	p.nowPlaying = p.nowPlaying.previous
	return p.nowPlaying
}

func (p *playlist) reversePlayList() {
	currentNode := p.head
	var previousNode *song
	p.tail = p.head

	for currentNode != nil {
		next := currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = next
	}
	p.head = previousNode
	fmt.Println(previousNode.name, previousNode.artist)
	for previousNode.next != nil {
		previousNode = previousNode.next
		fmt.Println(previousNode.name, previousNode.artist)
	}
}

func main() {
	playlistName := "myplaylist"
	myPlaylist := createPlayList(playlistName)
	fmt.Println("Created playlist")
	fmt.Println()

	fmt.Print("Adding songs to the playlist...\n\n")
	myPlaylist.addSong("name1", "artist1")
	myPlaylist.addSong("name2", "artist2")
	myPlaylist.addSong("name3", "artist3")
	myPlaylist.addSong("name4", "artist4")
	fmt.Println("Showing all songs in playlist...")
	myPlaylist.ShowAllSongs()
	fmt.Println()

	myPlaylist.startPlaying()
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	myPlaylist.playNextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()
	myPlaylist.playNextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	myPlaylist.playPreviousSong()
	fmt.Println("Changing previous song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()
	myPlaylist.playPreviousSong()
	fmt.Println("Changing previous song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)

	fmt.Println()
	fmt.Println("Reversing the playlist ")
	myPlaylist.reversePlayList()
}
