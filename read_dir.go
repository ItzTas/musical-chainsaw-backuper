package main

import (
	"log"

	"github.com/spf13/afero"
)

func (b *Backuper) readDir(src string) error {
	src = normalizePath(src)
	files, err := afero.ReadDir(b.reader, src)
	if err != nil {
		return err
	}
	for _, file := range files {
		path := src + "/" + file.Name()
		if file.IsDir() {
			err := b.readDir(path)
			if err != nil {
				log.Printf("Error reading directory %s: %v\n", path, err)
			}
			continue
		}
		content, err := afero.ReadFile(b.reader, path)
		if err != nil {
			log.Printf("Error reading file %s: %v\n", path, err)
			continue
		}
		f := File{
			filePath: path,
			fileName: file.Name(),
			content:  content,
		}
		b.files = append(b.files, f)
	}
	return nil
}
