package Algorithm

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are
stored in reverse order and each of their nodes contain a single digit. Add the two numbers and
return it as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

type ListNode struct {
	val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	head := &ListNode{val: 0, Next: nil}
	current := head
	carry := 0
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.val
		}
		if l2 == nil {
			y = 0
		} else {
			y = l2.val
		}
		//存的node的val取个位
		current.Next = &ListNode{val: (x + y + carry) % 10, Next: nil}
		current = current.Next
		//carry取进位
		carry = (x + y + carry) / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 { //进行收尾
		current.Next = &ListNode{val: carry % 10, Next: nil}
	}
	return head.Next
}

/*
	总结：
		这里学到Go中的结构体写法
			I.define: type structName { propName:propType}
			II.access: node.val ... ,结构体在定义时结构体名开头为小写时这个结构体在其他包是访问不到的。
			III.pointer:GO是值传递，想要在其它地方进行访问和修改要传递指针，类型为 *structName
			IV. GO是不支持继承的，只支持结构体的组合。
			V.结构体中不能包含自己，比如上面ListNode的next字段是*ListNode类型的而不是ListNode类型的。
	拓展：
		1. 方法
			I.普通函数: func funcName(arg1 argType) returnType {}
				这为一个普通的函数
			II.方法： 当一个函数被绑定到某个对象上（将对应类型写到func关键字前类似于 func (m Member)MemberFunc(arg ArgType）)，我们称这个函数为这个对象的方法。
				而且，并不是只有我们自己定义的结构体才能有方法，我们也可以将方法绑定到系统内建的类型上。
				这个m称为方法的接收器。类比其他语言中的this
			III.需要注意的:
				a.go为值传递类似于：
					func (m Member)MemberFunc(arg ArgType）)
					m.setArg('a different val')
					fmt.Println(m.Arg) 我们会发现这个属性的值并没有被修改
					我们需要传递Member类型的指针:
						func (m *Member)MemberFunc(arg ArgType）)
				b.方法也和字段一样，如果首字母是小写只能在包内能访问得到。
		2. tags 在定义结构体字段时，除了字段名称和数据类型外，还可以使用反引号为结构体字段声明元信息，这种元信息称为tag,用于编译阶段关联到字段当中，

			type Member struct {
				Id     int    `json:"id,-"`
				Name   string `json:"name"`
				Email  string `json:"email"`
				Gender int    `json:"gender,"`
				Age    int    `json:"age"`
			}
			Tag由反引号括起来的一些列用空格分隔的key:"value"键值对组成:Id int `json:"id" gorm:"AUTO_INCREMENT"`
		3. 组合
			何为组合，可以理解为一个结构体中又包含其他的结构体。

			例如：
				type Animal struct {
					Name   string  //名称
					Color  string  //颜色
					Height float32 //身高
					Weight float32 //体重
					Age    int     //年龄
				}
				//奔跑
				func (a Animal)Run() {
					fmt.Println(a.Name + "is running")
				}
				//吃东西
				func (a Animal)Eat() {
					fmt.Println(a.Name + "is eating")
				}

				type Cat struct {
					a Animal
				}
				Cat可以共用Animal中的一些字段
				而且也支持：
					type Lion struct {
						Animal
					}
					在初始化时：
					var lion = Lion{
						Animal {
						Name:'BigLion',
						Color:'Grey'
						}
					}
				不需要连续多个·去访问对于的字段。

*/
