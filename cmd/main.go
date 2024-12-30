package main

import (
	"errors"
	"fmt"
	"net/url"
	_ "net/url"
	"path"
	"storage"
	"sync"
)

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
			ext, err := storage.CheckExtensionFile(url)
			if err != nil {
				fmt.Println(err)
			}

			typeFile, err := storage.CheckTypeFile(ext)
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

func redirectFileToDirectory(urlStr string, typeFile string) error {
	var dir string

	if typeFile == storage.TypeToVideo {
		dir = storage.PathToVideo
	} else if typeFile == storage.TypeToPhoto {
		dir = storage.PathToPhoto
	} else if typeFile == storage.TypeToDocument {
		dir = storage.PathToDocument
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

	err = storage.DownloadFile(filePath, urlStr)
	if err != nil {
		return err
	}

	fmt.Println("Файл успешно загружен:", filePath)

	return nil
}
