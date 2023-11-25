package core

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	util "github.com/hktalent/go-utils"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

//var KvCc = util.NewKvCachedb()

func RandomStr(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// LinesInFile 读取文件 返回每行的数组
func LinesInFile(fileName string) ([]string, error) {
	var result []string
	var f *os.File
	var err error
	var rd io.Reader
	if !util.FileExists(fileName) {
		s1 := strings.Replace(fileName, ".txt", ".zip", -1)
		if util.FileExists(s1) {
			if f, err = os.Open(s1); nil == err {
				defer f.Close()
				rd, err = gzip.NewReader(f)
			}
		} else {
			err = errors.New("file not exists")
		}
	} else {
		f, err = os.Open(fileName)
		if err == nil {
			rd = f
			defer f.Close()
		}
	}
	if nil != err {
		return result, err
	}
	scanner := bufio.NewScanner(rd)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			result = append(result, line)
		}
	}
	if err = scanner.Err(); nil != err {
		fmt.Println(err)
	}
	return result, nil
}

// LinesReaderInFile 读取文件，返回行数
func LinesReaderInFile(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var readSize int
	var count int
	buf := make([]byte, 1024)

	for {
		readSize, err = r.Read(buf)
		if err != nil {
			break
		}
		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], '\n')
			if i == -1 || readSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
	}
	if readSize > 0 && count == 0 || count > 0 {
		count++
	}
	if err == io.EOF {
		return count, nil
	}
	return count, err
}

func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GetWindowWith() int {
	w, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0
	}
	return w
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func SliceToString(items []string) string {
	ret := strings.Builder{}
	ret.WriteString("[")
	ret.WriteString(strings.Join(items, ","))
	ret.WriteString("]")
	return ret.String()
}
func HasStdin() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return false
	}
	return true
}
