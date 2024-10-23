package funcs// package 创建包 funcs 用于存放封装函数的自定义包

/*  
	注意:存在自定义包里面的封装函数和变量名,必须大写开头
	否则无法在 可执行包 main 中调用
*/  

import(// 导入包
	"fmt"// go的基础包,用于输入输出
	"github/calculator/var"// 自定义包,用于存放全局变量
)

func Openmain()  {// 开始介绍函数模块
	fmt.Println("\n欢迎使用由开源作者 github/gfnmsl 制作的无 gui 计算器")
	fmt.Println("\n如果对该项目感兴趣,可以在 github 仓库下载源代码\ngithub仓库:  \n")
	fmt.Println("如果你无法访问github,可以b站私信,我发你\nb站主页: https://space.bilibili.com/589956651 \n")
	fmt.Println("后期会添加 gitee 下载渠道\n")
}

func Gd(){// 收集数值函数模块
	fmt.Println("输入需要计算的两个数")
	fmt.Scan(&vars.Number)// 将输入的值存放在 vars 文件里的全局变量 Number 和 Number0 中( vars 位于自定义包 var 中)
	fmt.Scan(&vars.Number0)
}

func Sa(){// 选择算法模块
	var r float32// 用于存放计算结果的局部变量
	
	for{// 创建死循环
	fmt.Println("选择算法(1 加法 2 减法 3 乘法 4 除法)")
	fmt.Scan(&vars.I)// 输入数子选择算法

	if vars.I == 1{// i == 1 选择加法 
		
	r = vars.Number + vars.Number0// 将变量 vars.Number 和 vars.Number0 的计算结果存进局部变量 r 中
	fmt.Println("加数一:",vars.Number,"加数二:",vars.Number0,"\n计算结果是:",r,"(算法:加法)")
	break// 退出循环

	}else if vars.I == 2{// i == 2 选择减法 

	r = vars.Number - vars.Number0// 将变量 vars.Number 和 vars.Number0 的计算结果存进局部变量 r 中
	fmt.Println("减数一:",vars.Number,"减数二:",vars.Number0,"\n计算结果是:",r,"(算法:减法)")
	break// 退出循环

	}else if vars.I == 3{// i == 3 选择乘法 

	r = vars.Number * vars.Number0// 将变量 vars.Number 和 vars.Number0 的计算结果存进局部变量 r 中
	fmt.Println("乘数一:",vars.Number,"乘数二:",vars.Number0,"\n计算结果是:",r,"(算法:乘法)")
	break// 退出循环

	}else if vars.I == 4{// i == 4 选择除法 

	r = vars.Number / vars.Number0// 将变量 vars.Number 和 vars.Number0 的计算结果存进局部变量 r 中
	fmt.Println("除数一:",vars.Number,"除数二:",vars.Number0,"\n计算结果是:",r,"(算法:除法)")
	break// 退出循环

	}else{// 输入错误
		
	fmt.Println("错误,检查你的输入是否符合规范")

	}
}

}