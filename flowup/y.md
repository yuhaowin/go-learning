Java 是 **静态类型（Static Typing）+ 强类型（Strong Typing）+ 名义类型（Nominal Typing）** 的组合，这三个维度分别回答不同的问题，我逐一说明，并和你之前聊过的
Go/TS 做对比。

## 1. 静态类型 vs 动态类型 —— "类型什么时候检查"

**Java 是静态类型**：类型检查发生在 **编译期**，变量声明时就要确定类型，运行前编译器就能发现类型不匹配的错误。

```java
int x = "hello"; // ❌ 编译期直接报错，程序都跑不起来
```

对比动态类型语言（比如 Python/JavaScript）：类型检查发生在 **运行期**，同一个变量运行时可以指向不同类型的值。

## 2. 强类型 vs 弱类型 —— "类型之间能不能随便互转"

**Java 是强类型**：不同类型之间不会做隐式的"魔法转换"，想转换必须显式处理。

```java
int a = 5;
String b = "3";
int c = a + b; // ❌ 编译错误，int 和 String 不能直接相加
```

对比弱类型语言（比如 JavaScript）：

```javascript
let a = 5;
let b = "3";
let c = a + b; // "53"，隐式把 5 转成字符串拼接了
```

（Java 里字符串拼接用 `+` 时对 `String` 有特殊处理，比如 `"a" + 5` 会变成 `"a5"`，这是语言层面专门为 `String`
设计的重载行为，不是"弱类型隐式转换"，属于个例，不影响 Java 整体强类型的定位。）

## 3. 名义类型 vs 结构化类型 —— "类型兼容性怎么判断"（重点，回应你之前的问题）

**Java 是名义类型（Nominal Typing）**：两个类型是否兼容， **看的是"名字"（类型声明），而不是"长得像不像"**。

```java
interface Flyable {
    void fly();
}

class Bird {
    public void fly() {
        System.out.println("Flying");
    }
}

Flyable f = new Bird(); // ❌ 编译错误！
// 即使 Bird 有一模一样的 fly() 方法，但没写 implements Flyable，
// Java 类型系统就认为它们毫无关系
```

必须显式声明：

```java
class Bird implements Flyable {
    public void fly() {
        System.out.println("Flying");
    }
}

Flyable f = new Bird(); // ✅ 现在可以了
```

这正是你之前问 Go/TS 时我们对比过的地方—— **Go 和 TypeScript 都是结构化类型（Structural Typing）**
，只要方法/字段结构匹配就自动兼容，不需要显式声明"我实现了谁"；而 **Java 必须靠 `implements`/`extends`
这种名义上的血缘关系**，类型系统才认可这层关系。

## 三个维度汇总对比

| 语言           | 静态/动态      | 强/弱类型          | 名义/结构化              |
|----------------|----------------|--------------------|--------------------------|
| **Java**       | 静态           | 强类型             | **名义类型**             |
| **Go**         | 静态           | 强类型             | **结构化类型**           |
| **TypeScript** | 静态（编译期） | 强类型（相对宽松） | **结构化类型**           |
| **Python**     | 动态           | 强类型             | （鸭子类型，运行时判断） |
| **JavaScript** | 动态           | 弱类型             | （鸭子类型，运行时判断） |

## 一个有意思的点

Java 虽然是名义类型，但 **泛型（Generics）在编译期做了类型檫除（Type Erasure）**——也就是说泛型信息只在编译期存在，编译成字节码后
`List<String>` 和 `List<Integer>` 在运行时其实都变成了普通的 `List`。这算是 Java
类型系统里一个常被拿出来讨论的"设计取舍"（主要是为了兼容 Java 5 之前的老代码），和你如果后面深入学 Go 的泛型（Go 1.18+
引入，是真泛型，运行时保留类型信息，没有类型檫除）会形成一个有意思的对比。