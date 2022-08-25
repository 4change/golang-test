package main

import (
	"bytes"
	"compress/gzip"
	"log"
)

func zip(data []byte) ([]byte, error) {

	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err := zw.Write(data)
	if err != nil {
		log.Print(err)
		return []byte(""), err
	}

	zw.Close()

	return buf.Bytes(), nil
}

func unzip(data []byte) ([]byte, error) {
	bs := bytes.NewBuffer(data)

	r, err := gzip.NewReader(bs)
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}

	var res bytes.Buffer
	_, err = res.ReadFrom(r)
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}

	return res.Bytes(), nil
}
