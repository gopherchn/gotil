package factory

// 抽象工厂
// 可以创建不同的类型的工厂
// 全部都集成在一起
type SystemConfigParser interface {
	SystemParse(data []byte)
}

type JSONSystemConfigParser struct {
}

func (j *JSONSystemConfigParser) SystemParse(data []byte) {

}

// d
type AbstractConfigParser interface {
	CreateSystemParser() SystemConfigParser
	CreateConfigParser() ConfigParser
}

// 实现AbstractConfigParser接口
type AbstractJSONConfigParserFactory struct {
}

func (a *AbstractJSONConfigParserFactory) CreateSystemParser() SystemConfigParser {
	// 此处也可以添加一些初始化的代码
	return &JSONSystemConfigParser{}
}

func (a *AbstractJSONConfigParserFactory) CreateConfigParser() ConfigParser {
	// 此处也可以添加一些初始化的代码
	return &JSONConfigParser{}
}
