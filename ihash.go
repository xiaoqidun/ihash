package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"golang.org/x/crypto/sha3"
	"hash"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

var (
	args    []string
	author  bool
	install string
)

var typeList = []string{
	"md5sum",
	"sha1sum",
	"sha256sum",
	"sha512sum",
	"sha3sum224",
	"sha3sum256",
	"sha3sum384",
	"sha3sum512",
}

func init() {
	flag.BoolVar(&author, "author", false, "")
	flag.StringVar(&install, "install", "", "")
	flag.Parse()
	args = flag.Args()
}

func main() {
	if author {
		AuthorInformation()
		return
	}
	if "" != install {
		InstallShortcut(install)
		return
	}
	fileName := GetFileName()
	hashIndex := InArray(fileName, typeList)
	if -1 == hashIndex {
		if 0 == len(args) {
			return
		}
		hashType := args[0]
		hashIndex = InArray(hashType, typeList)
		if -1 == hashIndex {
			return
		}
		args = args[1:]
	}
	if 0 == len(args) {
		stdinBytes, _ := ioutil.ReadAll(os.Stdin)
		fmt.Println(StrHash(typeList[hashIndex], bytes.TrimSpace(stdinBytes)))
		return
	}
	hashFileList := GetFileList(args)
	for _, hashFile := range hashFileList {
		FileHash(typeList[hashIndex], hashFile)
	}
}

func InArray(value interface{}, array interface{}) int {
	switch array.(type) {
	case []string:
		for k, v := range array.([]string) {
			if v == value.(string) {
				return k
			}
		}
	}
	return -1
}

func StrHash(hashType string, hashData []byte) string {
	var sum interface{}
	switch hashType {
	case "md5sum":
		sum = md5.Sum(hashData)
	case "sha1sum":
		sum = sha1.Sum(hashData)
	case "sha256sum":
		sum = sha256.Sum256(hashData)
	case "sha512sum":
		sum = sha512.Sum512(hashData)
	case "sha3sum224":
		sum = sha3.Sum224(hashData)
	case "sha3sum256":
		sum = sha3.Sum256(hashData)
	case "sha3sum384":
		sum = sha3.Sum384(hashData)
	case "sha3sum512":
		sum = sha3.Sum512(hashData)
	}
	return fmt.Sprintf("%x", sum)
}

func FileHash(hashType string, hashFile string) {
	msg := "progress:%.2f%%    filename:%s\r"
	success := "%x    %s\n"
	fileOpen, err := os.Open(hashFile)
	if err != nil {
		return
	}
	defer func() {
		_ = fileOpen.Close()
	}()
	fileStat, err := fileOpen.Stat()
	if err != nil || fileStat.IsDir() {
		return
	}
	fileSize := fileStat.Size()
	readSize := 0
	readBytes := make([]byte, 8192)
	var hashHandle interface{}
	switch hashType {
	case "md5sum":
		hashHandle = md5.New()
	case "sha1sum":
		hashHandle = sha1.New()
	case "sha256sum":
		hashHandle = sha256.New()
	case "sha512sum":
		hashHandle = sha512.New()
	case "sha3sum224":
		hashHandle = sha3.New224()
	case "sha3sum256":
		hashHandle = sha3.New256()
	case "sha3sum384":
		hashHandle = sha3.New384()
	case "sha3sum512":
		hashHandle = sha3.New512()
	}
	printLength := 0
	for {
		n, err := fileOpen.Read(readBytes)
		if err != nil {
			break
		}
		readSize += n
		hashHandle.(hash.Hash).Write(readBytes[:n])
		if 0 == readSize%67108864 {
			printMsg := fmt.Sprintf(msg, float64(readSize)/float64(fileSize)*100, fileStat.Name())
			printMsgLength := len(printMsg)
			if printLength > printMsgLength {
				printMsg += strings.Repeat(" ", printLength-printMsgLength)
			}
			fmt.Print(printMsg)
			_ = os.Stdout.Sync()
			printLength = len(printMsg)
		}
	}
	printMsg := fmt.Sprintf(success, hashHandle.(hash.Hash).Sum(nil), hashFile)
	printMsgLength := len(printMsg)
	if printLength > printMsgLength {
		printMsg += strings.Repeat(" ", printLength-printMsgLength)
	}
	fmt.Print(printMsg)
	_ = os.Stdout.Sync()
}

func GetFileName() string {
	filePath, _ := os.Executable()
	fileName := filepath.Base(filePath)
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return fileName
}

func GetFileList(args []string) (fileList []string) {
	var tempList []string
	for _, v := range args {
		globList, _ := filepath.Glob(v)
		tempList = append(tempList, globList...)
	}
	sort.Strings(tempList)
	for i := 0; i < len(tempList); i++ {
		if i > 0 && tempList[i] == tempList[i-1] {
			continue
		}
		fileList = append(fileList, tempList[i])
	}
	return
}

func InstallShortcut(dst string) {
	execPath, err := os.Executable()
	if err != nil {
		return
	}
	filePath, err := filepath.Abs(dst)
	if err != nil {
		return
	}
	fileStat, err := os.Stat(filePath)
	if err != nil || !fileStat.IsDir() {
		return
	}
	osPathSeparator := string(os.PathSeparator)
	if !strings.HasSuffix(filePath, osPathSeparator) {
		filePath += osPathSeparator
	}
	for i := 0; i < len(typeList); i++ {
		symlinkPath := filePath + typeList[i]
		if "windows" == runtime.GOOS {
			symlinkPath += ".exe"
		}
		err := os.Link(execPath, symlinkPath)
		if err != nil {
			fmt.Printf("为 %s\t<<===>>\t%s 创建链接错误\n", symlinkPath, execPath)
		} else {
			fmt.Printf("为 %s\t<<===>>\t%s 创建了硬链接\n", symlinkPath, execPath)
		}
	}
}

func AuthorInformation() {
	fmt.Println("welcome to our website https://aite.xyz/")
	fmt.Println("----------------------------------------")
	fmt.Println("腾讯扣扣：88966001")
	fmt.Println("电子邮箱：xiaoqidun@gmail.com")
	fmt.Println("----------------------------------------")
}
