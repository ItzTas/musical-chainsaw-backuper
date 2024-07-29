package main

import (
	"path/filepath"

	"github.com/spf13/afero"
)

func (b *Backuper) writeDir(dst string) error {
	dst = normalizePath(dst)
	err := b.writer.RemoveAll(dst)
	if err != nil {
		return err
	}
	for _, file := range b.files {
		filep := dst + "/" + file.filePath
		fileDir := filepath.Dir(filep)
		b.writer.MkdirAll(fileDir, 0711)

		err = afero.WriteFile(b.writer, filep, file.content, 0711)
		if err != nil {
			return err
		}
	}
	return nil
}
