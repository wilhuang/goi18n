package embedfile

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
)

//go:embed lang/*
var fs_lang embed.FS
var LANG = _EmbedDir{
	efs:    fs_lang,
	prefix: "lang",
}

type _EmbedDir struct {
	efs    embed.FS
	prefix string
}

func (e _EmbedDir) List() []string {
	files, err := e.efs.ReadDir(e.prefix)
	if err != nil {
		return nil
	}
	list := make([]string, len(files))
	for i, file := range files {
		list[i] = file.Name()
	}
	return list
}

func (e _EmbedDir) Open(name string) (fs.File, error) {
	return e.efs.Open(e.prefix + "/" + name)
}

func (e _EmbedDir) ReadFile(name string) ([]byte, error) {
	return e.efs.ReadFile(e.prefix + "/" + name)
}

func (e _EmbedDir) Copy(dst, src string) error {
	srcf, err := e.Open(src)
	if err != nil {
		return err
	}
	defer srcf.Close()

	dstf, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstf.Close()

	n, err := io.Copy(dstf, srcf)
	if err != nil {
		return err
	}

	srcInfo, err := srcf.Stat()
	if err != nil {
		return err
	}

	if n < srcInfo.Size() {
		return fmt.Errorf("not all bytes copied from source to destination: %d < %d", n, srcInfo.Size())
	}

	return nil
}
