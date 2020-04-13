package main



const   (
	RED = true
	BLACK = false
)

//红黑树结构
type RBnode struct {
	Left *RBnode  //左节点
	Right *RBnode //右节点
	Parent *RBnode //父节点
	Color bool   //颜色
	//DataItem interface{}  //数据
	Item  //数据接口
}
//存放数据接口
type Item interface {
	Less(than Item)bool
}

//定义红黑树结构
type RBtree struct {

	NIL *RBnode
	Root *RBnode
	count uint
}

//比较大小
func less(x , y Item) bool  {
	return x.Less(y)
}

func NewRBtree() *RBtree  {
	return new(RBtree).Init()    //创建一个内存初始化
}

func (rbt *RBtree)Init() *RBtree  {
	node := &RBnode{
		Left:   nil,
		Right:  nil,
		Parent: nil,
		Color:  BLACK,
		Item:   nil,
	}
	return &RBtree{
		NIL:   node,
		Root:  node,
		count: 0,
	}
}


//获取长度
func (rbt *RBtree)Len() uint  {
	return rbt.count
}

//取得红黑树的极大值
func (rbt *RBtree)max (x *RBnode) *RBnode  {
	if x== rbt.NIL{
		return rbt.NIL
	}
	for x.Right != rbt.NIL{
		x=x.Right
	}
	return x
}

//取得红黑树的极小值
func (rbt *RBtree)min (x *RBnode) *RBnode  {
	if x== rbt.NIL{
		return rbt.NIL
	}
	for x.Left != rbt.NIL{
		x=x.Left
	}
	return x
}

//搜索红黑树  有两种解决方式
func (rbt *RBtree)search(x *RBnode) *RBnode  {
	pnode := rbt.Root //根节点
	for pnode != rbt.NIL{   //循环所有节点   比我大就去左比我小就去右
		if less(pnode.Item,x.Item){
			pnode = pnode.Right
		}else if less(x.Item,pnode.Item) {
			pnode = pnode.Left
		}else {
			break //找到
		}
	}
	return pnode
}

func (rbt *RBtree) leftRotate (x *RBnode)  {
	if x.Right == rbt.NIL{
		return  //左旋 逆时针  右树不能为0
	}
	y := x.Right
	x.Right = y.Left //实现旋转的左旋
	if y.Left != rbt.NIL{
		y.Left.Parent = x //设定父亲节点
	}

	y.Parent = x.Parent //交换父节点
	if x.Parent == rbt.NIL{
		//意味着 此时此刻 根节点
		rbt.Root = y
	}else if x==x.Parent.Left{  //三步指针操作
		x.Parent.Left = y
	}else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

func (rbt *RBtree) rightRotate (x *RBnode)  {
	//因为右旋 所以左子树不可以为空
	if x.Left == nil{
		return
	}

	y := x.Left
	x.Left = y.Right
	if y.Right != rbt.NIL{
		y.Right.Parent = x   //设置祖先
	}


	y.Parent = x.Parent //y保存x的父节点
	if x.Parent == rbt.NIL{
		rbt.Root = y
	}else  if  x == x.Parent.Left {  //x小于根节点
		x.Parent.Left = y    //父亲节点的孩子是x ， 改父亲节点为子节点
	}else  {  //x 大于根节点
		x.Parent.Right = y
	}
	y.Right = x
	x.Parent = y
}

//插入一条数据
func (rbt *RBtree)Insert (item Item) *RBnode  {
	if item==nil{
		return nil
	}
	return rbt.insert(&RBnode{
		Left:   rbt.NIL,
		Right:  rbt.NIL,
		Parent: rbt.NIL,
		Color:  RED,
		Item:   item,
	})
}


//红黑树的插入算法  困难
func (rbt *RBtree) insert(z *RBnode)*RBnode  {
	//寻找插入位置
	x  := rbt.Root
	y := rbt.NIL
	for x != rbt.NIL{
		y = x  //备份位置
		if less(z.Item,x.Item){
			//如果小于的话 可以一直往左
			x= x.Left
		}else if less(x.Item,z.Item){
			x = x.Right

		}else {
			//相等
			return x
			//数据已经存在 无法插入
		}
	}

	z.Parent = y

	if y == rbt.NIL{  //根节点
		rbt.Root = z
	}else if less(z.Item,y.Item){
		y.Left = z  //小于左边插入
	}else {
		y.Right = z  //大于右边插入
	}

	rbt.count ++
	rbt.insertFixup(z)  //调整平衡
	return z



}

//插入之后调整平衡  difficult
func (rbt *RBtree)insertFixup (z *RBnode)  {
	for z.Parent.Color == RED{  //一直循环下去，直到根节点
		if z.Parent == z.Parent.Parent.Left{    //父亲节点在爷爷左边
			y := z.Parent.Parent.Right
			//旋转  再做判断
			if y.Color == RED{  //红色意味着有两个子节点 判断节点
				z.Parent.Color = BLACK
				y.Color = BLACK
				//如果我是红色节点  那么我的同级也是红色节点

				z.Parent.Parent.Color = RED //节点转换
				z = z.Parent.Parent  //平衡操作  循环前进的变量
			}else { //黑色意味着有1~2个子叶结点 无法确定
				if z == z.Parent.Right{
					//z 比父节点小   满足旋转条件
					z = z.Parent
					rbt.leftRotate(z) //左旋转
					//调整平衡
				} else {
					//比父节点大 标记红 黑  旋转
					z.Parent.Color = BLACK
					z.Parent.Parent.Color = RED
					rbt.rightRotate(z.Parent.Parent) //右旋调整平衡
				}
			}
		}else {  //父亲节点在爷爷左边
			y := z.Parent.Parent.Left
			if y.Color == RED{
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			}else {
				if z == z.Parent.Left{
					z = z.Parent
					rbt.rightRotate(z)
				}else {
					z.Parent.Color = BLACK
					z.Parent.Parent.Color = RED
					rbt.leftRotate(z.Parent.Parent) //右旋调整平衡
				}
			}
		}
	}
	rbt.Root.Color = BLACK //根节点的颜色必须是黑色
}

func (rbt *RBtree)GetDepth() int  {
	var getDeepth func(node *RBnode)int


	//函数包含
	getDeepth= func(node *RBnode) int {
		if node == nil{
			return 0
		}
		if node.Left == nil && node.Right == nil{
			return 1
		}
		var leftdeep int = getDeepth(node.Left)
		var rightdeep int = getDeepth(node.Right)
		if leftdeep > rightdeep{
			return leftdeep + 1
		}else {
			return  rightdeep +1
		}
	}
	return getDeepth(rbt.Root)
}

//第二个查找  是近似查找
func (rbt *RBtree)searchle(x *RBnode) *RBnode  {
	p := rbt.Root
	n := p  //备份根节点
	for n!=rbt.NIL{
		//如果n不等于nil  持续循环
		if less(n.Item,x.Item){
			p = n
			n = n.Right //大于
		}else if less(x.Item,n.Item){
			p = n
			n = n.Left //小于
		}else {
			return n //找到直接返回n 跳出循环
			break
		}
	}
	if less(p.Item,x.Item){
		return p
	}
	p = rbt.desuccessor(p)  //近似处理 平衡
	return p
}

func (rbt *RBtree)successor(x *RBnode) *RBnode  {
	if x == rbt.NIL{
		return rbt.NIL
	}

	if x.Right != rbt.NIL{
		return rbt.min(x.Right) //求左边最大值
	}

	y := x.Parent
	for y != rbt.NIL && x == y.Right{
		x = y
		y = y.Parent  //循环查找上一级
	}
	return y
}

func (rbt *RBtree)desuccessor(x *RBnode) *RBnode  {
	if x == rbt.NIL{
		return rbt.NIL
	}

	if x.Left != rbt.NIL{
		return rbt.max(x.Left) //求左边最大值
	}

	y := x.Parent
	for y != rbt.NIL && x == y.Left{
		x = y
		y = y.Parent  //循环查找上一级
	}
	return y
}


//封装 delete函数
func (rbt *RBtree)Delete (item Item) Item  {
	if item == nil{
		return nil
	}
	return rbt.delete(&RBnode{
		Left:   rbt.NIL,
		Right:  rbt.NIL,
		Parent: rbt.NIL,
		Color:  RED,
		Item:   item,
	})
}

func (rbt *RBtree)delete (key *RBnode) *RBnode  {
	z := rbt.search(key)
	if z == rbt.NIL{
		return rbt.NIL //无需删除
	}
	//新建节点起到备份作用
	ret := &RBnode{
		Left:   rbt.NIL,
		Right:  rbt.NIL,
		Parent: rbt.NIL,
		Color:  z.Color,
		Item:   z.Item,
	}
	var x *RBnode
	var y *RBnode
	if z.Left == rbt.NIL || z.Right == rbt.NIL{
		y = z  //直接替换删除
	}else {
		//否则重新替换删除
		y = rbt.successor(z)  //找到最接近的  右边最小
	}
	if y.Left != rbt.NIL{
		x = y.Left
	}else{
		x = y.Right
	}

	x.Parent = y.Parent

	if y.Parent == rbt.NIL{
		rbt.Root = x
	}else if y == y.Parent.Left{
		y.Parent.Left = x
	}else {
		y.Parent.Right = x
	}

	if y != z{
		z.Item = y.Item
	}
	if y.Color == BLACK{
		rbt.deleteFixUp(x)
	}
	rbt.count --
	return  ret

}

func (rbt *RBtree)deleteFixUp(x *RBnode)  {
	//红黑相间
	for x != rbt.Root && x.Color == BLACK{
		if x == x.Parent.Left{  //x在左边
			w := x.Parent.Right  //子1级节点
			if w.Color ==RED{
				w.Color = BLACK
				x.Parent.Color = RED
				rbt.leftRotate(x.Parent)
				w = x.Parent.Right  //循环步骤
			}
			if w.Left.Color == BLACK && w.Right.Color ==BLACK{
				w.Color = RED
				x = x.Parent

			}else {
				if w.Right.Color == BLACK{
					w.Left.Color = BLACK
					w.Color = RED  //同级的条件
					rbt.rightRotate(w) //右旋转
					w = x.Parent.Right //循环条件
				}
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				rbt.leftRotate(x.Parent)
				x = rbt.Root


			}
		}else {  //x在右边
			w := x.Parent.Left  //左边子节点
			if w.Color ==RED{
				w.Color = BLACK
				x.Parent.Color = RED
				rbt.rightRotate(x.Parent)
				w = x.Parent.Right  //循环步骤


			}
			if w.Left.Color == BLACK && w.Right.Color ==BLACK{
				w.Color = RED
				x = x.Parent //循环条件

			}else {
				if w.Right.Color == BLACK{
					w.Left.Color = BLACK
					w.Color = RED  //同级的条件
					rbt.leftRotate(w) //右旋转
					w = x.Parent.Left //循环条件
				}
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				rbt.rightRotate(x.Parent)
				x = rbt.Root

			}



		}
	}
	x.Color = BLACK  //循环到根节点   黑色



	return
}

