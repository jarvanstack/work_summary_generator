package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/pflag"
)

// 命令行参数
var (
	ttheme = pflag.StringP("theme", "t", "后端开发", "参数主题")
	ttimes = pflag.Uint32P("times", "c", 10, "生成次数")
	wwrods = pflag.Uint32P("words", "w", 50, "生成字数")
)

/*
${生成总结份数}
${主题}
${生成字数}

for 生成总结份数
article := ""
for {
    str := randomStr()
    article +=  str
    if 字数足够 {
        退出
    }
}
article = replace(article, "${THEME}", ${主题})
输出
*/

// readData 从文件读取配置
func readData() (data []string) {
	f, err := os.Open("./data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			return
		}
		data = append(data, string(line))
	}
}

func main() {
	pflag.Parse()
	theme := *ttheme
	times := *ttimes
	wordCount := *wwrods
	generate(theme, int(times), int(wordCount))
}

func generate(theme string, times int, wordCount int) {
	datas := readData()
	buf := bytes.Buffer{}
	for i := 0; i < times; i++ {
		count := 0
		for {
			str := randomOne(datas)
			buf.WriteString(str)
			count += utf8.RuneCountInString(str)
			if count >= wordCount {
				buf.WriteString("。")
				break
			} else {
				buf.WriteString("，")
			}
		}
		str := buf.String()
		str = strings.ReplaceAll(str, "${THEME}", theme)
		out(str)
		buf.Reset()
	}
}

func randomOne(datas []string) string {
	return datas[rand.Intn(len(datas))]
}

func out(s string) {
	fmt.Println()
	fmt.Println(s)
	fmt.Println()
}
