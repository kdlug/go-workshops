package main

import "fmt"

type song struct {
	title, artist string
}

type playlist struct {
	genre string
	songs []song
}

func main() {
	song1 := song{title: "Fly Away", artist: "Lenny Kravitz"}
	song2 := song{title: "Layla", artist: "Eric Clapton"}

	fmt.Printf("Song 1: %+v\nSong 2: %+v\n", song1, song2)
	if song1 != song2 {
		fmt.Printf("The songs are different\n")
	}
	// copy
	song2 = song1

	fmt.Printf("Song 1: %+v\nSong 2: %+v\n", song1, song2)
	if song1 == song2 {
		fmt.Printf("The songs are the same\n")
	}

	// struct with [] slice
	songs := []song{
		song1,
	}
	rock := playlist{genre: "rock", songs: songs}

	// -> the following doesn't change the song name, because song is a clone/copy
	song := rock.songs[0]
	song.title = "I belong to you"
	// the following does
	// rock.songs[0].title = "I belong to you"

	fmt.Printf("%-20s %20s\n", "TITLE", "ARTIST")

	for _, s := range rock.songs { // s is a copy of a struct
		// s := rock.songs[i] // in the background
		s.title = "destroy"
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}

	for _, s := range rock.songs {
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}
}
