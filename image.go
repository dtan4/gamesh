package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var meshURLbase = "http://tokyo-ame.jwa.or.jp/mesh/000/"
var backgroundImageURL = "http://tokyo-ame.jwa.or.jp/map/map000.jpg"
var maskImageURL = "http://tokyo-ame.jwa.or.jp/map/msk000.png"

func GetImage(id string) (string, error) {
	body, err := download(meshURL(id))

	if err != nil {
		return "", err
	}

	path, err := savePath(id)

	if err != nil {
		return "", err
	}

	err = saveToFile(body, path)

	if err != nil {
		return "", err
	}

	return path, nil
}

var listImageURL = "http://tokyo-ame.jwa.or.jp/scripts/mesh_index.js"

func ListImages() ([]string, error) {
	body, err := download(listImageURL)

	if err != nil {
		return nil, err
	}

	substr := strings.Replace(string(body), "Amesh.setIndexList(", "", -1)
	substr = strings.Replace(substr, ");", "", -1)
	substr = strings.Replace(substr, "\n", "", -1)
	substr = strings.Replace(substr, "\"", "", -1)
	substr = strings.Replace(substr, "[", "", -1)
	substr = strings.Replace(substr, "]", "", -1)
	images := strings.Split(substr, ",")

	return images, nil
}

func download(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func meshURL(id string) string {
	return meshURLbase + id + ".gif"
}

func savePath(id string) (string, error) {
	pwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return filepath.Join(pwd, id+".jpg"), nil
}

func saveToFile(body []byte, path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	defer file.Close()
	file.Write(body)

	return nil
}
