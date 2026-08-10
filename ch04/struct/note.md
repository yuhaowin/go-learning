
## 可寻址性（addressability）：为什么函数返回值不能直接赋值字段

```go
func EmployeeByID(id int) *Employee {
	return &Employee{}
}

func EmployeeByID1(id int) Employee {
	return Employee{}
}

EmployeeByID(id).Position = "xxx"  // 编译通过
EmployeeByID1(id).Salary = 0        // 编译错误：cannot assign to struct field
```

`EmployeeByID1(id)` 的返回值是一个 `Employee` 值类型的临时结果（rvalue），没有对应的内存地址，属于**不可寻址**的值。对不可寻址值的字段赋值，赋值了也无处可存、马上丢失，所以 Go 编译器直接拒绝编译。

而 `EmployeeByID(id)` 返回的是 `*Employee`，对指针解引用 `(*p).Position` 得到的是指针指向的内存，是可寻址的，所以可以赋值（Go 会自动帮你解引用，`EmployeeByID(id).Position` 等价于 `(*EmployeeByID(id)).Position`）。

可寻址的值包括：变量、指针解引用、可寻址值的字段/数组元素、slice 的元素（通过底层数组间接可寻址）。
不可寻址的值包括：函数返回的值类型结果、字面量、map 的元素（map[key] 也不可寻址，同理不能对 `m[key].Field = x` 赋值）。

修复方式：
1. 让函数返回指针 `*Employee`；
2. 或者先把返回值存入变量，再对变量的字段赋值：

```go
e := EmployeeByID1(id)
e.Salary = 0
```
