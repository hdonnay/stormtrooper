// Code generated by go-bindata.
// sources:
// bootstrap/cmds.go
// bootstrap/detect.go
// bootstrap/detect_linux.go
// bootstrap/list.go
// bootstrap/main.go
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
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

var _cmdsGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x52\x4d\x8b\x14\x31\x10\x3d\x77\x7e\x45\x19\x18\x48\xcb\x90\x15\x8f\xca\x5c\x76\x60\x05\x3d\x28\x78\x54\x91\x38\x5d\xe9\x0e\xdb\x9d\x6a\x93\x6a\xd7\x45\xe6\xbf\x5b\x49\xcf\x34\xe8\x69\xf7\x92\xaf\x97\xf7\xea\xd5\x4b\x66\x77\xba\x77\x3d\xc2\xe4\x42\x54\x2a\x4c\x33\x25\x06\xa3\x1a\xed\x27\xd6\x32\x05\x2a\xe3\x48\x7d\x99\x22\xf2\xcd\xc0\x3c\x5f\xd7\x4b\x1a\xcb\x92\x72\x19\x67\xc7\xc3\x75\xbe\xf1\x61\xc4\xf5\xa0\x55\xca\x2f\xf1\x04\x0f\x3d\xb2\x71\xa9\xcf\xf0\xe5\x5b\xe6\x14\x62\xdf\xc2\x1f\xd5\x78\x4a\x10\xf6\xe0\xe1\xcd\x01\x92\x8b\xe2\xa4\xde\x11\xa4\x09\x1e\x5e\xfe\xc2\xf4\x83\x32\xd6\x7d\x23\x96\xec\xdd\x2c\x54\xf6\x86\xb2\xfd\xcc\x1d\xa6\xb4\x07\xed\x98\x71\x9a\x59\x24\x81\x09\x3c\xf2\x69\x80\xdd\xcf\xaf\x51\x8b\x6e\x2b\xc4\xb3\x92\x61\xd9\x83\xdc\x2e\x65\xc4\xb5\xfd\xe4\x52\x46\x53\x51\x29\x53\x80\x17\x07\x88\x61\x5c\x0b\x49\xbb\xf6\xce\xb1\x1b\x8d\x20\x9b\x82\x8f\x85\x5d\xba\xb2\xb7\x4e\xd8\x8b\xa8\xf0\x70\x91\x10\xf0\x70\x00\xad\x2f\x4e\x65\x07\x3a\xc4\x0e\x7f\xdb\x81\xa7\x92\x92\x68\x34\x4d\xc2\xbc\xd9\x28\x41\xda\x77\x12\xca\x33\x5c\xd4\x7b\x22\x22\xbd\x3b\x5e\xf2\x91\x3a\x2c\x94\x2a\xb5\x1e\x7d\xfc\xf0\x3f\x59\xf7\xc4\x90\x2b\x08\xa7\x42\xd8\x75\x12\xcc\xbf\x22\x57\xf5\x0e\x3d\xa6\x8a\xdd\x52\xf7\x68\x8f\xa3\x64\x6f\x0a\x48\x9b\x6d\x09\xfe\x98\xd0\xb1\xa4\x77\x79\x63\xfb\x9e\x42\x34\x0f\x94\xee\x25\xef\xd8\x3e\xa7\x9b\xb5\x1e\x6d\x85\x56\xea\xf7\xad\x58\x10\x88\xe6\x47\x43\xfb\xcd\x54\xfb\xf6\x89\xcf\xf5\xe4\xef\x53\xfa\x80\xdd\xab\xd7\x1d\x38\x86\x5d\xae\xff\x26\xd4\x56\x56\x97\x67\x75\x56\x7f\x03\x00\x00\xff\xff\xcc\xf3\x80\x88\x27\x03\x00\x00")

func cmdsGoBytes() ([]byte, error) {
	return bindataRead(
		_cmdsGo,
		"cmds.go",
	)
}

func cmdsGo() (*asset, error) {
	bytes, err := cmdsGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmds.go", size: 807, mode: os.FileMode(420), modTime: time.Unix(1442345366, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _detectGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x24\xcd\xb1\x4e\x03\x41\x0c\x04\xd0\x9a\xfd\x8a\xa1\x22\x11\x90\x7c\x05\x05\x15\x05\x25\xa2\xd8\x1c\xbe\x3b\x8b\x8b\xbd\x78\xed\x08\x14\xdd\xbf\xe3\x83\xca\x9a\x91\xf5\xe6\x78\xc4\xfd\x29\x78\xf9\xc0\xed\xc2\x12\xdf\xa5\x64\xf3\xf2\x0a\xee\xa8\x98\x48\xc8\x78\xc0\xac\xfa\x09\x57\x34\x6d\xb1\x54\x27\x54\xf9\xf1\x99\x65\x42\x08\x7f\x45\xe6\x93\x86\xc3\x67\x82\x57\x9b\xc8\x53\x38\xe0\x79\xdc\xac\x2c\x8d\xee\x3a\x44\x3d\xc5\xb3\x1a\xa1\x37\x1a\x78\x4c\xf7\x42\xd6\x59\xe5\x21\x9f\x72\xef\x6f\x52\xf4\x51\xdb\xa1\x8c\x21\x43\x22\xbb\x3d\x9e\xe4\x82\x6b\xb9\x31\xf2\x30\xd9\xd2\xee\x5c\xdb\x5b\x77\xcb\xf9\xf7\xff\x73\x5d\xf7\x65\x2d\xbf\x01\x00\x00\xff\xff\x04\x84\x13\x91\xca\x00\x00\x00")

func detectGoBytes() ([]byte, error) {
	return bindataRead(
		_detectGo,
		"detect.go",
	)
}

func detectGo() (*asset, error) {
	bytes, err := detectGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "detect.go", size: 202, mode: os.FileMode(420), modTime: time.Unix(1442257113, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _detect_linuxGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x52\x3b\x6f\xdb\x30\x10\x9e\xc5\x5f\x71\xe5\x12\x19\x55\xe5\xcc\x06\x3c\xf4\xe1\x21\x43\x63\xc0\x0e\xba\xa4\x41\x40\x49\xa7\xf4\x50\xbe\xc0\x87\x6d\xa1\xc8\x7f\xef\x51\x6e\x0d\x77\x69\x87\x2e\xd4\xf1\xe3\x7d\x0f\x1e\xb5\x5c\xc2\xdb\x2e\x93\x1e\x40\x93\xcd\x27\x21\xbc\xea\xbf\xab\x17\x04\xa3\xc8\x0a\x41\xc6\xbb\x90\xa0\x16\x95\xec\xa6\x84\x51\x72\xe1\xe2\x12\x4f\xd8\x97\x32\xa6\x40\xf6\x85\xd1\x85\x10\xac\xb4\xdd\x83\x77\x3e\x6b\xc5\x9d\x90\x8e\x0e\xd0\x1e\xc8\x06\x67\xd0\x26\x38\xa8\x40\xaa\xd3\x18\x57\xf0\xe9\x6e\xff\xb0\xdb\x3e\xdf\xbf\xff\xbc\x01\x65\x87\xdf\xfb\x2f\x9b\x5d\x0b\x77\x63\x51\xca\xb6\xb4\x42\x72\x30\x60\xc2\x60\xc8\xf2\xe6\x1b\x45\x20\x3b\xba\x60\x54\x22\x67\x1b\x46\x70\x82\x23\x69\x0d\x1d\x42\xc4\x54\xfa\x19\x03\x34\x3e\x4d\x70\x0e\x57\xd4\x8a\xc7\x6d\x03\x01\xa3\xc7\x3e\xd1\x01\xf5\xd4\x8a\x31\xdb\x9e\x13\xd7\x0b\xd8\xd8\x03\xfc\x10\xd5\x00\xab\x75\xa9\x6b\xa3\xfc\xe3\x99\xfc\x74\xfe\xf0\x61\x25\xaf\x42\xcb\x15\x48\xd9\x5c\x81\x9c\x9c\x31\x90\xb7\x05\x7d\x5d\x88\xaa\x37\xb3\x5a\x99\x53\xfb\xd1\x19\xc3\x09\x6a\xa9\x63\xf7\x1c\x50\xa3\x8a\x28\x1b\x90\xef\x68\x5e\xc3\xbc\x46\xc9\x2c\x97\x53\x03\x18\x42\xa1\xb2\x42\xbb\xcd\xc9\xe7\x54\xf3\x09\x8d\x33\xfe\x66\x0d\x96\x74\x09\x5b\x05\x4c\x39\x58\x18\xd8\x4f\x54\x33\x63\x7e\xa0\xf6\x1e\x8f\x1f\xf2\x38\x62\xa8\x59\x8d\x9f\xa5\x3a\x5d\x24\x43\xbb\x43\x35\xec\xe7\x2b\xd5\x37\x5f\xed\xcd\xbf\x85\x87\xc7\x3f\xee\xfd\x04\xeb\x5f\x63\x8d\xed\x43\x20\xb3\xe7\xbf\x05\xeb\xd3\x95\xcd\x7f\xba\x94\x41\xfe\xc5\xe4\xc2\x7a\x15\x3f\x03\x00\x00\xff\xff\x34\x78\x79\xbd\xb9\x02\x00\x00")

func detect_linuxGoBytes() ([]byte, error) {
	return bindataRead(
		_detect_linuxGo,
		"detect_linux.go",
	)
}

func detect_linuxGo() (*asset, error) {
	bytes, err := detect_linuxGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "detect_linux.go", size: 697, mode: os.FileMode(420), modTime: time.Unix(1442344304, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _listGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x57\x70\x0c\x0d\xf1\x77\x77\xf5\x73\x0d\x72\x0c\x71\x75\x51\xd0\xd5\x55\x28\x4e\x4d\x55\x28\xc9\xcf\xcf\x29\xd6\xcf\xc9\x2c\x2e\xd1\x4b\xcf\xe7\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\x4d\xcc\xcc\xe3\xe2\x2a\x4b\x2c\x52\x48\xcb\xcc\x49\x2d\x56\xb0\x55\x88\x8e\x2d\x2e\x29\xca\xcc\x4b\xaf\xe6\xe2\x54\x4a\x49\x2d\x49\x4d\x2e\x89\xcf\xc9\xcc\x2b\xad\x00\xea\x52\xd2\x01\x8a\x41\x4d\x00\xb2\x6b\xb9\x00\x01\x00\x00\xff\xff\x84\x1b\x76\x30\x6c\x00\x00\x00")

func listGoBytes() ([]byte, error) {
	return bindataRead(
		_listGo,
		"list.go",
	)
}

func listGo() (*asset, error) {
	bytes, err := listGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "list.go", size: 108, mode: os.FileMode(420), modTime: time.Unix(1442252256, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x57\xdf\x57\xdb\xb8\x12\x7e\x4e\xfe\x0a\x5d\xdf\x93\x62\xb7\xb9\x06\x6e\x9f\x96\xdd\x3c\xb0\x6c\x69\xbb\xdb\x02\x87\xf4\x2d\xcd\xe1\x08\x47\x4e\x74\xb0\x25\x23\x29\x01\x76\xcb\xff\xbe\xdf\x48\xb6\xe3\x34\x81\xd2\x07\xb0\xad\x1f\x33\xdf\xcc\x7c\xf3\x23\x15\xcf\x6e\xf8\x5c\xb0\x92\x4b\xd5\xef\xcb\xb2\xd2\xc6\xb1\xb8\xdf\x8b\xb8\xc9\x16\x72\x25\xf6\xff\x96\x55\x84\xcf\xeb\x65\x2e\x35\xbd\xe4\x05\x9f\xfb\x67\xe9\xe8\x11\x16\xa5\xde\x97\x7a\xe9\x64\x41\x1f\x85\xf6\x07\xb4\x0d\xff\xf7\xc5\xbd\xc8\xe8\xb5\xe2\x6e\xb1\x9f\xcb\x42\xd0\x0b\x2d\x98\xa5\x72\xb2\x14\xf4\x6a\xa1\xd6\x3f\x9d\x91\x6a\x8e\x9b\x49\xbf\xbf\xe2\x86\x90\xac\x84\xb9\xd6\x56\xb0\x11\x23\xd5\xe9\xef\x5a\x17\x71\xb4\x8a\x86\x2c\xe7\x85\x15\x43\x16\x5d\x0b\x96\x2d\xb8\x73\x0f\x11\x2e\xf5\xee\xb4\xb9\x61\x41\x0c\xbe\xa4\x72\xc2\x28\x5e\x9c\x94\x33\x0b\x09\x25\xaf\x26\x61\x6f\x9a\x2f\x55\x16\x4f\xa6\xe1\x2b\xf9\xa7\xdf\xeb\x45\x77\x73\xe1\xa2\x23\x46\x8f\x61\xbf\xf7\x48\x18\x32\xad\xac\xf7\x07\x30\x93\x24\xc8\x88\x26\x07\xff\xfb\x65\xea\xff\x5d\xbd\xf6\x40\xdd\x43\x25\xd8\x3b\xb5\xea\x8a\x6f\x10\x90\x1a\x16\xfb\xed\x84\x7d\x16\x66\x2e\x62\xed\x16\xc2\x6c\x9f\x4d\x18\x40\xe4\xda\x30\x35\x64\x2b\x76\x34\x62\x86\x2b\x04\x26\x9c\x26\x7c\x62\xa2\xa6\xd0\xbf\x22\x68\x8f\xfd\xfe\xfe\x3e\xbb\xd0\xd5\xb2\xe0\x4e\x30\x9c\xf1\x00\xee\xa4\x5b\xf8\x8f\xca\xe8\x4c\x58\xbb\x67\x99\x50\x2b\x69\xb4\x2a\x85\x72\x9b\x60\xce\xc7\x71\xab\xf2\x6a\x53\xa5\x4d\xdf\x85\x5b\xe1\x44\xef\x9e\xf6\xea\xd0\xa4\xe3\xaa\x90\x2e\x5e\xc1\xf1\x23\x38\x9c\x60\xdd\x4f\x0e\xa6\x84\xec\x7e\x72\x38\xad\xc1\x6d\x68\x1a\x17\x32\x13\x10\xd5\x78\x9b\x64\x1a\xe1\x48\x68\xc9\x6f\x44\x1b\x85\x21\x2b\x84\x8a\x45\x02\xa9\x92\x36\x0f\x76\xf9\x43\x78\x44\xb8\x3e\x91\xa4\x13\x34\x04\x22\xdc\x76\x79\x1c\x0d\xec\x68\x60\xc1\x0c\xba\x41\xd0\xe4\x9b\x37\x84\x87\x4e\x2f\x11\x3b\x3c\x5a\x6c\x52\xc1\x08\x6f\x1d\xe8\x9a\x8e\x85\x3b\x05\xb9\x6c\x4c\x1f\x9f\xac\x9b\xf9\x2f\xf6\x8d\x85\xef\x05\xd8\x49\xc4\x4d\x6a\x7a\x8d\xd8\x52\x51\xe2\xc4\x49\x2b\x8f\xf2\xa7\xf6\x27\xb1\xf4\x82\x1b\x0b\x93\xfb\x3d\xb8\x9f\xb0\xc3\x0f\xb1\xb7\x75\x3b\xec\xe1\x50\x4a\xe1\x80\xdd\x39\x7b\xdd\x10\x9e\xec\x24\xf3\x4e\x6b\xf3\x10\x96\xb1\x9b\x09\x63\xe0\xfa\xa0\x5f\xcc\x98\xcd\x8c\xac\x9c\x65\x4e\xb3\xc1\xed\x57\x05\xe3\x09\x61\xe2\xcd\x9e\x89\x1c\xd4\xc1\xb5\x4b\x51\xea\x95\x38\x2e\x8a\x38\x6c\x06\x8d\x81\x8d\x59\x3e\x0f\x60\x9b\xcc\x4c\xff\xd4\x30\x85\x0e\x42\x0f\x76\xa3\xa4\x81\xf8\x82\x0b\x75\x4a\xa7\xef\xcf\xcf\xc7\xbb\xaf\x93\x9d\x49\xcb\xba\x6a\x1d\xd8\x86\x05\xde\xec\x1f\x8a\x5e\x7f\x1d\x5f\x9e\x7c\x48\x86\x2f\xba\x14\x5d\x4b\x15\xd1\xd9\xc7\xe0\x5c\x39\x64\x70\x27\x41\xf0\xbe\xe5\x2e\xae\x3c\x6d\x72\xbf\x3c\x1a\x31\x25\x0b\xf6\xea\x15\xcb\x65\xfa\xd1\xfe\x21\x4d\x9d\x0e\x64\xcd\x24\xba\x38\xfe\xf2\x21\x22\x0a\x56\xec\x4d\x9d\x1c\x6b\x97\x7c\x92\xd6\x8d\xf1\x6a\xb8\xd3\x26\xc1\x81\xce\x15\x48\x78\xa4\xf8\xfc\x44\xb4\x2d\x65\x0d\x44\x1c\x21\xc2\x84\xb0\xf6\x9e\xe8\xa4\x05\x1c\xdc\xe4\x99\xc7\xf8\x94\xac\xaf\x6e\x60\x3d\x4f\x88\xcd\x1d\x20\xb5\x23\xe0\xb0\x19\x0c\xfd\x91\x2f\x93\x61\x57\x63\xf2\xab\xbf\xfe\x9f\xe0\x30\x52\x4f\x79\x73\xca\x1d\x2f\x62\x6c\x24\xeb\x8a\x15\x78\x6b\x43\x8d\xe2\x19\xbf\x99\x77\x38\xcc\xd5\x8c\x85\x54\x0d\x07\x00\x44\x64\xf0\xdf\x03\x7d\x3d\xec\x19\x81\xac\x4d\x43\xba\x35\xf9\xc7\xd6\x15\xc5\x95\x55\x1b\xce\xd0\x8c\xd2\x2f\xa2\xac\x28\x6c\x51\x44\xb1\xd7\xda\xe1\x34\xaf\xa2\xa4\x35\xf8\x59\xc4\x3d\xe0\xd5\x95\x50\x4c\x2f\x0d\x43\x13\x43\x1d\x69\x15\xa0\x2b\xa6\xe7\xd8\xbb\x14\x1c\x7e\x25\x07\x1f\x9b\xb9\x45\x1d\x7c\xa9\xec\x90\x9b\x26\x3d\x29\xb4\x2f\x14\x6d\x46\xe4\xeb\x98\x9a\xf4\x14\x61\xa8\xa9\x4a\xcb\xb9\x5f\xf8\xa8\x72\xed\x6f\x90\xa6\xef\xa9\xb9\x8e\x24\x30\x7d\xbe\xd9\x8e\xa5\x77\x53\x9e\x9e\xf1\x52\x20\x88\xb8\xfe\x59\xcf\x76\x86\x70\x1b\xb6\xa7\x4b\xaf\x87\x9e\xe8\xa4\x5a\x0a\x4f\x1f\xaa\xc5\x79\xeb\x96\xdc\x3b\x25\xee\x24\x51\x57\xe0\x96\x3c\x12\xa7\xbb\x39\x78\x62\x04\xda\xd9\x33\x88\x7f\x46\x32\xce\x5d\x75\x18\x91\x9e\xe8\xea\x21\x86\x3a\x93\xef\x30\xf6\x29\x09\x0d\xb6\xf4\x64\x51\xea\x59\xfc\x9c\xbf\x76\x8a\x30\x79\x1b\x62\xd8\xba\x7e\x27\xcf\xd5\x6d\x09\xf6\xb5\x6d\x64\xa3\xb8\xb2\x66\x30\xd8\x6a\x1a\xa4\x50\xb4\xfd\x73\xbb\xa7\xf4\x7b\x57\x5b\xa5\xad\x6e\x60\xb0\x09\x2b\x1f\xed\x99\x76\xef\xee\x51\xa5\x3c\xd8\xa6\xa7\x12\x1c\xe1\xf9\x99\x77\xef\xfb\xa0\xae\xef\xbf\x94\xdd\x1d\xd3\x2d\x49\xf2\x33\x64\x7a\x26\xee\xc6\x19\x57\x0a\x69\x93\xd7\x8d\x00\x08\xb1\x52\x53\xd8\xcf\x05\x16\xb9\x7b\xef\x1a\x22\xd1\x58\xe0\x12\xf6\x1b\x7b\xcb\xbe\x7d\x63\x0e\x69\x46\xc5\x79\xef\xbf\x7b\xc1\xf1\x1b\x84\xec\xf5\x6c\xb5\x35\xae\x9c\xc5\x0e\x15\x80\xa1\x0c\xbc\x25\x91\x16\x73\x52\xb6\x60\xb6\x22\x49\x24\x22\xe3\xa8\xc0\x51\x56\xce\xa2\xa3\x3a\x89\x36\xea\xf2\xd3\xd5\x14\x35\x51\x51\x3c\x06\x36\x1e\xdc\x26\xbe\xae\x42\xea\xe1\x74\xf8\xdd\xb8\x84\xc5\xff\x63\x31\x1a\x46\xc9\x3a\x91\xba\xc3\xe9\xc4\x5f\x9b\xc6\xcf\x5e\x0b\x30\x51\x78\x03\x4c\x51\x5f\x42\x0f\xf2\x07\x7f\x0e\xba\xc5\x04\x46\x23\x36\x0d\x4d\xb7\x5d\xe4\x8d\xd8\x16\x28\xa2\xc9\x97\x85\x3b\x6a\x48\x7e\x51\x8f\x5b\x4b\x75\xa3\xf4\x9d\x62\xcd\x6c\x3c\xb8\x1d\x32\x39\x57\x9a\x2c\xf0\x02\x5d\xdb\x62\x3a\xa9\x84\xe1\xd2\xa0\x5c\xbd\xa0\x65\xac\x19\xd9\xa4\x47\xdd\x9d\xf0\xc7\x9a\x91\x91\x06\xac\x76\x8c\x27\x99\xda\x4f\xcb\xc4\x56\xbb\xae\x4b\x4d\x41\x79\x5f\xe8\xeb\xef\xca\x0b\x84\x0d\x1b\x13\x92\x9d\x04\x6f\x60\x18\xd3\x98\xe2\xa5\x37\x03\xc2\x66\xe7\x2e\xd4\x86\x97\x95\xae\x0f\xe7\x1a\xd8\x69\x4c\x83\x0d\xa9\x6f\xe2\xb5\x58\x88\xf0\x62\xe9\x27\x10\xee\xf9\xf0\x7b\x88\x96\x6a\x3c\x45\x08\xbc\x64\xaf\xe9\x57\x54\x0a\xa6\xec\x6a\x14\x41\xc3\x13\x23\x4d\xfe\x44\xcd\xec\x5a\xe5\xc3\x8c\x76\x67\x6f\x64\xc5\x70\x56\x69\x47\x1d\x2f\x5b\x3a\x7e\x5d\x88\xb6\xd3\x84\xd2\xf7\xea\xe0\xf0\xf0\x90\x8c\x3f\x68\x5b\xce\x0b\x29\xe7\x8b\xda\xc0\x1e\x7d\x75\xa4\xa8\xf2\xf1\xdb\x54\x75\xc4\x06\xda\x53\x27\xef\xf6\xa6\x27\x5b\x0f\x79\x66\xc4\x82\x6b\x74\x59\x62\x80\x08\xe6\x62\x9d\xb4\x62\x0a\x60\xb5\x23\xe8\x7d\xbd\xe1\xc7\x3b\xd6\x42\x5b\x6f\x48\xd5\xae\xe3\xf7\x70\x58\x46\x77\xc5\x22\x68\x52\x7f\xd3\x0f\xad\x11\xf1\x6e\xa3\x43\xd0\xce\xe5\x52\xed\x20\xf6\xb6\xa7\xb7\x3c\xf6\x63\x87\x81\x35\xec\xfc\xaf\xe0\x9a\xce\xe0\xd6\xe1\xd0\x63\xff\xdf\x00\x00\x00\xff\xff\xe4\x86\x77\x93\xcc\x0f\x00\x00")

func mainGoBytes() ([]byte, error) {
	return bindataRead(
		_mainGo,
		"main.go",
	)
}

func mainGo() (*asset, error) {
	bytes, err := mainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.go", size: 4044, mode: os.FileMode(420), modTime: time.Unix(1442350466, 0)}
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
	"cmds.go":         cmdsGo,
	"detect.go":       detectGo,
	"detect_linux.go": detect_linuxGo,
	"list.go":         listGo,
	"main.go":         mainGo,
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
	"cmds.go":         {cmdsGo, map[string]*bintree{}},
	"detect.go":       {detectGo, map[string]*bintree{}},
	"detect_linux.go": {detect_linuxGo, map[string]*bintree{}},
	"list.go":         {listGo, map[string]*bintree{}},
	"main.go":         {mainGo, map[string]*bintree{}},
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
