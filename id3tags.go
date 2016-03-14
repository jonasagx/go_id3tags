package id3tags

import (
	"os"
)

//Mp3 ...Contains metadata of the file like title, artist, album, year info along with file name and its path
type Mp3 struct {
	Filename string
	Path     string
	Title    string
	Artist   string
	Album    string
	Year     string
}

func getLastBytes(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	_, err = file.Seek(-int64(128), os.SEEK_END)
	if err != nil {
		panic(err)
	}
	b := make([]byte, 128)
	_, err = file.Read(b)
    if err!=nil{
        panic(err)
    }
	return b, nil
}

func setLastBytes(filename string, b []byte)(error){
	file, err:=os.OpenFile(filename,os.O_RDWR,0655)
	if err!=nil{
		panic(err)
	}
	_,err = file.Seek(-int64(128),os.SEEK_END)
	if err!=nil{
		panic(err)
	}
	_,err=file.Write(b)
	if err!=nil{
		panic(err)
	}
	return nil
}

//GetID3Tags ...Extracts mp3 metadata from the file
func (m *Mp3) GetID3Tags() {
	b, _ := getLastBytes(m.Path + m.Filename)
	if string(b[:3]) == "TAG" {
		m.Title = string(b[3:33])
		m.Artist = string(b[33:63])
		m.Album = string(b[63:93])
		m.Year = string(b[93:97])
	}
}
//SetID3Tags ...Writes tag metadata to file
func (m *Mp3) SetID3Tags() {
	b := make([]byte,128)
    copy(b[:],"TAG")
    copy(b[3:33],m.Title)
    copy(b[33:63],m.Artist)
    copy(b[63:93],m.Album)
    copy(b[93:97],m.Year)
    setLastBytes(m.Path+m.Filename,b)
}