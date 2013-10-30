package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("upload.gtpl")
		currTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20) // 設置 max memeory
		// 驗證 token ... ...

		fmt.Println("do copy user's upload file to web side :D")
		userUploadFile, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println("get user upload file err:", err)
			return
		}
		defer userUploadFile.Close()

		copyDestFile, err := os.OpenFile("./uploadfiles/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("open copy dest file err: ", err)
			return
		}
		defer copyDestFile.Close()

		if _, err := io.Copy(copyDestFile, userUploadFile); err != nil {
			fmt.Fprintf(w, "upload your file (%s) fail, err=%v", handler.Filename, err)
		} else {
			fmt.Fprintf(w, "upload your file (%s) success", handler.Filename)
		}
	}
}
