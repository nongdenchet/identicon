package main

import (
	"github.com/nongdenchet/identicon/encode"
	"github.com/nongdenchet/identicon/process"
)

func main() {
	process.GenerateImage(encode.GetMD5Hash("nongdenchet"), 500)
	process.GenerateImage(encode.GetMD5Hash("hello world"), 500)
	process.GenerateImage(encode.GetMD5Hash("grab"), 500)
}
