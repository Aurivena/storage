package storage

import (
	"errors"
	"fmt"
	uuid2 "github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

func RedirectFileToDirectory(urlStr string, typeFile string, typ string) error {
	var dir string

	if typeFile == TypeToVideo {
		dir = PathToVideo
	} else if typeFile == TypeToPhoto {
		dir = PathToPhoto
	} else if typeFile == TypeToDocument {
		dir = PathToDocument
	} else if typeFile == TypeToArchive {
		dir = PathToArchive
	} else {
		return errors.New("Ошибка. Такого типа файла не существует")
	}

	uuid, err := uuid2.NewV7()
	if err != nil {
		return err
	}

	var buf strings.Builder

	buf.WriteString(uuid.String())
	buf.WriteString(time.Now().UTC().Format("2006-01-02_15-04-05"))

	fileName := buf.String()

	filePath := fmt.Sprintf("%s/%s%s", dir, fileName, typ)

	err = DownloadFile(filePath, urlStr)
	if err != nil {
		return err
	}

	fmt.Println("Файл успешно загружен:", filePath)

	return nil
}

func DownloadFile(filepath string, url string) error {

	if !(strings.HasPrefix(url, "http") || strings.HasPrefix(url, "https")) {
		file, err := os.Open(url)
		if err != nil {
			return fmt.Errorf("не удалось открыть локальный файл: %v", err)
		}
		defer file.Close()

		out, err := os.Create(filepath)
		if err != nil {
			return fmt.Errorf("не удалось создать файл: %v", err)
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			return fmt.Errorf("не удалось скопировать данные: %v", err)
		}
		return nil
	}

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

func CheckExtensionFile(urlStr string) (string, error) {

	if strings.Contains(urlStr, "https") || strings.Contains(urlStr, "http") {
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
	} else {
		var buf strings.Builder
		var res strings.Builder
		for i := len(urlStr) - 1; i >= 0; i-- {
			if urlStr[i] == '.' {
				res.WriteByte('.')
				for j := buf.Len() - 1; j >= 0; j-- {
					res.WriteByte(buf.String()[j])
				}

				return res.String(), nil
			} else {
				buf.WriteByte(urlStr[i])
			}
		}

		return "", nil
	}

}

func CheckTypeFile(typeFile string) (string, string, error) {
	for _, v := range TypesVideo {
		if typeFile == v {
			return TypeToVideo, typeFile, nil
		}
	}

	for _, v := range TypesPhoto {
		if typeFile == v {
			return TypeToPhoto, typeFile, nil
		}
	}

	for _, v := range TypesDocument {
		if typeFile == v {
			return TypeToDocument, typeFile, nil
		}
	}

	for _, v := range TypesArchive {
		if typeFile == v {
			return TypeToArchive, typeFile, nil
		}
	}

	return "", "", errors.New("type file not found")
}
