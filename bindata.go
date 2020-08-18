// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// gtk/main_window.glade
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _gtkMain_windowGlade = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x4f\x73\xda\x3e\x10\xbd\xe7\x53\xe8\xa7\xeb\x6f\xc0\xfc\x69\xd2\xb4\x63\x9c\x69\xd3\x21\x97\xe4\x92\xa6\xd3\xa3\x47\x96\x16\x5b\x41\x48\xae\xfe\xe0\xd0\x4f\xdf\x31\x4e\x06\x83\x4d\x6c\x48\xa0\x34\xa7\x8c\xa5\xb7\xda\x7d\x7e\x7a\xbb\xd8\xbf\x7a\x9a\x09\x34\x07\x6d\xb8\x92\x23\xdc\xef\xf6\x30\x02\x49\x15\xe3\x32\x1e\xe1\x1f\x0f\xe3\xce\x25\xbe\x0a\xce\xfc\xff\x3a\x1d\x74\x03\x12\x34\xb1\xc0\x50\xc6\x6d\x82\x62\x41\x18\xa0\x61\x77\x78\xd1\xed\xa1\x4e\x27\x38\xf3\xb9\xb4\xa0\x27\x84\x42\x70\x86\x90\xaf\xe1\x97\xe3\x1a\x0c\x12\x3c\x1a\xe1\xd8\x4e\xff\xc7\xab\x83\x86\xdd\xc1\x00\x7b\xcb\x7d\x2a\x7a\x04\x6a\x11\x15\xc4\x98\x11\xbe\xb1\xd3\x2f\xec\xd1\x19\x3b\x03\x69\x31\xe2\x6c\x84\x19\xcc\x39\x85\xd2\xd3\x1c\x86\x90\x9f\x6a\x95\x82\xb6\x0b\x24\xc9\x0c\x46\xd8\xa5\x29\x68\x1c\x0c\xce\xcf\x7d\xef\x65\xa9\x7e\xa7\xb1\x90\x86\x5c\x52\x0d\x45\xb8\x7e\x13\x20\x25\x31\xac\x01\x7a\xeb\x08\xdf\x2b\x8a\x68\x57\x8f\x50\x2a\xbd\x56\x4e\xda\xc6\x92\x84\xca\xf2\x92\x1a\xf3\x7b\x2e\xfd\x53\xfe\xd7\xb4\x77\x4e\x84\x83\x16\x31\x8f\x4c\xd2\x4f\x2e\x99\xca\x0a\x82\x32\x79\x47\xb8\xdc\xc2\x09\x25\x32\x9c\x28\xea\x0c\x0e\xc6\x44\x18\x68\x4a\xcb\x72\x2b\x00\x23\xab\x89\x34\x82\x58\x12\x09\x18\xe1\x05\x18\x1c\x4c\xc8\x14\x28\x99\x75\x62\xc7\x9b\x82\x68\x30\xfc\x77\x0e\x6d\x79\x68\xb6\xac\x26\x4c\x95\xe1\x96\x2b\x89\x03\x0a\xf9\xdd\x68\x82\x31\x98\x10\x27\x6c\x98\x71\x66\x13\x1c\x0c\x2e\x7a\x6d\x11\x09\xf0\x38\xc9\x49\xbf\xac\x42\x68\xc2\x05\x2b\xfe\xaf\x63\xfe\x96\x2c\x94\x7b\x91\xa6\x2d\x31\x5f\xab\x1e\x6e\xf8\x92\x86\x07\xed\x2a\x2c\xec\xf3\xba\xaa\x19\xd6\x67\x39\xe6\x02\xae\x13\xa5\x0c\xe8\xaf\xce\x5a\x25\x8b\x84\x23\x5b\x3c\xcc\x97\x71\x39\x42\xdd\x3b\x61\x36\x09\x73\x57\x02\x63\x71\x30\x18\x56\x98\xda\xbb\xe8\x7d\x0b\xaf\xbd\x77\x89\xca\xc2\x84\x33\x06\x72\xb7\x03\xb7\x0a\xdd\x5b\x63\xb6\x74\x0b\x57\x91\x08\x9d\x72\x19\xbf\x1e\xff\x09\x07\xfd\x8a\xb5\x6e\xdb\xbc\xd8\xbe\xd9\xf7\x2a\xc7\xf9\xde\x86\x02\xda\x28\xe2\x3a\x01\x3a\x2d\x6b\x81\x46\xb7\x4a\xa5\x63\xa5\x61\x9e\xbb\xe1\xab\xf9\x09\x12\x81\xa8\x75\x85\x3c\x06\x9a\x14\x41\x4e\x52\x20\x1a\x28\xf0\x39\x98\xf0\xf9\xee\xef\x08\x67\x9a\x64\x21\x97\x8c\x53\x62\x95\x7e\x2d\xe5\xe3\x49\xa5\xea\x73\xef\x2b\x95\xef\x29\x97\x65\xa5\x98\xa5\x52\x96\x3d\x78\x37\xd3\xe8\xf7\x2f\x8e\xa5\x89\x5d\x60\xa4\x34\x48\xd4\x4c\x17\x6d\xc3\x48\x37\x03\xcd\xe9\xe1\x34\xd1\xde\x72\x17\x38\xa8\xce\x71\xef\x2b\x8a\x07\x15\xc7\x02\xd6\x9b\xc9\xbd\x93\x7b\x1b\xc7\xbd\x93\x48\x49\x54\x4c\xab\x9f\xdb\x16\xba\x29\xb0\xde\x61\xbb\x92\x01\x99\x8f\x24\xf3\x6d\x83\x4c\x0b\x61\xbe\xd1\xac\x4e\xc1\x6f\xb6\xb2\xfc\x5e\xda\xba\x2d\x54\xb2\x1c\xa9\xa2\x3b\xc5\xdc\x01\xa7\x93\x62\xf4\x2b\x41\x4f\x73\xae\xc9\x34\x49\x77\x3b\xa9\xa0\x84\x26\x44\x1b\x1c\x0c\x3e\xfe\x75\xd1\x7c\x38\x7a\x97\xfa\xb6\xf4\x92\x83\xb5\xa8\x4d\xe5\x0c\x07\xa7\xde\xdc\x36\x3f\x05\xfc\x9b\x9d\xed\x0d\xf6\xb3\x9e\xe5\xda\x62\x21\x31\x64\x17\xe9\xcb\x8f\x81\x88\xac\xc6\x60\x3f\x15\x84\x42\xa2\x04\x03\xed\x55\xd0\xab\xb0\xbe\x57\xfa\x8a\xf3\x27\x00\x00\xff\xff\x0c\x70\x5e\xc2\x1e\x12\x00\x00")

func gtkMain_windowGladeBytes() ([]byte, error) {
	return bindataRead(
		_gtkMain_windowGlade,
		"gtk/main_window.glade",
	)
}

func gtkMain_windowGlade() (*asset, error) {
	bytes, err := gtkMain_windowGladeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gtk/main_window.glade", size: 4638, mode: os.FileMode(420), modTime: time.Unix(1597709028, 0)}
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
	"gtk/main_window.glade": gtkMain_windowGlade,
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
	"gtk": &bintree{nil, map[string]*bintree{
		"main_window.glade": &bintree{gtkMain_windowGlade, map[string]*bintree{}},
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
