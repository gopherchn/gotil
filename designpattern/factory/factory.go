package factory

// 工厂方法需要实现需要实现工厂初始化创建的方法
// 解决复杂对象的创建
type ConfigParserFactory interface {
	CreateParser() ConfigParser
}

type YamlConfigParserFactory struct {
}

func (y *YamlConfigParserFactory) CreateParser() ConfigParser {
	// 此处可以增加初始化代码
	return &YamlConfigParser{}
}

type JSONConfigParserFactory struct {
}

func (j *JSONConfigParserFactory) CreateParser() ConfigParser {

	// 此处可以随意增减初始化代码
	return &JSONConfigParser{}
}

func NewConfigParserFactory(t string) ConfigParserFactory {
	switch t {
	case "json":
		return &JSONConfigParserFactory{}
	case "yaml":
		return &YamlConfigParserFactory{}
	}
	return nil
}
