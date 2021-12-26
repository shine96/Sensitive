package common

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func IsFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println(info)
		return false
	} else {
		return true
	}
}

func DownloadFile(url string, localPath string) error {
	var (
		buf     = make([]byte, 32*1024)
		written int64
	)
	tmpFilePath := localPath
	//创建一个http client
	client := new(http.Client)
	//client.Timeout = time.Second * 60 //设置超时时间
	//get方法获取资源
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	//读取服务器返回的文件大小
	if IsFileExist(localPath) {
		return err
	}
	//创建文件
	file, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if resp.Body == nil {
		return errors.New("body is null")
	}
	defer resp.Body.Close()
	//下面是 io.copyBuffer() 的简化版本
	for {
		//读取bytes
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			//写入bytes
			nw, ew := file.Write(buf[0:nr])
			//数据长度大于0
			if nw > 0 {
				written += int64(nw)
			}
			//写入出错
			if ew != nil {
				err = ew
				break
			}
			//读取是数据长度不等于写入的数据长度
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	fmt.Println(err)
	if err == nil {
		file.Close()
		err = os.Rename(tmpFilePath, localPath)
	}
	return err
}

func ReadLine(r *bufio.Reader) (string, error) {
	line, isprefix, err := r.ReadLine()
	for isprefix && err == nil {
		var bs []byte
		bs, isprefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return string(line), err
}
