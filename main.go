package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

var (
	pathSource = "freeCodeCamp-main/curriculum/challenges/english/15-javascript-algorithms-and-data-structures-22/"
	path       = "pseudo-answer/"
)

func main() {

}

func body() {
	dir1, err := os.ReadDir("freeCodeCamp-main/curriculum/challenges/english/15-javascript-algorithms-and-data-structures-22/")
	if err != nil {
		log.Println(err)
	}

	for _, k1 := range dir1 {
		err := os.MkdirAll(path+k1.Name(), 0777)
		if err != nil {
			log.Println(err)
		}

		dir, err := os.ReadDir(pathSource + k1.Name())
		if err != nil {
			log.Println(err)
		}

		for _, k := range dir {
			fmt.Println(k.Name())
			file, err := os.OpenFile(pathSource+k1.Name()+"/"+k.Name(), os.O_RDWR, 0777)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(k.Name() + "\t" + "open file")

			data := read(file)

			fmt.Println(k.Name() + "\t" + "read file")

			numberStep := reg(string(data))

			fmt.Println(k.Name() + "\t" + "regexp")

			newFile, err := os.Create(path + k1.Name() + "/" + numberStep + ".md")
			if err != nil {
				log.Println(err)
			}

			_, err = newFile.Write(data)
			if err != nil {
				log.Println(err)
			}

			fmt.Println(k.Name() + "\t" + "new file")
			file.Close()
			newFile.Close()

			fmt.Println(k.Name() + "\t" + "close")
		}
	}
}

func reg(str string) string {
	re := regexp.MustCompile(`Step (\d+)`)
	matches := re.FindStringSubmatch(str)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func read(file *os.File) []byte {
	info, err := file.Stat()
	if err != nil {
		log.Println(err)
	}
	var data = make([]byte, info.Size())

	for {
		_, err = file.Read(data)
		if errors.Is(err, io.EOF) {
			break
		}
	}

	return data
}
