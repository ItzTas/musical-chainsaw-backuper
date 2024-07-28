package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/afero"
)

type File struct {
	filePath string
	fileName string
	content  []byte
}

type Backuper struct {
	writer afero.Fs
	reader afero.Fs
	files  []File
}

func main() {
	src := flag.String("src", "", "source of the directory to copy")
	dst := flag.String("dst", "", "destination of the directory to copy")
	flag.Parse()

	if *src == "" || *dst == "" {
		log.Fatal("src and dst params required")
	}

	writer := afero.NewOsFs()

	reader := afero.NewReadOnlyFs(afero.NewOsFs())

	backuper := Backuper{
		reader: reader,
		writer: writer,
		files:  []File{},
	}

	err := backuper.readDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range backuper.files {
		fmt.Println(string(file.content))
	}
}
