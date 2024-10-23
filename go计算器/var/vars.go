package vars// package 创建包 vars 用于存放全局变量的自定义包

/*  
	注意:存在自定义包里面的封装函数和变量,名称必须大写开头
	否则无法在 可执行包 main 中调用
*/ 

var (// 全局变量
	Number float32// DG 函数中的加数一
	Number0 float32// DG 函数中的加数二
	I int8// Sa 函数里的选择变量
)