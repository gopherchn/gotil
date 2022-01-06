package factory

import (
	"fmt"
)

type ConfigParser interface {
	Parse(data []byte)
}

type JSONConfigParser struct {
}

func (j *JSONConfigParser) Parse(data []byte) {
	fmt.Println("JsonConfigParser: ", string(data))
}

type YamlConfigParser struct {

}

func (y *YamlConfigParser) Parse(data []byte) {
	panic("not implemented") // TODO: Implement
}

// 简单工厂直接创建对象
func NewConfigParser(typ string) ConfigParser {
	switch typ {
	case "yaml":
		return &YamlConfigParser{}
	case "json":
		return &JSONConfigParser{}
	}
	return nil
}