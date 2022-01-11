package prototype

import (
	"encoding/json"
	"fmt"
	"time"
)

// 原型对象：对象的创建成本比较高
// 同一个类型的创建的对象差别不大
// 深拷贝：递归复制，序列化与反序列化
// 浅拷贝: 仅复制对象的引用
// 产生样品

type Keyword struct {
	Word     string
	Visit    int
	UpdatedAt *time.Time
}

func (k *Keyword) Clone() *Keyword {
	newKeyworkd := Keyword{}

	b, err := json.Marshal(k)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	err = json.Unmarshal(b, &newKeyworkd)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &newKeyworkd
}


// 初始化之后便
type Keywords map[string]*Keyword
func (w Keywords) Clone(words [] *Keyword) Keywords {

	// 此处浅克隆的意义
	newKeyWords := Keywords{}
	for k, v := range w {
		newKeyWords[k] = v
	}
	// 深拷贝
	// words在当前程序调用完成之后会销毁，所以需要进行深度克隆
	for _, word := range words {
		newKeyWords[word.Word] = word.Clone()
	}

	return newKeyWords
}
