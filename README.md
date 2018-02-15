# s3

## usage
```go
	cli, err := s3.NewS3("アクセスキー", "アクセスシークレット", "リージョン", "バケット名")
	if err != nil {
		panic(err)
	}
	// アップロードして公開する
	cli.UpLoadPublic("./s3.go", "hoge/s3.go")
```