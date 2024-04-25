package embedfile

import (
	"embed"
	"io/fs"
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
	var list []string
	for _, file := range files {
		list = append(list, file.Name())
	}
	return list
}

func (e _EmbedDir) Open(name string) (fs.File, error) {
	return e.efs.Open(e.prefix + "/" + name)
}

func (e _EmbedDir) ReadFile(name string) ([]byte, error) {
	return e.efs.ReadFile(e.prefix + "/" + name)
}
