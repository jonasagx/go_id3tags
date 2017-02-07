package main

import (
    "fmt"
    "github.com/jonasagx/go_id3tags"
)

func main() {
    var mp3file id3tags.Mp3
    mp3file.Filename = "Pixies  - Where Is My Mind-5iC0YXspJRM..mp3"
    mp3file.Path = "musicProvider/"
    mp3file.GetID3Tags()              //read tags from mp3 file
    fmt.Println(mp3file.Title)        //Burn
    // mp3file.Artist = "Pixies"	 //set Artist
    mp3file.SetID3Tags()              //write tags to mp3file
    fmt.Println(mp3file.Artist)
}