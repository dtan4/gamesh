package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func ListImages() (string, error) {
	resp, err := http.Get("http://tokyo-ame.jwa.or.jp/scripts/mesh_index.js")

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	images := strings.Replace(string(body), "Amesh.setIndexList(", "", -1)
	images = strings.Replace(images, ");", "", -1)
	images = strings.Replace(images, "\n", "", -1)
	images = strings.Replace(images, "\"", "", -1)
	images = strings.Replace(images, "[", "", -1)
	images = strings.Replace(images, "]", "", -1)
	images = strings.Replace(images, ",", "\n", -1)

	return images, nil

}
