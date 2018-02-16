package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func copyFiles(inFilePath, outFilePath string) {

	inFile, err := os.Stat(inFilePath)

	if nil != err {
		fmt.Println("exception: file failed...")
		panic(err)
	}

	if inFile.IsDir() {

		err = os.MkdirAll(outFilePath, 0777)

		if err != nil {
			fmt.Printf("%s", err)
		} else {
			fmt.Println("Create Directory OK!")
		}

		inPath, error := ioutil.ReadDir(inFilePath)

		if nil != error {
			fmt.Println("exception: fileDir failed...")
			panic(error)
		}

		for _, v := range inPath {
			newInFilePath := inFilePath + string(os.PathSeparator) + v.Name()
			newOutFilePath := outFilePath + string(os.PathSeparator) + v.Name()
			fmt.Println(inFilePath)
			copyFiles(newInFilePath, newOutFilePath)
		}
	} else {
		outFilePath = outFilePath + string(os.PathSeparator) + inFile.Name()
		CopyFile(inFilePath, outFilePath)
	}
}

// CopyFile 文件的复制方法
func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		// panic(err)
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		// panic(err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
