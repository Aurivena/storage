package storage

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

func DownloadFile(filepath string, url string) error {
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

func CheckTypeFile(typeFile string) (string, error) {
	for _, v := range TypesVideo {
		if typeFile == v {
			return TypeToVideo, nil
		}
	}

	for _, v := range TypesPhoto {
		if typeFile == v {
			return TypeToPhoto, nil
		}
	}

	for _, v := range TypesDocument {
		if typeFile == v {
			return TypeToDocument, nil
		}
	}
	return "", errors.New("type file not found")
}
