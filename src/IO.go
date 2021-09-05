package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func BuildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		r.Close()
		gr.Close()
	}, nil
}

func CountLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func main() {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	couts, err := CountLetters(sr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(couts)

	data, err := ioutil.ReadFile("/Users/ashu/Documents/GoLang/src/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	myData := []byte("this is new type of data, that i am going to write in a file")

	herr := ioutil.WriteFile("/Users/ashu/Documents/GoLang/src/newFile.data", myData, 0777)
	if herr != nil {
		//print it out
		fmt.Println(herr)
	}
	

}
