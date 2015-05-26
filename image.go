package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var listImageURL = "http://tokyo-ame.jwa.or.jp/scripts/mesh_index.js"

func ListImages() ([]string, error) {
	resp, err := http.Get(listImageURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

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
