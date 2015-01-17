package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _templates_404_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x34\x8e\x31\x0e\xc2\x30\x0c\x45\x67\x38\x85\xc9\xc2\x54\x32\xb0\xba\xb9\x41\x25\x06\x2e\xe0\x82\xdb\x5a\x84\x24\x72\x4a\xa5\xde\x1e\xb7\x88\xc1\x1e\xbe\xfd\xfc\x8c\x4f\x59\xe0\x11\xa9\xd6\xd6\x51\x64\x9d\x61\xef\x8d\xa4\x21\xbb\x70\x3c\xe0\x74\x0d\x37\x1a\x19\x3a\xa9\x55\xd2\x78\x42\x6f\x89\xe5\x25\xdc\x27\x86\xb2\x8d\xd6\xfc\x01\x52\x86\x98\xf3\xcb\x56\x60\xc8\x0a\x52\xd3\xd9\x6e\x2d\x24\x91\xfa\xc8\x17\xf4\xe5\x47\x21\xfd\x75\xfd\x9c\xc0\xaa\x29\x2a\x6f\xd2\xd5\xc1\xa4\x3c\xb4\xce\xbb\xd0\x91\x24\xd8\xac\xe8\x29\xec\x24\x7a\xfb\x33\x7c\x03\x00\x00\xff\xff\x83\x9f\x13\x5f\xad\x00\x00\x00")

func templates_404_html_bytes() ([]byte, error) {
	return bindata_read(
		_templates_404_html,
		"templates/404.html",
	)
}

func templates_404_html() (*asset, error) {
	bytes, err := templates_404_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/404.html", size: 173, mode: os.FileMode(420), modTime: time.Unix(1421529581, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_error_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x49\xc9\x2c\x53\x48\xce\x49\x2c\x2e\xb6\x55\x4a\xcc\x49\x2d\x2a\x51\x00\x93\xba\x29\x89\x79\xe9\xa9\x45\x4a\x76\x5c\x9c\x36\x19\xc6\x76\xae\x45\x45\xf9\x45\x8a\x36\xfa\x40\x26\x50\xa0\xc0\x2e\x24\x23\xb5\x28\x55\xa1\x3c\xb1\x58\x21\x51\xa1\xa0\x28\x3f\x29\x27\x35\x57\x21\x3f\x4f\xa1\x24\x23\x55\xa1\x38\xb5\xa8\x2c\x15\xa4\xb6\x00\xac\x14\x66\x76\x71\x6e\x62\x4e\x8e\x12\xc4\x20\x2b\x85\xea\x6a\xbd\xda\x5a\xb0\x12\x1b\x7d\xa0\x03\xec\x00\x01\x00\x00\xff\xff\xf5\xff\x98\x03\x86\x00\x00\x00")

func templates_error_html_bytes() ([]byte, error) {
	return bindata_read(
		_templates_error_html,
		"templates/error.html",
	)
}

func templates_error_html() (*asset, error) {
	bytes, err := templates_error_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/error.html", size: 134, mode: os.FileMode(420), modTime: time.Unix(1421529581, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x94\xbd\x6e\xdb\x30\x10\xc7\xe7\xf4\x29\xae\xdc\x15\xa1\x6d\x50\x64\x90\xb5\x14\xdd\x8a\x36\x48\x02\x14\x1d\x29\xeb\x64\x11\xa1\x78\x2c\x79\xb4\x63\x04\x79\xf7\xf2\x43\x32\x6c\xa4\x35\x60\xa0\x83\xc5\xd3\x7d\x91\xfc\xe9\x7f\x6e\x7a\xb5\x85\xb5\x96\xde\xaf\x84\xa3\x9d\x68\xdf\x5d\x1d\xbb\xd6\xa4\xab\x67\x5f\x7d\xf8\x08\xc9\xf2\xd3\x62\x4d\x7d\xf5\x79\x31\x68\x18\x3c\x72\xf5\x29\xbf\xeb\xcd\x1c\x88\xc6\x12\x48\x4d\xaf\x9a\x81\xdc\x04\x72\xcd\x8a\xcc\x4a\xd4\xd2\xaa\xda\xa2\xf3\x64\x04\x4c\xc8\x23\xf5\x2b\x71\xf7\xe3\xe1\x31\xe7\x9e\x1c\x21\xd5\x55\x1b\x47\xc1\x96\xd8\x55\xa3\x65\x87\x1a\xa2\x7f\x25\x8c\x9c\x50\xb4\xdf\xe3\xb3\xa9\xb3\x7b\x4e\x51\xc6\x06\x06\xde\x5b\x5c\x09\xc6\x67\x16\x27\xdd\xd6\x64\xd8\x91\x16\x90\xca\xe7\x26\x60\xb5\x5c\xe3\x48\xba\xc7\xd8\xf7\x17\x05\x07\x83\x72\x9e\x41\x9a\x1e\x62\x2d\x43\x49\x73\xf8\x3b\x28\x87\xfd\xbc\x91\x5d\x1a\x8f\xa8\x6d\xd5\x69\x5a\x3f\x89\x36\x57\xa7\x74\xd8\x29\xad\xa1\x43\xf0\x23\xed\x0c\x30\xc5\x6e\x7b\x32\x08\x3c\x4a\x86\xad\xf2\x8a\x7d\xb4\x95\x87\x1d\x76\xf1\x05\xa1\x0f\x4e\x99\x4d\xf4\x21\xe0\x16\x0d\x37\xb5\x2d\x44\xea\x88\xe4\x32\x36\x38\x49\xa5\x45\xfb\x35\x2d\xff\xa6\x53\xb2\xce\xe1\x99\x33\xde\xf2\xc9\x01\x90\x7d\xef\xd0\x7b\x71\x0e\x48\xbe\xe2\xa0\x50\xf7\x10\x0d\xb2\x49\x04\x52\x5f\xc3\xcf\x99\x10\x19\xbd\x87\xe0\xb1\xb0\x98\x3b\x26\x5c\xe9\x28\x51\x33\xb0\xa7\x00\xb2\xa3\x74\xea\x94\x91\xd1\xe4\x2f\x33\x04\x0e\x6e\x66\xe5\xaf\x2f\xa5\xd5\x05\x66\x32\x4b\xbc\x63\x03\xf1\x57\x59\xa7\x26\xe9\xf6\xd9\x2e\x17\x98\x59\xf9\xd0\x4d\x8a\x45\x7b\x8f\x1b\xe5\x19\x5d\x53\x97\x06\xa7\x7b\x36\x75\xda\x28\x0d\x52\xf1\x2c\xcb\xe8\xe2\xe3\xc2\x71\xbb\x5d\x86\xec\x76\x19\xaa\xdb\x32\x4d\x47\x45\x56\x9a\xf8\xc9\xf3\xb3\xea\x71\x90\x41\xf3\xdb\x29\x2a\xe1\x11\x65\x1f\xe5\x25\xda\x2f\xc1\xb9\x84\xf0\x0e\xc9\x6a\x3c\xc2\x15\xf4\x52\xa1\xe3\x15\x67\x5a\xa0\xe2\x74\x7a\x96\x1b\x3c\xc8\x4c\xbd\x4d\xab\xa2\x7e\xa7\x2c\xfe\x83\xec\x93\x8c\xf3\x14\xbd\x8f\x02\x54\x33\xa6\xa0\x0b\xa5\xb2\xe7\x61\x3d\x87\xe1\x66\xc1\x70\xb3\x60\xb8\xf9\x2f\x18\x1e\xd0\x6d\xd1\xc1\x03\x4b\xf6\x7f\xd7\x4c\x29\xe8\xa8\xdf\x1f\x04\xde\x7e\x4b\xff\x06\xf7\x38\x44\x91\x8e\xf0\x48\xf4\x04\x4d\x97\x11\xb1\x9a\xd0\x89\xb6\x8a\xba\x68\xcf\x4b\xb1\xb4\x1d\x88\x38\x15\x24\x64\x32\xaa\x38\x4a\xd1\x97\x13\xbd\xbc\x5c\xbf\xbe\x1e\x29\xea\x84\xd5\xbc\xfc\x09\x00\x00\xff\xff\xc6\x49\xe2\x1c\xbb\x05\x00\x00")

func templates_index_html_bytes() ([]byte, error) {
	return bindata_read(
		_templates_index_html,
		"templates/index.html",
	)
}

func templates_index_html() (*asset, error) {
	bytes, err := templates_index_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/index.html", size: 1467, mode: os.FileMode(420), modTime: time.Unix(1421529581, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_layout_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x57\x5d\x6f\xdb\xbc\x15\xbe\xb6\x7f\x05\xa3\x05\xa8\x0c\x44\x52\x9c\xee\xc5\xba\xc4\x32\xf0\x36\x0d\xb6\x16\x5d\xdb\x35\x19\xba\x22\xc8\x05\x2d\x1e\x59\x4c\x28\x51\x25\x29\x27\x5e\x9a\xff\xbe\x43\x4a\x94\x14\x27\xd9\x56\x60\x17\xad\x29\xf2\x7c\x9f\xe7\x7c\x64\xb1\xf7\xee\xf3\xe9\xc5\xf7\x2f\x67\xa4\x30\xa5\x58\x4e\x17\xf6\x87\x08\x5a\xad\xd3\x00\xaa\x60\x39\x9d\x2c\x0a\xa0\x0c\x7f\x27\x8b\x12\x0c\x25\x59\x41\x95\x06\x93\x06\x8d\xc9\xa3\x37\xc1\xf0\x50\x18\x53\x47\xf0\xa3\xe1\x9b\x34\xf8\x67\xf4\x8f\xdf\xa3\x53\x59\xd6\xd4\xf0\x95\x80\x80\x64\xb2\x32\x50\x21\xd7\xfb\xb3\x14\xd8\x1a\x46\x7c\x15\x2d\x21\x0d\x36\x1c\x6e\x6b\xa9\xcc\x88\xf4\x96\x33\x53\xa4\x0c\x36\x3c\x83\xc8\x7d\x1c\x10\x5e\x71\xc3\xa9\x88\x74\x46\x05\xa4\xf3\x56\x8c\xe1\x46\xc0\xf2\x6c\x83\x5c\xe4\x2b\xac\xb9\x36\x0a\xf5\xca\x6a\x91\xb4\x2f\x53\x4b\xb4\x17\x45\xe4\xad\x94\xc6\x3e\xd6\xe4\xf4\xfc\x9c\x44\x91\xe3\x16\xbc\xba\x21\x85\x82\x3c\x0d\x92\xa4\x02\xc3\x2a\x1a\xaf\x3c\x61\xc6\xaa\x38\x93\x65\xd2\x5f\x24\xaf\xe3\x79\x3c\x4f\x32\xad\x87\xbb\xb8\xe4\x48\xa5\x75\x40\x14\x88\x34\xd0\x66\x2b\x40\x17\x00\x26\x18\x14\xec\xbe\xf4\x1a\x51\xc3\xb5\x8e\x33\x21\x1b\x96\x0b\xaa\xc0\xa9\xa3\xd7\xf4\x2e\x11\x7c\xd5\x29\xb9\xa5\x26\x2b\x50\x33\xea\x4e\x18\x55\x37\x62\xfb\x8c\xf2\xc1\xcd\xbf\x5e\xfc\xed\xe3\x6f\xe4\xbc\xe0\x25\xa1\x15\xc3\x90\xe8\x5a\x56\x2c\xbe\xd6\xe4\xfd\xd9\x1b\xa2\x9b\xda\x06\x9a\xc8\xbc\x23\x04\x01\x25\x86\x4e\x3b\xe2\x12\x18\xa7\xe4\x47\x03\x8a\x83\xf6\x21\xb2\x42\xbf\xfd\xfe\xf5\xd3\xfb\x4f\x7f\x39\x1e\x8b\x63\x12\x74\xf5\xca\x90\x5b\xa9\x6e\x08\xcf\xc9\x56\x36\xc4\x26\x92\x98\x02\x48\x4d\xd7\x80\x5f\x94\xe4\x5c\xc0\x71\x92\x8c\x84\x5d\x22\xad\x30\x68\x0d\xf9\xf3\x95\xbd\x9b\x2c\x74\xa6\x78\x6d\x88\x56\x59\x1a\x58\x20\x69\x64\x90\x5a\xc7\x25\xbd\xf3\x29\x70\xe1\xb0\xe8\xfc\x4d\x17\x7c\x83\xd1\xf8\x53\x7c\x38\x7c\xa3\x39\xc1\x72\x91\xb4\x72\x7e\x4d\xa6\xea\x1d\x4a\xe6\xf1\x1f\xe3\xa3\xfe\xc2\x46\x76\x57\xec\x62\xef\x12\x2a\xc6\xf3\x2b\xe7\xcd\x22\xe9\x8a\x63\xb1\x92\x6c\xeb\xde\x8b\x39\xc9\x04\xd5\x3a\x0d\x0c\xdc\x99\x28\xc3\xc8\x82\x0a\x96\xdf\x40\xa0\x42\x20\x46\x12\xd9\x28\x02\x16\xac\xc8\x3d\x77\x3c\x8c\x6f\x3c\x93\x45\x3f\xe5\x95\x65\x71\x4e\x8c\x9e\x94\xbc\x6d\x2f\x77\x18\x44\x74\xa7\xa3\xf9\x11\xb1\x27\x5d\xfa\x53\xc9\xa2\x37\xee\x20\xd6\xdd\x01\x6f\x64\x9e\x63\xed\x46\x47\xfe\xc1\x7f\x77\x72\x5b\xc1\x9c\x61\xe9\x2b\x25\x55\xe0\x75\x14\x9c\x31\xa8\x08\x96\x1c\xe2\xc6\xfd\x1f\x31\xec\x10\xde\x48\xcb\x58\xbc\x5e\x9e\x59\x1e\xf4\xe9\x75\x7f\x59\x0f\xb2\xa2\x52\xaf\x6d\x24\xeb\xfe\x71\xd5\x18\x23\x2b\xaf\x62\x65\x2a\x82\xff\x22\x06\x39\x6d\x04\x56\xce\xa9\x00\x8a\xd2\x5a\x2a\x6f\x5e\x82\xf6\x75\x21\x18\x8e\xf7\xf7\x64\xcb\x41\x30\xf2\xf0\x30\x1d\xbf\x74\x07\x0f\xe0\xeb\xbf\x23\xa8\xb7\x1e\x85\x63\x70\x60\x11\x4a\x06\xf1\xb5\x45\xfd\xd6\xe1\xa2\x3b\x3e\x4d\xfe\xa3\x26\xf2\x81\x6e\xe8\x79\x2b\xe7\x59\xa9\xff\x6b\x33\xb9\xde\xed\x25\x4f\xf4\x8e\x60\x1d\xe6\x4d\x95\xd9\xee\x16\xee\xcf\xee\x9d\xff\x1b\xaa\x88\xe1\x25\x28\x92\x92\xf6\x66\x62\xdb\xe4\x31\x39\x3c\x68\xbf\x58\xd3\xf6\xc3\xe1\x46\x1b\xaa\x90\xa0\x17\x35\xf3\x8c\x13\x53\x70\x1d\x5b\x76\x14\x56\x61\x25\xbf\xa3\x06\xc2\x59\xbc\x06\x73\x81\x2a\xc2\xd9\x49\x4b\xf7\xd0\x0b\x92\xf5\xcb\x72\xbc\xe2\x17\x64\x91\x88\xf4\xea\xbc\x5c\xf7\xf3\x70\x32\xed\x3d\xc3\x62\xb4\xc6\xbe\x6d\xd1\x92\x92\xfd\x30\xe8\xa0\xb3\x0c\x66\x31\x65\xec\xd4\x02\x28\xec\x01\x54\x2b\x5e\x52\xb5\xc5\x37\x5b\x7f\x61\xd0\xf1\xe3\x37\xda\x17\x64\x82\x67\x37\xc1\xc1\x60\x31\x74\x41\x9c\x64\x8d\x52\x58\x93\xbd\x87\xfb\xa1\xb5\x6d\x16\xd7\xb4\xbd\x8e\x15\x94\x72\x03\x5e\xdb\xda\x57\xc0\x2e\x7d\x4b\xe6\xc5\x3c\xcc\x46\xae\x40\x59\x9b\xed\x47\x1c\x4f\x9d\x1b\x82\x3f\x76\x41\xe0\x53\xb4\x56\xb2\xa9\x23\x6e\xa0\xec\x5d\xf8\x8e\x6d\xf5\x96\x0b\x41\x56\xe0\x3a\x6b\xce\x95\x36\x7b\x81\x17\xed\x7d\x21\xbd\x0b\x3e\x0b\x0e\x15\xb1\xf3\x3f\x9c\x4d\x3b\x3b\x63\x3b\x5c\x42\x9f\xa7\x46\x89\x63\x12\x24\xb4\xe6\x49\x0d\xb2\xc6\x59\xdd\x65\x76\xc2\xa8\xa1\x17\xdb\x1a\xf0\xf9\x5a\xcb\x2a\xe8\x12\x34\xb3\x10\x46\x3a\x4c\xa5\xb5\xd7\xfb\xb9\x63\x8a\x7d\x0a\x31\xf4\xbd\x29\xd6\x7f\x2b\xf1\xc0\xcf\x77\x8c\xc1\xe5\x55\xe7\x41\x6f\xa8\xac\x43\xef\xd5\x04\x67\x84\x95\x60\xcd\x37\x8d\x26\x7b\x29\x39\x3a\x3c\x1c\x00\x86\x01\xfc\x83\x63\xeb\xc3\x14\x05\xa3\x6c\xf5\x87\x11\x42\x86\x9c\xc5\xb4\xae\xb1\x8d\x87\x8f\xc0\xe5\x33\x39\x51\x60\x1a\x55\x79\x40\x4e\xfb\x68\xa0\xc9\xd6\xa0\x76\x3c\x68\xf8\x70\xfe\xf9\xd3\xd8\xd8\x3d\x47\xf3\xf3\x27\x71\x87\xb8\xa0\xfa\xf3\x6d\xf5\x45\xc9\x1a\xfb\xe5\x36\x0c\xba\xf0\xce\xfe\xff\x2e\xf4\x16\x5d\xa0\x90\xff\xe8\xc5\x7e\x0c\x34\x2b\x42\x67\x5f\x6b\xcf\xa8\x14\x38\xae\x57\x08\x3b\x5f\x10\x3e\x65\x30\xae\x5e\x4b\x10\xdb\xcb\xd9\x88\x28\x97\xaa\xa4\xc6\x00\xeb\x70\xad\x6b\xba\x53\x9c\x75\x23\x44\xa4\xf8\xba\xc0\xde\x58\x52\x21\xbc\xc3\x56\x52\x6c\xe4\x47\x69\xd7\xb9\x73\xa3\x78\xb5\x0e\x67\xbd\x07\xae\x66\xc4\xf3\xc5\xf2\x6a\xa7\x58\x5e\xcd\x62\xbb\x0b\x84\x97\xce\x42\xbb\x53\x1e\x0c\x76\x5d\xf5\x22\x3b\xf0\xc5\x75\xa3\x8b\x10\xc4\xcc\x83\x7a\x9c\x47\x4f\x23\xa0\x5a\x9b\x82\x2c\x09\xa2\xce\xa6\x0a\xa1\x82\x8b\x6b\xa7\xa7\x23\xf2\x82\x9f\xe1\x4b\xd3\x67\x19\xfb\x26\x30\x34\x8e\x1d\x14\xb4\xc5\xe0\x7b\x67\x6f\x9a\x6e\xdb\xa6\x6c\x4c\xd8\x15\xfa\x81\xad\x88\xc3\x17\x2a\x50\x37\xab\x92\x1b\xec\x6f\x1e\x6e\x10\xd7\xca\xad\x1d\xef\xda\x11\x1b\x8e\x0d\x68\xe7\xfd\x38\x65\xed\xcc\x0f\xfa\xb6\x61\x83\xe9\x52\xe1\xda\xdc\xaf\x36\x13\xdc\xf4\x0b\xc9\xf0\xb1\x6e\xcc\xa3\x0e\x73\x4c\x9c\xe4\x58\xe3\xc2\x49\x05\xff\x17\x76\xce\xa7\x8d\xa6\x87\xa8\xed\x28\x7d\xf5\x58\xbe\xcb\xc3\x2b\x8b\x7e\x18\xbc\xb1\xee\xf0\x0a\xd5\x1c\xbb\x36\x89\x3e\xe5\x32\x6b\xf4\xf0\xfe\x5c\x5f\x99\xb7\x89\x1a\x36\x95\x2e\x15\x4f\x8a\xeb\x85\x71\xe0\x83\x75\x32\xe0\x69\x94\x13\xdc\x7d\x31\x21\x4c\x96\x44\x48\xca\xc8\x8a\xe3\x9e\xad\x4d\x93\xe7\x6e\xe3\x76\x1d\xc8\xf5\x75\x21\x65\x3d\x6d\x5d\x18\x26\xea\x7d\x9f\x24\xeb\x6f\x37\xc6\xda\xe4\xe2\x1c\x6b\x0f\x4f\x32\x49\xda\x09\xf9\xd2\xd0\x9b\x8d\x3b\xd0\xcb\xa9\x3f\xd9\xad\x8e\xff\x12\xda\x9d\x11\xda\x45\xe1\x21\x6c\x37\xaf\xb6\xae\x47\x6b\x0d\x2e\x77\x6e\x69\x5e\x24\xed\x9f\x9e\xff\x0e\x00\x00\xff\xff\xdd\xb4\x13\xf0\x8b\x0e\x00\x00")

func templates_layout_html_bytes() ([]byte, error) {
	return bindata_read(
		_templates_layout_html,
		"templates/layout.html",
	)
}

func templates_layout_html() (*asset, error) {
	bytes, err := templates_layout_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/layout.html", size: 3723, mode: os.FileMode(420), modTime: time.Unix(1421530765, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"templates/404.html": templates_404_html,
	"templates/error.html": templates_error_html,
	"templates/index.html": templates_index_html,
	"templates/layout.html": templates_layout_html,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"404.html": &_bintree_t{templates_404_html, map[string]*_bintree_t{
		}},
		"error.html": &_bintree_t{templates_error_html, map[string]*_bintree_t{
		}},
		"index.html": &_bintree_t{templates_index_html, map[string]*_bintree_t{
		}},
		"layout.html": &_bintree_t{templates_layout_html, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
