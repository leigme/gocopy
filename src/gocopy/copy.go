package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func copyFiles(inFilePath, outFilePath string) {

	inPath, err := ioutil.ReadDir(inFilePath)
	if nil != err {
		fmt.Println("exception: file failed...")
		panic(err)
	}

	err = os.MkdirAll(outFilePath, 0777)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Print("Create Directory OK!")
	}

	for _, v := range inPath {
		inFilePath = inFilePath + string(os.PathSeparator) + v.Name()
		outFilePath = outFilePath + string(os.PathSeparator) + v.Name()
		fmt.Println(inFilePath)
		fmt.Println(outFilePath)
		// if v.IsDir() {
		// 	copyFiles(inFilePath, outFilePath)
		// } else {
		// 	CopyFile(inFilePath, outFilePath)
		// }
	}
}

// CopyFile 文件的复制方法
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
