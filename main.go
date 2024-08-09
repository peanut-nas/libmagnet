package main

import (
	"C"
	"bt/metainfo"
	"os"
)

func main() {
	var magnet, err = readTorrentFile("test.torrent")
	if err != nil {
		println(err.Error())
	}
	println(magnet)
}

//export TorrentToMagnet
func TorrentToMagnet(str *C.char) *C.char {
	var path = C.GoString(str)
	var magnet, err = readTorrentFile(path)
	if err != nil {
		return C.CString(err.Error())
	}
	return C.CString(magnet)
}

func readTorrentFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	mi, err := metainfo.Load(f)
	if err != nil {
		return "", err
	}
	ts := TorrentSpecFromMetaInfo(mi)
	m := metainfo.Magnet{
		InfoHash:    ts.InfoHash,
		DisplayName: ts.DisplayName,
	}
	return m.String(), nil
}
