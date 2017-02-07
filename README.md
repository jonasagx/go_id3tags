# go_id3tags

Extract and write ID3 tag metadata to mp3 files.

### Example

``` go
package main

import (
	"fmt"
	"github.com/jonasagx/go_id3tags"
)

func main() {
	var mp3file id3tags.Mp3
	mp3file.Filename = "Ellie Goulding - Burn.mp3"
	mp3file.Path = "/path/to/mp3/"
	mp3file.GetID3Tags()              //read tags from mp3 file
	fmt.Println(mp3file.Title)        //Burn
	mp3file.Artist = "Ellie Goulding" //set Artist
	mp3file.SetID3Tags()              //write tags to mp3file
}
```
