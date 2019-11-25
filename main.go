package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var pathDir string

//TODO зарефакторить http сервер и уменьшить время обработки запроса
func main() {
	flag.StringVar(&pathDir, "p", ".", "Enter path to directory")
	flag.Parse()
	http.HandleFunc("/metrics", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("0.0.0.0:9100", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, checkDir())
}

func checkDir() string {
	size, _ := DirSize(pathDir)
	fmt.Print("check_leveldb_store ")
	fmt.Println(size)
	o := "check_leveldb_store " + strconv.FormatInt(size, 10)
	return o

}

//TODO разобрать данную функцию
func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
