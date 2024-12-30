package main

import (
	"fmt"
	_ "net/url"
	"storage"
	"sync"
)

func main() {
	urls := []string{
		"C:\\Users\\vladk\\Pictures\\Screenshots\\Снимок экрана 2024-12-25 090242.png",
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

			typeFile, typ, err := storage.CheckTypeFile(ext)
			if err != nil {
				fmt.Println(err)
			}

			err = storage.RedirectFileToDirectory(url, typeFile, typ)
			if err != nil {
				fmt.Println(err)
			}
		}(url)
	}
	wg.Wait()
}
