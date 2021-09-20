package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Expected a root filder path as an argument")
	}

	fmt.Println("<catalog>")
	filepath.WalkDir(os.Args[1], walkdir)
	fmt.Println("</catalog>")
}

func walkdir(path string, d fs.DirEntry, err error) error {
	if err != nil {
		log.Fatalf(err.Error())
	}

	if !d.IsDir() {
		process(path, d)
	}

	return nil
}

func process(path string, d fs.DirEntry) {
	ext := strings.TrimPrefix(filepath.Ext(path), ".")
	if ext != "pdf" && ext != "djvu" && ext != "epub" && ext != "chm" && ext != "doc" && ext != "docx" {
		return
	}

	dir := filepath.Dir(path)
	nam := withoutExt(d.Name())
	pat := withoutExt(path)
	img := imageExt(pat)
	cod := code(pat, nam)

	dir = strings.Replace(dir, "&", "&amp", -1)
	nam = strings.Replace(nam, "&", "&amp", -1)

	fmt.Printf("<book year=\"?\" author=\"?\" fmt=\"%s\" title=\"%s\"\n", ext, nam)
	if cod == "" {
		fmt.Printf("      pub=\"?\" ed=\"1\" isbn10=\"?\" isbn13=\"?\" asin=\"?\" lang=\"e\" img=\"%s\" path=\"%s\"/>\n", img, dir)
	} else {
		cod = strings.Replace(cod, "&", "&amp", -1)
		fmt.Printf("      pub=\"?\" ed=\"1\" isbn10=\"?\" isbn13=\"?\" asin=\"?\" lang=\"e\" img=\"%s\" path=\"%s\" code=\"%s\"/>\n", img, dir, cod)
	}
}

func imageExt(pathWithoutExt string) string {
	switch {
	case fileExists(pathWithoutExt + ".jpg"):
		return "jpg"
	case fileExists(pathWithoutExt + ".png"):
		return "png"
	case fileExists(pathWithoutExt + ".gif"):
		return "gif"
	default:
		log.Fatalf("No known image foud for %s", pathWithoutExt)
		return ""
	}
}

func code(pathWithoutExt, name string) string {
	switch {
	case fileExists(pathWithoutExt + ".zip"):
		return name + ".zip"
	case fileExists(pathWithoutExt + ".rar"):
		return name + ".rar"
	case fileExists(pathWithoutExt + ".7z"):
		return name + ".7z"
	default:
		return ""
	}
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func withoutExt(s string) string {
	if pos := strings.LastIndexByte(s, '.'); pos != -1 {
		return s[:pos]
	}
	return s
}
