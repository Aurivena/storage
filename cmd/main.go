package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	_ "net/url"
	"os"
	"path"
	"sync"
)

const (
	pathToVideo = "video"
	pathToPhoto = "photo"
)

const (
	typeToVideo = "typeVideo"
	typeToPhoto = "typePhoto"
)

var typesVideo = []string{".mp4", ".avi", ".mkv", ".webm", ".mov", ".wmv"}
var typesPhoto = []string{".giv", ".png", ".jpeg", ".pdf", ".bmp", ".jpg", ".jpe", ".jfif"}

func main() {
	urls := []string{
		"https://static-cse.canva.com/blob/191106/00_verzosa_winterlandscapes_jakob-owens-tb-2640x1485.jpg",
		"https://aif-s3.aif.ru/images/019/507/eeba36a2a2d37754bab8b462f4262d97.jpg",
		"https://example.com/video.mp4",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			ext, err := checkExtensionFile(url)
			if err != nil {
				fmt.Println(err)
			}

			typeFile, err := checkTypeFile(ext)
			if err != nil {
				fmt.Println(err)
			}

			err = redirectFileToDirectory(url, typeFile)
			if err != nil {
				fmt.Println(err)
			}
		}(url)
	}
	wg.Wait()
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return err
}

func checkExtensionFile(urlStr string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	fileName := path.Base(parsedURL.Path)
	ext := path.Ext(fileName)

	if ext == "" {
		return "", errors.New("не удалось извлечь расширение из URL")
	}

	return ext, nil
}

func checkTypeFile(typeFile string) (string, error) {
	for _, v := range typesVideo {
		if typeFile == v {
			return typeToVideo, nil
		}
	}

	for _, v := range typesPhoto {
		if typeFile == v {
			return typeToPhoto, nil
		}
	}
	return "", errors.New("type file not found")
}

func redirectFileToDirectory(urlStr string, typeFile string) error {
	var dir string

	if typeFile == typeToVideo {
		dir = pathToVideo
	} else if typeFile == typeToPhoto {
		dir = pathToPhoto
	} else {
		return errors.New("Ошибка. Такого типа файла не существует")
	}

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fileName := path.Base(parsedURL.Path)

	filePath := fmt.Sprintf("%s/%s", dir, fileName)

	err = downloadFile(filePath, urlStr)
	if err != nil {
		return err
	}

	fmt.Println("Файл успешно загружен:", filePath)

	return nil
}
