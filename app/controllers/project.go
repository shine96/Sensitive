package controllers

import (
	"Sensitive/app/common"
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func CreateProject(c *gin.Context) {
	project := c.PostForm("project")
	if project == "" {
		common.Success(c, 200, "参数为空", nil)
	} else {
		if common.IsFileExist("./"+project+"/dict.txt") == false {
			os.Mkdir("./"+project, 0777)
			common.DownloadFile("https://media.bnickolas.com/dict_1640512181457.txt", "./"+project+"/dict.txt")
			common.Success(c, 200, "项目创建成功", nil)
		} else {
			common.Success(c, 200, "项目已存在", nil)
		}
	}
}

func ProjectCheck(c *gin.Context) {
	project := c.PostForm("project")
	context := c.PostForm("context")
	if project != "" && context != "" {
		filter := common.Sfilter()
		err := filter.LoadWordDict("./" + project + "/dict.txt")
		if err != nil {
			common.Success(c, 200, "无敏感词字典", nil)
		} else {
			result := filter.FindAll(context)
			if len(result) > 0 {
				common.Success(c, 200, "存在敏感词", result)
			} else {
				common.Success(c, 200, "无敏感词发现", nil)
			}
		}
	} else {
		common.Success(c, 200, "参数为空", nil)
	}
}

func DelProject(c *gin.Context) {
	project := c.PostForm("project")
	if project != "" {
		_, err := os.Stat("./" + project)
		if err != nil {
			if os.IsExist(err) {
				os.Remove("./" + project)
				common.Success(c, 200, "项目移除成功", nil)
				return
			}
			if os.IsNotExist(err) {
				common.Success(c, 200, "项目不存在", nil)
				return
			}
		} else {
			common.Success(c, 200, "项目不存在", nil)
		}
	} else {
		common.Success(c, 200, "参数为空", nil)
	}
}

func ProjectAddWord(c *gin.Context) {
	project := c.PostForm("project")
	word := c.PostForm("word")
	if word != "" && project != "" {
		f, err := os.OpenFile("./"+project+"/dict.txt", os.O_WRONLY, 0666)
		if err != nil {
			common.Success(c, 200, "文件读取失败", nil)
		} else {
			file, _ := os.Open("./" + project + "dict.txt")
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

func SaveAll(c *gin.Context) {
	data := c.PostForm("data")
	project := c.PostForm("project")
	if data != "" && project != "" {
		if common.IsFileExist("./" + project + "/dict.txt") {
			f, _ := os.OpenFile("./"+project+"/dict.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
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

func DictList(c *gin.Context) {
	project := c.Param("project")
	if project != "" {
		if common.IsFileExist("./" + project + "/dict.txt") {
			f, _ := os.Open("./" + project + "/dict.txt")
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
	} else {
		common.Success(c, 200, "参数为空", nil)
	}
}
