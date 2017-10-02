// Code generated by go-bindata.
// sources:
// www/index.html
// www/javascript/mapzen.whosonfirst.catalog.init.js
// www/javascript/mapzen.whosonfirst.catalog.js
// www/javascript/mapzen.whosonfirst.geojson.js
// www/javascript/mapzen.whosonfirst.render.js
// www/css/mapzen.whosonfirst.catalog.css
// www/css/mapzen.whosonfirst.render.css
// DO NOT EDIT!

package http

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _wwwIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\xb1\x8e\xdb\x30\x0c\x86\xe7\xcb\x53\x10\x5c\xba\x45\x2f\x20\x79\x69\x7b\x40\x81\x02\x37\x74\xb8\xb1\x50\x2c\xfa\xac\x44\x96\x5c\x91\xf2\xd5\x7d\xfa\xc2\x96\x73\x97\x14\xdd\x72\xf1\x62\xf9\x27\xfd\xf1\x27\x4d\x58\xf7\x32\x84\x66\x07\xa0\x7b\xb2\x6e\x39\x00\x68\xf1\x12\xa8\x79\xee\xd3\x27\x86\xa7\x08\x8f\x3e\xb3\xc0\x67\x2b\x36\xa4\x17\xf8\x31\x52\x28\xf1\x44\x59\xab\x9a\x57\xdf\x09\x3e\x9e\x40\xe6\x91\x0c\x0a\xfd\x16\xd5\x32\x23\x64\x0a\x06\x59\xe6\x40\xdc\x13\x09\x42\x9f\xa9\x33\xd8\x32\xab\xc1\x8e\x7f\x28\xee\x8f\xbc\x5f\x33\xd5\x6d\x98\xd7\x3e\x71\x8a\xdd\x62\x74\xdf\x56\xa3\x1f\xce\xcd\x14\x1d\xe5\x6b\x2c\xb7\xd9\x8f\x72\x09\x3e\xda\xc9\x56\x15\x81\x73\x6b\xf0\x5d\x38\x43\x07\xbf\xf4\x8d\x8d\x56\x55\xbf\x85\x75\x69\xf0\x85\xd2\x91\xd3\x7d\xd8\x5b\xf3\xf7\x40\x9f\xbf\xd7\x3d\xd9\x3e\x7a\xb9\x2e\x00\xdb\xb5\xac\xbe\x3a\xef\xbe\x3e\x24\x37\x6f\xb5\x9d\x9f\xc0\x3b\x83\xad\x8d\x93\x65\xac\xea\x9b\x0e\x06\x7f\x15\xca\x33\x36\xbb\x07\xdd\xa5\x3c\x34\xbb\x07\x00\xed\xe3\x58\x2e\xfd\xe2\x8a\x78\x4d\xdd\x4f\xef\x10\xa2\x1d\xe8\xfd\x69\xb2\xa1\x90\x41\x84\x31\xd8\x96\xfa\x14\x1c\x65\x83\x5f\xa3\x50\x06\x0b\xcf\x4f\x8f\xf0\xed\xcb\xba\x69\x0b\xf8\x50\x44\x52\xdc\xc8\x5c\x0e\x83\xdf\xd8\x21\xa5\x53\x19\xb1\xf9\xbe\xde\xb5\xaa\x89\x8b\x2b\x55\x6d\x6d\xb6\x95\xf3\xd3\x3f\x3d\x18\x74\x24\xd6\x87\x75\x2c\xff\x0b\x67\xe2\x12\xe4\x3a\xfc\x76\xd4\xaa\x0e\x4b\xab\xfa\x0b\xf9\x1b\x00\x00\xff\xff\xdd\xde\xa7\xcd\x4a\x04\x00\x00")

func wwwIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwIndexHtml,
		"www/index.html",
	)
}

func wwwIndexHtml() (*asset, error) {
	bytes, err := wwwIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/index.html", size: 1098, mode: os.FileMode(420), modTime: time.Unix(1506709930, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwJavascriptMapzenWhosonfirstCatalogInitJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\xc1\x6a\xc3\x30\x10\x44\xcf\xeb\xaf\x58\x7c\x92\x21\xe8\x07\x4c\x0e\x6d\x09\x34\xd0\x63\xa1\xc7\xa2\x4a\xeb\x44\x44\x59\x19\x69\x15\x93\x86\xfc\x7b\x91\x6d\x9a\x5b\x4e\xd2\xa2\x99\xc7\x8c\x76\xf2\xec\xe2\xa4\x8d\x73\xbb\x0b\xb1\x7c\xf8\x2c\xc4\x94\x54\x1b\xa2\x71\xed\x06\x87\xc2\x56\x7c\x64\xac\xb3\xa2\xaa\xe9\x6e\x0d\x34\x70\x31\x09\x7f\x8a\x48\x64\xdc\xa2\x8b\xb6\x9c\x89\x45\x1f\x48\x76\x81\xea\xf5\xf5\xba\x77\x95\x12\x4f\x65\x6c\xbb\xbe\x5a\x16\xb9\x8e\x6c\x83\xb7\x27\xdc\xfe\xc3\x55\x77\x6b\x1a\x00\x49\x57\xbc\x35\x88\x88\x00\x33\xdf\xbb\x67\xec\x29\x0e\xdf\xde\xcd\x6c\x80\x59\x3a\x9a\x94\x69\xcf\xa2\xbc\xd3\x17\x13\x0a\x75\x7d\xe5\xc2\xd9\x8c\xbf\xc4\x7a\x3a\xc6\x1c\x79\xf0\x29\x8b\xb6\x46\x4c\x88\x07\xbd\x04\x54\xde\x3d\xaa\xaa\x94\xc7\x25\xd0\x92\x22\xe1\x16\x9f\x10\x12\xb1\xa3\x34\x9b\xe6\x24\xf0\x30\x52\x2e\x41\xf2\xb3\x0e\xab\x64\x2d\x01\xeb\xa8\xcd\x38\x12\xbb\xb7\xa3\x0f\x4e\xa5\xe5\xed\x3e\x1f\xf7\x1a\xcb\x1a\xb1\x47\x54\x54\x17\x01\x60\x23\xe7\x18\x48\x87\x78\x50\xed\xd7\xfb\xcb\x67\xbb\x41\x5a\xc5\x00\xeb\x77\x26\x92\x92\x18\x07\x13\x32\xf5\x0d\xdc\xfb\xa6\xf2\xfe\x02\x00\x00\xff\xff\xae\xcc\x05\x3a\xfd\x01\x00\x00")

func wwwJavascriptMapzenWhosonfirstCatalogInitJsBytes() ([]byte, error) {
	return bindataRead(
		_wwwJavascriptMapzenWhosonfirstCatalogInitJs,
		"www/javascript/mapzen.whosonfirst.catalog.init.js",
	)
}

func wwwJavascriptMapzenWhosonfirstCatalogInitJs() (*asset, error) {
	bytes, err := wwwJavascriptMapzenWhosonfirstCatalogInitJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/javascript/mapzen.whosonfirst.catalog.init.js", size: 509, mode: os.FileMode(420), modTime: time.Unix(1505838717, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwJavascriptMapzenWhosonfirstCatalogJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x5b\x73\xdb\xb8\x15\x7e\x96\x7e\x05\x16\x9d\x59\x93\x13\x85\xce\xa5\x0f\x1d\xab\x7a\x68\xb6\xce\x5e\xea\x24\x9d\xc4\xed\xec\x0c\x47\xd5\xc0\xe4\xb1\x84\x35\x45\x30\x00\x14\x45\xcd\xfa\xbf\x77\x0e\xee\x94\x68\x5b\x6a\x5f\x9a\x87\xd8\x04\xbe\x73\xc1\x39\x1f\xce\x01\xe0\x2f\x4c\x92\x35\xeb\xfe\x0d\x2d\x99\xf9\x5f\x7e\xff\x9d\x7c\xbb\x9f\x8e\xed\x57\xb1\x5d\x09\x25\xda\x5b\x2e\x95\x0e\x90\xde\xa0\x83\x0f\xe0\x8b\x8a\x69\xd6\x88\x25\x99\x91\xec\x76\xd3\x56\x9a\x8b\x36\xcb\xbf\x8d\xc7\x84\x10\x82\x96\x2b\x56\xad\x80\xcc\xac\xbc\x1f\x54\xd0\xdc\xe2\xd8\x78\x34\x1e\x9d\x35\x42\xdc\x6d\xba\xb3\x0b\x12\xe4\x79\x3d\x21\xd5\x4d\xfe\x6d\x3c\x42\x01\xfb\x3f\x8a\x75\x52\x68\x41\x66\xa4\x16\xd5\x66\x0d\xad\x2e\x1a\x51\x31\x94\x28\xcc\x4c\x25\x9a\x69\x04\x6f\x64\x43\x66\x4e\xe4\x19\x39\x3b\x3f\x3f\x23\xcf\x48\x10\x58\x09\xa5\xd3\xef\x8e\xe9\x55\xcb\xd6\x80\x50\x5e\x1b\x2c\xaf\xa7\xfb\x0e\x88\xb6\x11\xac\x26\xb3\xe8\xaa\x54\x1d\xfa\x39\x1a\x8f\x46\x08\xd0\x4c\x2e\x01\x83\x28\x55\x57\xd8\x8f\xa9\x9d\xe5\xb7\x24\xb3\x03\x85\x04\x56\xef\x3e\x69\xa6\x81\x7c\x37\x23\x7f\x34\xf2\x68\x41\x82\xde\xc8\x16\xf1\xf7\x51\xa3\xd2\x4c\x6f\xd4\xa2\x12\x35\x46\xd1\x6a\x28\xcf\xec\xe8\xd9\x1c\xc1\x3d\x9c\x86\xaf\xfa\x00\x77\x0d\x5f\xb5\xc5\x3a\xa5\x92\x6d\x13\x90\x04\xd5\x89\x56\x41\x84\x21\xa6\x66\x9a\x91\x19\x69\x37\x4d\xe3\x24\xb5\xdc\x11\xef\xab\x9b\xfd\xe5\xd3\x87\xf7\x45\xc7\xa4\x82\x4c\xb2\x6d\x9e\xf8\x5e\x31\x5d\xad\x48\x06\x61\x75\x95\x68\x95\x68\xa0\x68\xc4\x32\xa3\x97\x1f\x3f\xd2\x09\xa6\x68\x42\x20\xef\x2d\x9f\xdc\xb2\x46\x41\xaa\xe8\x26\x43\x63\xb9\x4b\xc6\xfd\x41\x52\x24\x7c\x46\x3f\x61\x4b\x7e\x7d\x77\xf5\x93\xd6\xdd\x47\xf8\xbc\x01\xa5\x33\x2f\x22\xe1\x73\xc1\xea\xfa\xf2\x0b\xb4\xfa\x8a\x2b\x0d\x2d\xc8\x8c\x62\x26\xe9\xc4\xa5\x34\xef\x69\x45\x01\xd1\x41\x9b\xd1\x1f\x2f\xaf\xbd\x9f\x5a\x6e\x20\xd5\xa8\xa0\xad\x8d\x89\xfb\x89\xa1\xb1\x84\xb6\x06\x99\xd2\xd8\x71\x63\xdf\xdb\x4a\xc8\x5a\x59\x8a\x94\xd4\x7d\x82\xa6\xf3\xf0\x41\xe7\x09\x8d\x2b\xb1\x69\x0d\xa1\xec\x5c\xd1\x40\xbb\xd4\xab\x83\x20\xdc\xec\x16\x7a\xd7\xf9\x8d\x16\x86\x71\x0c\x6d\x95\xa9\x4a\x25\x36\xb2\xea\x0f\xdb\xff\x6f\x85\x24\x19\x22\xf8\xec\xc5\x94\x70\xf2\x67\x6b\x7d\x4a\xf8\xb3\x67\x29\xcb\x65\xf4\xa7\xe4\x81\x30\xc6\xcb\x92\xa2\x49\x3a\x4f\x58\xff\x9d\x77\xae\xd4\xf3\xc0\x06\xf7\x23\xce\xa0\x37\x44\x12\x23\x88\x33\xc6\xf3\xa2\xdb\xa8\x55\xa6\x53\x5e\x41\xa3\x20\xb0\x30\x8a\x5b\xa4\xf4\x48\xc3\x94\x74\x69\x56\x9d\x12\x32\xd2\x22\x89\x12\xbb\x69\x20\xad\x2b\x95\x04\xa6\xe1\xb2\x01\xfc\xca\xa8\x99\xa7\x87\x72\x36\x90\x8b\x15\xb0\x1a\xe4\x63\xf2\xab\x20\xdc\x13\x29\x58\xd7\x41\x5b\xff\xb0\xe2\x4d\x9d\xed\x09\xe3\x66\x7c\x2f\x6a\xc8\xa8\x15\xa1\xf9\x80\xdf\xbb\xee\x34\xeb\x89\xc0\x71\xb6\x4d\x32\x07\x2c\xaf\x98\x5a\x9d\x64\x39\x11\x38\xce\x32\x0a\x0c\x59\xde\x48\x7e\x92\xe1\x88\x3f\xce\xee\x46\xf2\xc1\x50\xf3\x35\x6f\x97\xa7\x05\x3b\x15\x39\x32\xdc\x46\x64\xc8\xbe\x5a\x89\xed\x69\x44\x8b\x02\xc7\xd9\x7e\x3e\x64\xb6\x06\xcd\x78\xa3\x4e\xb2\xdc\x97\xf9\x1f\x8c\x5b\x05\x0b\x29\xb6\x8f\x19\x96\x91\x63\x01\xdf\x33\xda\xdb\x72\x4f\x60\x93\x0d\xf2\x04\x32\x21\xf4\x13\xc8\xc8\xc0\xa7\x8c\xa7\x84\x79\x02\xdb\x0f\x72\x3f\x76\xa6\x5c\xf5\x9d\x0d\x4a\xf2\xe1\x82\xaf\x09\x6f\x6d\x8d\xf4\x75\xde\x9c\x22\x70\xa0\xd4\xb1\xc2\x63\x74\x62\x03\x8b\xd5\x37\x39\x5a\x98\x8e\x91\x80\x52\x99\xd8\xbc\x3c\xfa\x66\xb7\xb0\xe9\xf1\x8d\x6b\x34\xd0\xa0\x10\xfc\x40\x6b\xf2\x9a\x63\x8b\x4a\x9a\x4b\x72\x8a\x49\xbd\x70\x2d\x6b\x0f\x18\xdc\x40\x89\xd2\x17\xdd\xf9\x14\xcf\x57\x09\xcc\x35\x59\x0b\x8a\xad\x2e\xa9\x4e\x61\x16\x2b\x49\x6f\xd2\xe6\x37\x4a\xdb\xcd\x3e\x27\xe7\xe4\xe5\x0b\xff\x6f\x3a\x3a\x3f\x27\x2d\x6b\x85\x82\x4a\xb4\xb5\x3a\x74\x14\x89\x17\x74\x98\x2a\x69\x9c\x7c\x68\x45\x8b\x0a\x9a\xe6\xe9\x8d\x9b\x36\x27\x94\x38\x6a\xdb\x5a\x7c\x9e\x0f\x84\xd3\x04\xfc\x29\xd3\x75\x34\x1d\xf0\x47\x19\x46\xf4\xa0\x59\xb3\x2b\x4f\x30\x1b\xf0\x47\x99\x45\x34\x9a\xed\x27\xfc\x14\x7b\x1e\x5e\x28\xd0\x7f\xd1\x5a\xf2\x9b\x8d\x86\x8c\x62\x32\x9f\x23\x5f\xf0\xac\xc9\xf3\x29\x82\xf7\x25\x8e\x71\x10\x85\x07\xb3\x61\x2b\xcb\x29\xf9\x88\x12\xc7\x65\xc4\xe0\x0b\x2d\xde\xf2\xaf\x50\x67\xaf\x73\xf2\x8c\x50\x45\xf7\xa2\xe5\xab\xd6\x09\x8e\xa4\x22\x7b\x51\xab\x1a\xa6\x14\x9d\x10\x5a\x35\xbc\xba\x7b\xbe\x86\xa3\xa4\x1e\x89\x75\x4f\xea\xa8\xbe\xe5\x24\xf6\xd7\xb9\x06\xcd\x8e\xea\x5b\x4e\xc6\x5c\x0e\x4d\x65\x99\x11\x0a\x52\x0a\x49\x4d\x31\x1b\x79\x45\x0f\xad\xdc\x61\xfd\xb2\xef\xbd\xc2\x20\x37\xd0\x08\x71\x75\x41\x62\x10\x18\x36\xe3\xe3\xb0\xb0\x79\x1e\x87\x79\x06\x3f\x61\x33\x12\xee\x71\x60\x9a\xa4\x7d\xae\xf7\x12\x28\x5a\xc3\x8b\xf4\xb6\x6e\x6f\xa3\x58\x2b\x4d\x9f\x01\x24\x21\x24\xf7\xf4\x51\xac\xe1\xd0\x14\xcb\x61\xde\xe4\xae\x25\x8d\x14\x34\xb7\x85\x39\x67\x39\xbb\x66\xff\x85\x5c\xec\xf9\x96\x76\xe5\xcc\xaf\x2c\x52\xc0\xbc\x93\x94\x1b\xc9\xe7\xae\xb8\x4f\xad\xe8\xd0\x65\xc6\xdd\x94\x8d\x46\x7b\x01\x1d\x8f\xce\x56\xbc\x06\xef\x48\x7a\x0b\xcd\xbf\x8d\x0f\xce\x73\x29\x2f\x97\xa0\x1d\x29\xdf\xec\x7e\xae\x13\x4a\xf7\x8f\x73\x05\x6f\x5b\x90\x3f\x5d\xbf\xbb\x22\x33\x42\x69\xbc\xf7\xa6\x01\x48\xed\x62\x2c\xbc\x69\x13\xa9\xd4\xc1\xcc\x2c\x7c\xaf\x49\xc7\x10\x24\x97\xd5\x1b\x51\xef\x42\xbb\xc3\x8f\xde\xed\x78\xb0\x1d\xc7\xe9\xad\xc4\x78\x3f\x76\x6e\xad\xf9\x97\xb0\x52\x87\xde\xdb\x6b\xbc\xc6\x8d\xe6\x2f\x61\xa9\xf6\x5b\x60\x7a\x23\xd1\xbe\x59\x1f\x3a\xb1\xd0\x62\xe1\x86\x7b\xaf\x16\xee\xf5\xaa\xc3\xc8\xbb\xf9\x92\xe2\x00\x48\xcd\x41\x45\xaf\x3d\x68\xe0\xa5\xcd\x3e\x30\xb8\x1f\x0b\xd4\x9e\x19\xf4\xbe\xfb\xe9\x66\xf1\x80\xe8\xc5\x9a\x75\x47\x86\x63\xcd\xba\xc1\x50\xac\x59\xd7\x8f\xc3\x69\xd7\x67\x23\xf1\xe8\x55\x42\xf6\xd5\x37\x70\xab\x8f\xec\x18\xe1\x91\x85\x2f\x57\x47\xca\x58\xa1\x60\x23\x6c\x50\xd6\x85\xc7\x9e\xa0\xcc\x4f\xba\x48\x07\x69\xdd\x0f\x7a\x50\x16\x6e\x82\xfd\xf9\xa8\x30\x6a\x38\x38\xb2\x6b\xd9\xdf\x22\xff\xfd\xc6\xed\x69\x45\x3b\x41\xb1\xa5\xad\x64\xdb\xc5\x9a\x75\x99\xe3\x65\x1e\x2a\x8a\x9f\x49\x77\xb5\x07\xa5\x45\xe5\x28\x62\xdb\x97\x20\xcb\x47\x2c\xc2\xe1\xb1\xd3\xf7\xad\xb0\xe1\x6f\xc4\xd7\xe1\x0d\xb0\x04\xf1\x9b\x12\x6d\x51\x83\xe4\x5f\x60\x81\xc0\xc4\xe9\xe4\xae\x8c\x7d\xf7\xaa\x68\x98\xbe\x6a\x97\x19\xc2\xca\x97\xf3\x89\x51\x5c\xbe\x98\xa7\x3c\x69\xe1\x00\xf9\xda\x23\x5f\xcd\x83\xd2\xde\xd3\xe5\x9b\x0f\xff\x78\xff\xd7\x4f\x74\x42\xd4\x76\x42\xda\xf0\x24\xd8\xc3\x5c\x5f\x7e\xba\xa6\x13\x92\xa1\x27\x33\x04\xe5\xee\xf1\xdb\xff\xbb\x2a\xde\xd9\xf5\xb1\x8e\xff\x0d\x76\x69\x5a\xb1\xc6\x0d\xb5\x1e\x1b\x90\xe7\xac\xe3\xcf\xef\x60\x17\x5f\xa3\x52\xb5\x76\x73\x07\xe5\x98\xd5\x33\x4c\x60\x9e\x66\x21\x53\xdb\x92\x36\x4c\xd3\xb9\x75\xce\x7d\xe4\xe4\xfb\xef\x89\x9d\x33\xb7\x12\x3f\x87\x1f\xb9\x49\x99\x2b\x08\xff\xe4\xb0\xcd\x70\xf1\x2f\x5f\x85\x47\xd9\xb4\x3b\x85\x17\x3a\xc4\xdf\x72\xfd\x46\x6c\xda\x5a\x65\xa5\x8d\xd7\x3c\x1f\xc8\x79\xc3\x76\xa6\x4a\x5f\x61\x8e\x7f\xf9\xf4\xe1\x7d\x4a\x46\xb3\x3f\x11\x50\xb0\xba\xbe\x16\x7e\x6b\x3a\x8e\xf6\x8b\x6e\x4a\x55\x53\x7f\x0f\xde\x5f\x1f\xe9\x18\xbd\x53\x98\x23\x1b\x4d\xc8\x3a\xd4\x84\x7a\x2b\xef\x29\xe8\x84\xd2\x4b\xae\x68\xfa\x70\x3a\xdc\xcc\x46\xf8\x5b\x49\xdf\x81\x66\x18\xf7\xf4\x51\x3d\x9d\xc9\x93\x7b\xf3\x12\xc4\x3a\x79\x62\xb5\xb0\x1f\x41\xac\x69\x7c\x62\x45\xcc\x90\x36\x07\x7b\xe8\x4d\xf5\x21\xb1\x1f\xa0\xd5\x52\xf0\xba\x2f\x7a\x7e\x4e\x6a\x68\x40\x0f\x28\xdf\x9f\xda\x53\xe0\x56\xe2\x8b\x07\x62\x92\xe1\xd8\x5b\xbf\x25\x37\xc4\x0b\x42\xdf\xda\x09\x3a\x49\xdc\x05\x2d\x77\x17\xe6\x37\x3f\x1a\x2b\xd0\x85\xb5\x30\x89\x2e\xfb\xbf\x33\x58\x45\x7b\x6c\xec\x9f\xc4\x1b\xa6\x34\xaf\x14\x30\x59\xad\xa8\xa9\x79\x3d\x9f\xf7\xf3\x98\x94\x2f\x03\x29\x91\x45\xeb\x0b\x1c\x72\x89\xb6\x93\xf8\xa3\x50\x5d\xc3\x75\x46\x27\xee\x1a\x60\x0b\x57\x23\x5a\x37\x5f\xbe\x08\x2a\xd5\xb6\x61\xda\x0f\xbf\x0c\xc3\x2d\x24\xe8\x57\xe9\x70\x44\xbf\x9e\x07\xe5\x95\x70\x2f\x32\xa5\x8b\x51\x69\x15\x4f\x9c\xd9\xf9\x24\x8c\x1b\x15\x0f\x8f\x5b\xc3\xc9\xb8\xd3\xf3\xd0\xb8\xd7\x13\x8e\xc3\xd1\x29\x47\x36\x9f\x63\xbb\x21\x2f\x08\xfd\xbb\x68\x76\x4b\xd1\x86\x2c\x53\xe3\x3d\x6f\x99\x06\x45\x2f\x48\xe9\x57\x33\xb7\x79\x7d\x90\x34\x41\xe1\x3e\x6d\xa8\xe7\x0d\xed\x13\x27\xed\x5d\x29\x75\x8c\xc7\xc7\x33\x47\xf3\x06\x5e\xff\x69\x80\x32\xf1\x91\xcb\x2f\xfd\x3e\x04\x83\xd7\x81\x53\xb8\x4b\x3c\xd0\x2c\x88\xfe\x8b\x9a\x3f\x49\xe2\xb5\xfa\x0f\x34\x88\x0c\x16\x94\xf4\xb9\xec\x8e\xf0\xd6\xa0\xac\x2b\xb1\x13\xdf\x15\x6b\xa6\xab\x55\x26\xc1\x16\xf7\x51\x25\x5a\xcd\xdb\x0d\x1c\xdc\x25\x11\x7f\x47\xbe\x9b\x79\xeb\x78\x77\xa1\xf1\x22\xe5\xd6\x61\xf6\xf8\xdd\xbc\xa4\xe2\xe6\x37\xa8\xb4\x75\x7f\xbf\x06\xa4\xf3\x87\x97\x56\xbb\x67\xee\xe6\x51\xdb\xf4\xff\x26\xbb\xbd\xf6\xfe\xf6\xe7\x5f\xc9\xbb\x4b\x3a\x21\xe9\x11\x7f\xb8\x47\xdc\xdb\xbe\x7f\x3f\x1d\x87\x66\xed\x80\x78\xf8\x72\xa3\xf7\x79\x96\x4f\x47\xe3\xff\x04\x00\x00\xff\xff\x78\x01\x9c\xc0\xa9\x1f\x00\x00")

func wwwJavascriptMapzenWhosonfirstCatalogJsBytes() ([]byte, error) {
	return bindataRead(
		_wwwJavascriptMapzenWhosonfirstCatalogJs,
		"www/javascript/mapzen.whosonfirst.catalog.js",
	)
}

func wwwJavascriptMapzenWhosonfirstCatalogJs() (*asset, error) {
	bytes, err := wwwJavascriptMapzenWhosonfirstCatalogJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/javascript/mapzen.whosonfirst.catalog.js", size: 8105, mode: os.FileMode(420), modTime: time.Unix(1506711247, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwJavascriptMapzenWhosonfirstGeojsonJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\xcf\x6e\x9b\x40\x10\xc6\xcf\xf8\x29\xa6\x27\x83\x82\xb0\x9d\xaa\xad\x14\x42\xa4\xaa\x52\x6f\x7d\x02\x84\xa2\x8d\x19\x60\x5b\x98\x45\xb0\x24\x4e\x1d\xbf\x7b\xb5\xff\x00\x1b\x5b\xbd\xb1\xf3\x7d\xfb\xdb\x99\xd9\x59\x5e\x59\x07\x0d\x6b\xff\x22\x41\xe2\x3e\x3e\x3e\xe0\x78\x8a\x57\x66\x15\xbd\x55\xa2\x17\x54\xf0\xae\x97\xa3\xe5\x2c\x68\xed\x57\xfc\x51\x89\xe2\x77\x2f\x14\xda\x2f\x06\xda\x4b\x2e\xc8\x0f\x8e\xab\x95\xa7\x8e\xed\xb1\x2e\x20\x01\xb5\xf4\xd6\x39\x76\xfc\x15\x9f\x5f\x5e\xc4\x61\xfd\x00\xa3\xd9\x02\xf4\x1e\xcf\xe3\x05\xb8\x48\xba\xd6\xd6\x2c\x38\x2a\xc1\xeb\x50\x0e\x1d\xc1\x85\x18\x2b\xed\xb4\xdc\x2a\xdf\x5b\x5c\x67\x90\x24\xb0\xfe\x89\x4c\x0e\x1d\xfe\x10\x75\x8d\xfa\xc8\xb5\x3d\x4b\xa7\x58\x18\xb5\x87\x64\x42\xbb\x98\xc5\x6b\xdf\x5e\x0c\xa4\xba\xe3\xb4\xa8\x46\x2a\x65\x15\x4f\xa4\xfe\xad\x66\xca\x31\x50\x8e\x05\x27\xcc\xe3\xb9\xa4\x5b\x74\x45\x22\xbc\xb9\x8b\x70\xb1\x4b\x6b\x85\xe8\xc0\x57\x06\x9e\x6c\x63\xe0\xf0\x68\x92\x8b\x81\xdf\xdd\xb9\xd2\x34\x40\xb5\x08\x12\x7d\x0b\xd1\xac\xfb\xbe\xab\x21\xe5\x59\x10\xcf\xfc\xcf\xae\x04\xe5\x4a\xb7\xb6\xfa\x51\xd2\xc9\x68\x69\x77\x26\xb9\x12\xb4\x74\x7f\x29\x4d\xbb\x3e\x67\xee\x30\x75\x55\xfe\x27\xd3\xb1\x40\x4d\x97\x6f\x8f\x7e\xb4\x31\x7b\xe7\x9e\xe7\x32\x32\xba\x45\x9f\x16\x18\x41\x13\x46\x90\xc1\x08\x3a\xc3\xe8\x3c\x8c\x7e\x1d\xa3\xcb\xb0\x18\x53\xd2\x93\x8d\x8d\x18\x57\xa9\xd1\x6f\x62\xc6\x6c\x4c\xf9\x4f\x36\x36\xc7\x98\x6c\xf4\xc7\x88\x99\xc1\xec\xb0\xa7\xa6\x1d\xa1\x29\x27\x34\xe9\x84\x76\x30\xe6\xb3\x8f\x75\x8f\xf0\x9f\x07\x30\x8e\xfd\x66\x03\xdf\x73\xd6\x4a\xcc\xa1\xe8\x44\x03\x95\x94\xed\xc3\x66\x53\xf2\x3e\xea\x25\xdb\xff\xc1\xc3\xbe\x62\x54\x62\xb4\x17\xcd\x86\x6d\x76\xdf\xee\xbf\x7c\xdd\x8d\x53\x59\xa2\x68\xe6\x8f\x45\xad\x51\x76\xef\xe7\x8f\x45\x74\xb9\x7d\x52\x4d\xa4\x57\x9c\x98\xc4\x7e\xf6\x5c\x6a\x26\x95\x23\xcd\x42\x1d\x02\x00\xa8\xa9\x34\xa1\xc5\xa0\x43\x02\xe3\xa8\x2b\x74\xba\xcd\xdc\x03\xd4\x53\x0f\xb6\xb5\x0a\x1a\xb5\x43\x5f\xf9\xa3\x2f\xe5\x59\xba\x53\x73\x6e\x0c\x54\x5e\x33\x6c\x9d\xe1\x34\x25\xd8\x70\x32\xb7\xfd\x8b\xc9\x2a\x6a\x38\x45\xac\x6d\xeb\x77\x9f\x86\xba\x0e\x75\xfa\xc1\x94\x7a\xc3\x0e\x73\x33\x3b\x2c\xcd\xf1\x19\x99\xca\x9b\x64\x2a\x2f\xc9\x33\xf3\x25\x59\x99\xe3\x8b\xa1\x31\x99\x87\xf6\x1c\x83\xf2\x6c\x8a\xa1\x03\x2e\xc6\xe7\xa8\x26\x50\xad\xd5\x9f\xde\xc1\xd4\xbf\x23\x5e\xad\x4e\x81\x1f\xc4\xab\x7f\x01\x00\x00\xff\xff\x2a\x82\x1a\x4f\x4a\x06\x00\x00")

func wwwJavascriptMapzenWhosonfirstGeojsonJsBytes() ([]byte, error) {
	return bindataRead(
		_wwwJavascriptMapzenWhosonfirstGeojsonJs,
		"www/javascript/mapzen.whosonfirst.geojson.js",
	)
}

func wwwJavascriptMapzenWhosonfirstGeojsonJs() (*asset, error) {
	bytes, err := wwwJavascriptMapzenWhosonfirstGeojsonJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/javascript/mapzen.whosonfirst.geojson.js", size: 1610, mode: os.FileMode(420), modTime: time.Unix(1506709909, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwJavascriptMapzenWhosonfirstRenderJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x56\x41\x6e\xdb\x3a\x10\x5d\x4b\xa7\x18\xe8\x2f\x22\xc1\xfe\x4a\xfe\x36\x8a\xf0\x11\x14\xdd\x76\xd5\x5d\x11\x04\xb4\x38\xae\x59\xd3\x92\x41\x8e\x63\xa7\x89\x81\xde\xa1\x37\xec\x49\x8a\x21\x29\x59\xb2\x65\x65\x91\x44\x12\xdf\x7b\x33\x9c\x79\x33\xc8\x8b\x30\xb0\x11\xdb\x9f\x58\x43\xd9\x3e\xbc\xbf\xc3\xdb\xb1\x88\xfd\x5b\xbe\x5f\x35\xb6\xa9\x97\xca\x58\xea\x20\x83\x8f\x01\x3e\x82\xcf\x0d\xd6\x12\x0d\x94\x90\x2e\x77\x75\x45\xaa\xa9\xd3\xec\x2d\x8e\x23\x8e\x6a\x51\x2f\xa1\x04\x7e\x8d\x6e\x3c\xf0\x59\x0a\x12\x37\xf7\xd0\x81\xe5\x1c\x2a\x3a\x38\x4a\xe4\x48\x7b\x23\xb6\x5b\xa7\x28\x9b\x6a\xb7\xc1\x9a\xf2\xca\xa0\x20\xfc\xac\x91\xdf\xd2\x44\xaa\x97\x24\x2b\x18\x1f\xb0\xb9\x45\x7a\x24\x32\x6a\xb1\x23\x4c\x93\x4a\x0b\x6b\x93\x39\x24\x72\xb7\xd9\xa2\x61\x6c\x2b\xee\x93\x70\x5c\xfe\x51\x4b\x48\x1f\x8d\x11\xaf\xb9\xb2\xee\x6f\x2a\xb3\xec\x8d\x4f\xa2\xee\x5e\x7c\x89\x70\xcb\x67\xad\x2c\xb5\x19\x3b\x91\x63\x1c\x01\x00\xf0\x23\x6a\x8b\xc0\x82\xf4\xba\xc5\x66\x99\xca\x0c\xca\x12\x92\x66\xf1\x03\x2b\x4a\xa6\x54\xa5\xaa\xce\x55\x3b\x41\x4f\xe3\xd4\xab\xa6\x26\xac\xe9\x8c\x4b\x78\xe8\x71\x87\x31\xae\xd5\xcf\x6e\x45\x9d\x64\x45\x14\x8a\x10\x08\x39\x97\xb2\x96\x9f\x56\x4a\xcb\x34\x04\xeb\xa7\xd3\xd6\xba\x0f\xf3\x4c\x8f\x32\x48\x3b\x53\xb7\xed\x2b\xbc\xf8\x71\x3e\xe8\xbd\xaa\x68\xb4\xf7\x5d\x11\xf9\xa2\x24\x16\x1a\x27\xd2\x77\xe7\xc1\x00\xee\xf9\x5a\xfb\xbd\x90\xfb\xfd\xaf\x44\x12\x4a\x5b\x47\x6b\x83\x2d\x1b\x03\xe9\x1a\x54\x0d\xd2\xb5\xa7\x2b\xb5\x69\xf6\x53\xf1\x4d\x08\xee\xb0\x5a\x2c\x50\xbb\x36\x40\x09\x6b\xff\xbd\x3b\x7c\xae\xe8\xc0\xa3\xc1\xb7\x84\xff\xf9\xb2\x30\x83\x24\x4f\x60\x06\x6b\xb8\xbf\x80\xaf\x50\x4c\x77\x8e\x56\x17\xa1\x2f\xd1\x5f\xf1\x40\x5f\x1a\x89\xe9\x29\xb5\x40\xf2\xfa\x83\x0e\x3a\x8c\x3b\x1e\x31\xda\xd5\x34\x64\x9b\x46\x47\x5a\x34\xf2\xf5\xdc\xd6\x82\x44\x2a\xbf\xad\x9f\xe6\xae\x0e\x81\x11\xd4\x07\x49\x30\xb9\x2f\x68\x9a\xfd\xe0\xdc\xe7\x1d\x10\xe7\x87\x03\xab\x7a\xbe\x77\xc5\xc0\xa8\xcd\xbe\x1b\x58\x0f\x0a\x76\x75\xd0\xc2\x3b\x35\xea\x59\x95\x27\xfd\x63\xab\x56\xcd\xce\x17\x2a\xd7\x58\x7f\xa7\xd5\x60\xb1\x84\xc3\x12\xee\xba\xe1\x77\x21\x2f\xc6\x37\xf9\xf3\xeb\x77\x32\xbe\x54\x4e\x3a\x0f\x25\xfc\x77\x5d\x27\xd4\xfa\xee\x69\x5c\xc6\xb9\x45\xd9\xa9\x9e\xee\xf4\xe5\x6c\x30\x4d\x95\x77\x05\x28\x78\xf0\x77\x2d\x40\xcd\x66\x21\x8d\xae\xf7\x8a\x70\x33\xa1\xac\x55\xdf\xb4\xd7\x8d\xa2\x7c\xf2\x3c\x21\xff\xf0\x84\xa8\x7e\x4b\x39\xc6\x35\xcf\xf0\xcd\x06\x67\x0c\x1e\x6f\x37\x43\x47\xba\xcd\x6d\x18\xef\xb6\x67\xbb\xb5\xe4\x47\x5c\x16\xd1\xed\xad\x4f\x7f\x45\x1b\x6d\xb7\x58\x29\xa1\xab\x95\x30\x36\x95\x59\xd1\x67\xf0\x9e\xfd\x78\x0d\x33\x9a\x1f\xcf\xb6\x98\x92\x7d\x4b\x8c\x00\x48\x91\xc6\x13\x86\x1b\x77\x0a\x3d\xb9\x17\x4e\x1b\xc1\xc9\xf6\x2b\x17\x76\x41\xbf\x64\x8c\x69\x4b\x16\xc5\x11\xff\x1b\xd0\x77\x60\x11\xc7\xc7\x2c\xcd\x8a\x28\xfe\x1b\x00\x00\xff\xff\xee\x07\x3a\xfe\x68\x08\x00\x00")

func wwwJavascriptMapzenWhosonfirstRenderJsBytes() ([]byte, error) {
	return bindataRead(
		_wwwJavascriptMapzenWhosonfirstRenderJs,
		"www/javascript/mapzen.whosonfirst.render.js",
	)
}

func wwwJavascriptMapzenWhosonfirstRenderJs() (*asset, error) {
	bytes, err := wwwJavascriptMapzenWhosonfirstRenderJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/javascript/mapzen.whosonfirst.render.js", size: 2152, mode: os.FileMode(420), modTime: time.Unix(1506645651, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwCssMapzenWhosonfirstCatalogCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\xdd\x8e\x9b\x30\x10\x85\xef\x79\x8a\xa9\xa2\xbd\x24\x62\xd3\x46\xaa\xcc\xd3\x38\xf6\x00\x56\xfc\x43\xc7\x43\x36\xdb\x2a\xef\x5e\x81\x21\x05\xc2\x6e\x7a\xb1\xca\x55\x66\x6c\x7f\x67\xe6\x1c\x76\x4a\xfa\x8b\x8c\xf0\x27\x03\x00\x70\x92\x6a\xe3\x05\xbc\xa2\x2b\xb3\x5b\x96\xed\x9c\x6c\xc7\x56\x83\xa6\x6e\x58\xc0\xf7\xa2\x68\xaf\xa9\xf9\xab\x43\x7a\x1f\xdb\xa7\x40\x1a\x49\xc0\xa1\xbd\x42\x0c\xd6\x68\xd8\x15\x45\x51\xce\xaf\x1e\x87\x8b\x09\x73\xcd\x37\x8a\x3d\x3b\x3f\x05\xe6\xe0\xc4\x5d\x41\x82\x54\x81\xdc\x48\x6a\xa5\xd6\xc6\xd7\x02\xf6\xc7\xfe\xcc\xac\x94\x5b\xac\x78\x26\x5e\x23\x4b\x63\xe3\x7f\x28\x0c\x17\xa4\xca\x86\x37\x01\x51\x51\xb0\x16\xbe\x19\xd7\x06\x62\xe9\x79\x31\xc2\x38\xfc\x6a\x86\x45\x75\x73\x08\xc2\xd8\x59\xfe\x4a\x25\xb0\x2d\x05\x9e\x6b\x89\xa1\x23\x85\x77\xc3\x57\x57\x3f\xe7\x3d\x74\x3f\x95\x7b\xcb\x32\x96\x27\x3b\xb1\xac\xf1\x38\xc1\x0e\x93\x75\x55\xf0\x9c\x57\xd2\x19\xfb\x2e\x20\x4a\x1f\xf3\x88\x64\xaa\xd4\x7c\x33\x9a\x1b\xf1\x5a\x14\x2f\xe5\x6c\x71\x79\x6c\xa5\x1a\x12\xd0\xf9\x88\x23\x87\x46\xc8\x05\x89\x8d\x92\x36\x97\xd6\xd4\x5e\x00\x87\xb6\x5c\xc6\x66\xb5\x9f\x54\xe8\xff\x0f\xef\x34\x93\x47\x52\x9d\x6b\x0a\x9d\xd7\xb9\x0a\x36\x90\xd8\xa1\xee\x7f\xe5\xc2\xc1\xe4\x1e\x37\xc6\xa7\x3a\xe3\x95\x47\x72\x9f\xc5\x47\xce\x86\x31\xfb\xc3\x2a\xc6\x02\xf6\x87\x23\xba\x41\x8e\x1e\xe5\x6c\x3e\xbc\x3c\xff\x01\x6c\x94\xfa\x4f\xe9\xda\xa2\xfb\x67\x92\xbc\x62\x82\x3b\x35\xed\xff\xc7\xcf\x97\xa7\xbe\x0f\x2f\xb5\x34\x59\x3d\xb8\x1a\xcd\x6f\x14\xd1\xc9\xbe\x3b\x4f\xd2\x71\x33\x49\xf3\x24\x3e\x9e\xb8\x65\xd9\x5e\x59\xa3\xce\xb9\x9b\x18\xc9\x16\x38\xd9\x0e\x67\xcb\xd7\xa8\x02\x49\x36\xc1\x8b\xce\x6b\xa4\x3e\x74\xe9\x3a\x12\x05\xfa\xc8\x5d\x68\x8d\x3f\x2f\x89\x7f\x03\x00\x00\xff\xff\x08\x5c\xb8\xab\x1a\x05\x00\x00")

func wwwCssMapzenWhosonfirstCatalogCssBytes() ([]byte, error) {
	return bindataRead(
		_wwwCssMapzenWhosonfirstCatalogCss,
		"www/css/mapzen.whosonfirst.catalog.css",
	)
}

func wwwCssMapzenWhosonfirstCatalogCss() (*asset, error) {
	bytes, err := wwwCssMapzenWhosonfirstCatalogCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/css/mapzen.whosonfirst.catalog.css", size: 1306, mode: os.FileMode(420), modTime: time.Unix(1506706910, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwCssMapzenWhosonfirstRenderCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xdf\x6a\xf3\x30\x0c\xc5\xef\xf3\x14\xfa\xf8\xd8\xa5\x4b\x57\xd8\x8d\xf7\x34\x4e\xac\x24\x62\xfe\x87\xac\x6c\x2d\xa3\xef\x3e\x12\xbb\xa1\x66\x2d\x8c\x5c\xe9\x28\x3a\xd2\xef\x78\x71\xf0\xdd\x01\x00\x38\xca\xa2\xb2\x5c\x1c\x2a\xb9\x24\xd4\x10\x62\xc0\xf7\xad\xe5\x0d\x4f\x14\xf4\x31\x9d\x4b\x9d\x8c\xb5\x14\xa6\x22\x5c\xbb\xee\x60\x17\x9f\x90\x41\x4c\xef\x70\xb7\x0b\xa8\x66\xa4\x69\x16\x7d\x42\x5f\x06\xc7\x18\x44\x8d\xc6\x93\xbb\x68\xc8\x26\x64\x95\x91\x69\x2c\xcd\x2f\xb2\x32\xeb\xd7\xe3\xf1\xa5\xd4\x7d\x64\x8b\xac\x72\x32\xc3\xba\x0c\x96\x90\x51\x1e\xec\x13\xae\x2b\x3f\x91\x85\x06\xe3\x94\x71\x34\x05\x0d\x12\x53\x7b\x2f\xec\x04\x95\x68\x13\xd6\xba\x71\x9d\xab\x5f\x6f\x86\x8f\x89\xe3\x12\xac\x1a\xa2\x8b\xac\xff\xa3\x5d\xbf\xe6\x3c\x89\xe9\x3e\xaa\xaa\xf6\x51\x24\xfa\xda\xd8\x16\xdc\x35\x1d\x8e\xa2\x21\x47\x47\x16\x64\xa6\x00\xff\xc8\xa7\xc8\x62\x82\x34\x26\xbc\x85\xf7\xe4\xc7\xdd\x54\xf0\x2c\x95\x78\x35\xfe\xcd\xd7\x46\x70\x38\xbd\xa1\x6f\x71\xf5\x48\x9c\x45\x0d\x33\x39\x7b\x43\xbf\x83\x7b\x72\x68\x6b\xe1\xcc\x63\x87\x5b\x10\x7f\x30\x29\x8f\x79\x9b\x7f\x88\xd5\x52\x34\x5a\x4d\xf5\xb0\xeb\x4d\x04\xd7\xee\x27\x00\x00\xff\xff\xcd\x95\x93\x60\xe8\x02\x00\x00")

func wwwCssMapzenWhosonfirstRenderCssBytes() ([]byte, error) {
	return bindataRead(
		_wwwCssMapzenWhosonfirstRenderCss,
		"www/css/mapzen.whosonfirst.render.css",
	)
}

func wwwCssMapzenWhosonfirstRenderCss() (*asset, error) {
	bytes, err := wwwCssMapzenWhosonfirstRenderCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/css/mapzen.whosonfirst.render.css", size: 744, mode: os.FileMode(420), modTime: time.Unix(1506645651, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"www/index.html":                                    wwwIndexHtml,
	"www/javascript/mapzen.whosonfirst.catalog.init.js": wwwJavascriptMapzenWhosonfirstCatalogInitJs,
	"www/javascript/mapzen.whosonfirst.catalog.js":      wwwJavascriptMapzenWhosonfirstCatalogJs,
	"www/javascript/mapzen.whosonfirst.geojson.js":      wwwJavascriptMapzenWhosonfirstGeojsonJs,
	"www/javascript/mapzen.whosonfirst.render.js":       wwwJavascriptMapzenWhosonfirstRenderJs,
	"www/css/mapzen.whosonfirst.catalog.css":            wwwCssMapzenWhosonfirstCatalogCss,
	"www/css/mapzen.whosonfirst.render.css":             wwwCssMapzenWhosonfirstRenderCss,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"www": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"mapzen.whosonfirst.catalog.css": &bintree{wwwCssMapzenWhosonfirstCatalogCss, map[string]*bintree{}},
			"mapzen.whosonfirst.render.css":  &bintree{wwwCssMapzenWhosonfirstRenderCss, map[string]*bintree{}},
		}},
		"index.html": &bintree{wwwIndexHtml, map[string]*bintree{}},
		"javascript": &bintree{nil, map[string]*bintree{
			"mapzen.whosonfirst.catalog.init.js": &bintree{wwwJavascriptMapzenWhosonfirstCatalogInitJs, map[string]*bintree{}},
			"mapzen.whosonfirst.catalog.js":      &bintree{wwwJavascriptMapzenWhosonfirstCatalogJs, map[string]*bintree{}},
			"mapzen.whosonfirst.geojson.js":      &bintree{wwwJavascriptMapzenWhosonfirstGeojsonJs, map[string]*bintree{}},
			"mapzen.whosonfirst.render.js":       &bintree{wwwJavascriptMapzenWhosonfirstRenderJs, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
