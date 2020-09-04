package sql2struct

import (
	"html/template"
	"os"
	"strings"
)

const templateText = `
	Output 0: {{title .Name1}}
	Output 1: {{title .Name2}}
	Output 2: {{.Name3 | title}}
`

func main(){
	// strings.Title()返回s中每个单词的首字母都改为标题格式的字符串拷贝
	// template.FuncMap是自定义函数的作用，前面是函数名，后面是调用的函数名，定义完成之后，后面需要用Funcs解析
	funcMap := template.FuncMap{"title": strings.Title}
	tpl, _ := template.New("go-programmming-tour").Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "go",
		"Name2": "programming",
		"Name3": "tour",
	}

	/**
	在写template的时候，会经常用到"."。比如{{.}}、{{len .}}、{{.Name}}、{{$x.Name}}等等
	在template中，点"."代表当前作用域的当前对象。它类似于java/c++的this关键字，类似于perl/python的self。如果了解perl，它更可以简单地理解为默认变量$_。

	例如，前面示例test.html中{{.}}，这个点是顶级作用域范围内的，它代表Execute(w,"hello worold")的第二个参数"hello world"。也就是说它代表这个字符串对象。

	再例如，有一个Person struct。
	这里{{.Name}}和{{.Age}}中的点"."代表的是顶级作用域的对象p，所以Execute()方法执行的时候，会将{{.Name}}替换成p.Name，同理{{.Age}}替换成{{p.Age}}。

	但是并非只有一个顶级作用域，range、with、if等内置action都有自己的本地作用域。它们的用法后文解释，这里仅引入它们的作用域来解释"."。

	例如下面的例子，如果看不懂也没关系，只要从中理解"."即可。
	 */
	_ = tpl.Execute(os.Stdout, data)
}