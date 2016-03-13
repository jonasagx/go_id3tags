package main

import (
	"fmt"
	//"io/ioutil"
	"os"
	//"strings"
	"log"
)

type mp3 struct {
	filename string
	path string
	title  string
	artist string
	album  string
	year   string
}

func getLastBytes(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Seek(-int64(128), os.SEEK_END)
	if err != nil {
		fmt.Println(err)
	}
	b := make([]byte, 128)
	_, err = file.Read(b)
	return b, nil
}

func setLastBytes(filename string, b []byte)(error){
	file, err:=os.OpenFile(filename,os.O_RDWR,0655)
	if err!=nil{
		log.Fatal(err)
	}
	_,err = file.Seek(-int64(128),os.SEEK_END)
	if err!=nil{
		log.Fatal(err)
	}
	_,err=file.Write(b)
	if err!=nil{
		log.Fatal(err)
	}
	return nil
}

func (m *mp3) getID3Tags() {
	b, _ := getLastBytes(m.path + m.filename)
	if string(b[:3]) == "TAG" {
		m.title = string(b[3:33])
		m.artist = string(b[33:63])
		m.album = string(b[63:93])
		m.year = string(b[93:97])
		fmt.Println(m.title,m.artist,m.album,m.year)
	}
}
func (m *mp3) setID3Tags() {
	b:=make([]byte,128)
	setLastBytes(m.path+m.filename,b)
}
