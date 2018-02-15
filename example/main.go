package main

import (
	"github.com/makki0205/s3"
	uuid "github.com/satori/go.uuid"
)

func main() {
	mys3, err := s3.NewS3("アクセスキー", "アクセスシークレット", "リージョン", "バケット名")
	if err != nil {
		panic(err)
	}
	mys3.Up("./s3.go", "hoge/s3.go")
	uid := uuid.NewV4().String()
}
