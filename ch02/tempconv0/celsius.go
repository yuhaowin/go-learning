package tempconv0

import "fmt"

// 类型声明的方法 （用已有的数据类型，自定义出区别不同业务的类型）
// type 类型名字 底层类型
// 一个类型声明语句创建了一个新的类型名称，和现有类型具有相同的底层结构。
// 主要是用来分隔不同概念的类型，这样即使它们底层类型相同也是不兼容的。

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32) //是类型转换操作，它们并不是函数调用。
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9) //是类型转换操作，它们并不是函数调用。
}

// 命名类型还可以为该类型的值定义新的行为。
// 这些行为表示为一组关联到该类型的函数集合，我们称为类型的方法集。
// 下面的声明语句，Celsius类型的参数c出现在了函数名的前面，表示声明的是Celsius类型的一个名叫String的方法，该方法返回该类型对象c带着°C温度单位的字符串
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

//func String() string {
//	return fmt.Sprintf("%g°C", c)
//}
