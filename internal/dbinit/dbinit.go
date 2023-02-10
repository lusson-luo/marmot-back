package dbinit

import (
	"bufio"
	"io"
	"os"
	"regexp"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	dblinkString := g.Cfg().MustGet(gctx.New(), "database.default.link").String()
	reg := regexp.MustCompile(`sqlite::@file\((.*?)\)`)
	dbLinks := reg.FindStringSubmatch(dblinkString)
	if len(dbLinks) != 2 {
		panic("sqlite link 正则匹配失败")
	}

	scriptFile, err := os.OpenFile(g.Cfg().MustGet(gctx.New(), "initSql").String(), os.O_RDONLY, 0666)
	if err != nil {
		panic("未找到 sql 初始化文件")
	}

	reader := bufio.NewReader(scriptFile)
	var sqlStr string = ""
	for {
		str, err := reader.ReadString('\n')
		sqlStr += str
		if err == io.EOF { // io.EOF 表示文件的末尾
			break
		}
	}
	_, err = os.Stat(dbLinks[1])
	if os.IsNotExist(err) {
		file, _ := os.OpenFile(dbLinks[1], os.O_CREATE, 0666)
		defer file.Close()
	}

	db, err := gdb.New(gdb.ConfigNode{
		Link: dblinkString,
	})
	db.GetAll(gctx.New(), sqlStr)

}
