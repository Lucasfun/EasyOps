package _interface

//一个unit具有名称、被依赖unit数组、安装func
type UnitInterface interface {
	InstallFunc(callback func()) bool
	GetName() string
	GetNext() []string
}
