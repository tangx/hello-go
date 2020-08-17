package bufferhandler

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/fatih/color"
)

type Preset = map[string][]byte

var Presets = map[string]Preset{}

func Register(name string, preset Preset) {
	Presets[name] = preset
}

func init() {
	Register("gt", Preset{".git-template.yaml": []byte(`
	# <类型>: (类型的值见下面描述) <主题> (最多50个字)
	# feat: 增加了一个什么样的新功能
	
	# 解释为什么要做这些改动
	# |<----  请限制每行最多72个字   ---->|
	
	# 提供相关文章和其它资源的链接和关键字
	# 例如: Github issue #23
	
	# --- 提交 结束 ---
	# 类型值包含
	#    feat = 'Features', // 功能点
	#    fix = 'Bug Fixes', // bug 修复
	#    docs = 'Documentation', // 文档改动
	#    style = 'Styles',  // 格式化, 缺失分号等; 不包括生产代码变动
	#    refactor = 'Code Refactoring',  // 重构代码
	#    perf = 'Performance Improvements', // 性能提升
	#    test = 'Tests',  // 添加缺失的测试, 重构测试, 不包括生产代码变动
	#    build = 'Builds', // 编译
	#    ci = 'Continuous Integration',  // CI
	#    chore = 'Chores',  // 杂事, 更新grunt任务等; 不包括生产代码变动
	#    revert = 'Reverts',  // 回滚
`)})
}

// func Write(name string) []byte {
// 	b := Presets[name]

// 	for _, data := range b {
// 		// err := ioutil.WriteFile(f, data, os.ModePerm)
// 		// if err != nil {
// 		// 	panic(err)
// 		// }

// 		return data
// 	}

// 	return nil
// }

func WriteTo(w io.Writer, name string) {

	b := Presets[name]
	for _, data := range b {
		n, err := w.Write(data)
		if err != nil {
			color.Red(fmt.Sprint(err))
		}
		color.Green(strconv.Itoa(n))
	}
}

func bufferWriteTo() {
	buf := bytes.Buffer{}
	n, err := buf.Write([]byte("你好"))
	if err != nil {
		color.Red(fmt.Sprint(err))
	}
	color.Green(strconv.Itoa(n))

	_, _ = buf.WriteTo(os.Stdout)
}
