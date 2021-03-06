package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestListImages(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintln(w, `Amesh.setIndexList(["201505270545","201505270540","201505270535","201505270530","201505270525","201505270520","201505270515","201505270510","201505270505","201505270500","201505270455","201505270450","201505270445","201505270440","201505270435","201505270430","201505270425","201505270420","201505270415","201505270410","201505270405","201505270400","201505270355","201505270350","201505270345"]);`)
	}))

	defer server.Close()

	// TODO: mock HTTP client
	listImageURL = server.URL

	c.Convey("Get image list", t, func() {
		images, err := ListImages()

		c.Convey("It should raise no error", func() {
			c.So(err, c.ShouldBeNil)
		})

		c.Convey("It should return the image list", func() {
			expected := []string{"201505270545", "201505270540", "201505270535", "201505270530", "201505270525", "201505270520", "201505270515", "201505270510", "201505270505", "201505270500", "201505270455", "201505270450", "201505270445", "201505270440", "201505270435", "201505270430", "201505270425", "201505270420", "201505270415", "201505270410", "201505270405", "201505270400", "201505270355", "201505270350", "201505270345"}

			c.So(images, c.ShouldResemble, expected)
		})
	})
}
