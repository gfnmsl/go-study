package main// package 创建包 main 特殊包名,可执行的包(也就是程序入口)

/*  
	注意:存在自定义包里面的封装函数和变量名,必须大写开头
	否则无法在 可执行包 main 中调用
*/ 

import(// 导入需要的包

	"github/calculator/func"// 自定义包,用于存放封装函数
		
)

func main() {
	funcs.Openmain()// 调用自定义包 func 中的 funcs 文件里的 Openmain 函数模块
	funcs.Gd()// 调用自定义包 func 中的 funcs 文件里的 Gd 函数模块
	funcs.Sa()// 调用自定义包 func 中的 funcs 文件里的 Sa 函数模块
}