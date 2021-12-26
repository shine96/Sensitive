package controllers

import (
	"Sensitive/app/common"
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func Index(c *gin.Context) {
	common.Success(c, 200, "success", nil)
}

func Check(c *gin.Context) {
	filter := common.Sfilter()
	err := filter.LoadWordDict("./dict.txt")
	if err != nil {
		common.Success(c, 200, "无敏感词字典", nil)
	} else {
		context := c.PostForm("context")
		result := filter.FindAll(context)
		if len(result) > 0 {
			common.Success(c, 200, "存在敏感词", result)
		} else {
			common.Success(c, 200, "无敏感词发现", nil)
		}
	}
}

func Add(c *gin.Context) {
	word := c.PostForm("word")
	if word != "" {
		f, err := os.OpenFile("./dict.txt", os.O_WRONLY, 0666)
		if err != nil {
			common.Success(c, 200, "文件读取失败", nil)
		} else {
			file, _ := os.Open("./dict.txt")
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				lineText := scanner.Text()
				if lineText == word {
					common.Success(c, 200, "关键词已存在", nil)
					return
				}
			}
			n, _ := f.Seek(0, os.SEEK_END)
			_, er := f.WriteAt([]byte("\n"+word), n)
			if er != nil {
				common.Success(c, 200, "写入失败", nil)
			} else {
				defer f.Close()
				common.Success(c, 200, "添加成功", nil)
			}
		}
	} else {
		common.Success(c, 200, "参数为空", nil)
	}
}

func List(c *gin.Context) {
	if common.IsFileExist("./dict.txt") {
		f, _ := os.Open("./dict.txt")
		defer f.Close()
		r := bufio.NewScanner(f)
		var line string
		for r.Scan() {
			lineText := r.Text()
			if lineText != "" {
				line += lineText + ","
			}
		}
		line = line[0 : len(line)-1]
		common.Success(c, 200, "ok", line)
	} else {
		common.Success(c, 200, "fail", nil)
	}
}

func Save(c *gin.Context) {
	data := c.PostForm("data")
	if data != "" {
		if common.IsFileExist("./dict.txt") {
			f, _ := os.OpenFile("./dict.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
			defer f.Close()
			list := strings.Split(data, ",")
			for _, text := range list {
				f.WriteString(text + "\n")
			}
			fmt.Println(list)
			common.Success(c, 200, "保存成功", nil)
		} else {
			common.Success(c, 200, "字典文件不存在", nil)
		}
	} else {
		common.Success(c, 200, "参数为空", nil)
	}
}
