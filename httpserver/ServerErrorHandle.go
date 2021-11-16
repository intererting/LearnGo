package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errorHandler(appHandler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := appHandler(writer, request)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func HandlerFileRequest(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/test/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	datas, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	_, err = writer.Write(datas)
	if err != nil {
		return err
	}
	return nil
}
