# 页面模块化
---

页面自定义需要调用引擎的```Content```方法，需要返回一个对象```types.Panel```

以下是```types.Panel```的定义：

```go
type Panel struct {
	Content     template.HTML   // 页面内容
	Title       string          // 页面标题
	Description string          // 页面描述
	Url         string
}
```

对应的ui，可以看下图：

![](http://quizfile.dadadaa.cn/everyday/app/jlds/img/006tNbRwly1fxoz5bm02oj31ek0u0wtz.jpg)

## 如何使用

```go
package datamodel

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
)

func GetContent() (types.Panel, error) {

	components := template2.Get(config.Get().THEME)
	colComp := components.Col()

	infobox := components.InfoBox().
		SetText("CPU TRAFFIC").
		SetColor("blue").
		SetNumber("41,410").
		SetIcon("ion-ios-gear-outline").
		GetContent()

	var size = map[string]string{"md": "3", "sm": "6", "xs": "12"}
	infoboxCol1 := colComp.SetSize(size).SetContent(infobox).GetContent()
	row1 := components.Row().SetContent(infoboxCol1).GetContent()

	return types.Panel{
		Content:     row1,
		Title:       "Dashboard",
		Description: "this is a example",
	}, nil
}
```

## Col 列

一个col列对应的是```ColAttribute```这个类型，有三个方法如下：

```go
type ColAttribute interface {
	SetSize(value map[string]string) ColAttribute   // 设置大小
	SetContent(value template.HTML) ColAttribute    // 设置本列的内容
	GetContent() template.HTML                      // 获取内容
}
```

关于```size```，参考的例子是：```map[string]string{"md": "3", "sm": "6", "xs": "12"}```

## Row 行

一个row行对应的是```RowAttribute```这个类型，有两个方法如下：

```go
type RowAttribute interface {
	SetContent(value template.HTML) RowAttribute  // 设置内容
	GetContent() template.HTML                    // 获取内容
}
```
