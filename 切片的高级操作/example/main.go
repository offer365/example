package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"sort"
	"unsafe"
)

var (
	a []int               // nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
	b = []int{}           // 空切片, 和 nil 不相等, 一般用来表示一个空的集合
	c = []int{1, 2, 3}    // 有3个元素的切片, len和cap都为3
	d = c[:2]             // 有2个元素的切片, len为2, cap为3
	e = c[0:2:cap(c)]     // 有2个元素的切片, len为2, cap为3
	f = c[:0]             // 有0个元素的切片, len为0, cap为3
	g = make([]int, 3)    // 有3个元素的切片, len和cap都为3
	h = make([]int, 2, 3) // 有2个元素的切片, len为2, cap为3
	i = make([]int, 0, 3) // 有0个元素的切片, len为0, cap为3
)

func append_example1() {
	a = append(a, 1)                 // 追加1个元素
	a = append(a, 1, 2, 3)           // 追加多个元素, 手写解包方式
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
}

// 在开头一般都会导致内存的重新分配，而且会导致已有的元素全部复制1次。因此，从切片的开头添加元素的性能一般要比从尾部追加元素的性能差很多。
func append_example2() {
	var a = []int{1, 2, 3}
	a = append([]int{0}, a...)          // 在开头添加1个元素
	a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片
	fmt.Println(a)
}

// 由于append函数返回新的切片，也就是它支持链式操作。我们可以将多个append操作组合起来，实现在切片中间插入元素：
func append_example3(i int, x int) {
	var a = []int{1, 2, 3}
	a = append(a[:i], append([]int{x}, a[i:]...)...) // 在第i个位置插入x
	fmt.Println(a)
	a = []int{1, 2, 3}
	a = append(a[:i], append([]int{4, 5, 6}, a[i:]...)...) // 在第i个位置插入切片
	fmt.Println(a)
	// 每个添加操作中的第二个append调用都会创建一个临时切片，并将a[i:]的内容复制到新创建的切片中，然后将临时创建的切片再追加到a[:i]。
}

func append_example4(i, x int) {
	var a = []int{1, 2, 3}
	a = append(a, 0)     // 切片扩展1个空间
	copy(a[i+1:], a[i:]) // a[i:]向后移动1个位置
	a[i] = x             // 设置新添加的元素
	fmt.Println(a)
}

func InsertOne(slice *[]int, index int, elem int) {
	if index >= len(*slice) {
		return
	}
	*slice = append(*slice, 0)
	copy((*slice)[index+1:], (*slice)[index:])
	(*slice)[index] = elem
}

func Insert(slice *[]int, i int, elem []int) {
	if len(elem) == 0 || i >= len(*slice) {
		return
	}
	*slice = append(*slice, elem...) // 为x切片扩展足够的空间
	fmt.Println(*slice)
	copy((*slice)[i+len(elem):], (*slice)[i:]) // a[i:]向后移动len(x)个位置
	copy((*slice)[i:], elem)                   // 复制新添加的切片
}

// copy 可以把第二个参数的元素拷贝到第一个元素里面
//	s1:=[]int{1,2,3}
//	s2:=[]int{4,5,6}
//	copy(s2[1:],s1)
//  [4,5,6]
//	  [1,2,3]
//  从索引是1的位置替换  4,1,2
//  copy(s2[1:],s1)
//  [4,5,6]
//	    [1,2,3]
//  从索引位置是2的位置替换  4,5,2

// 删除切片元素
//
// 根据要删除元素的位置有三种情况：从开头位置删除，从中间位置删除，从尾部删除。其中删除切片尾部的元素最快：
func delete_example1(n int) {
	var a = []int{1, 2, 3}
	a = a[:len(a)-1] // 删除尾部1个元素
	a = a[:len(a)-n] // 删除尾部N个元素
}

// 删除开头的元素可以直接移动数据指针：
func delete_example2(n int) {
	var a = []int{1, 2, 3}
	a = a[1:] // 删除开头1个元素
	a = a[n:] // 删除开头N个元素
}

// 删除开头的元素也可以不移动数据指针，但是将后面的数据向开头移动。可以用append原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）：
func delete_example3(n int) {
	var a = []int{1, 2, 3}
	a = append(a[:0], a[1:]...) // 删除开头1个元素
	a = append(a[:0], a[n:]...) // 删除开头N个元素
}

// 也可以用copy完成删除开头的元素：原地完成？？
func delete_example4(n int) {
	var a = []int{1, 2, 3}
	a = a[:copy(a, a[1:])] // 删除开头1个元素
	a = a[:copy(a, a[n:])] // 删除开头N个元素
}

// 对于删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用append或copy原地完成：
func delete_example5(i, n int) {
	var a = []int{1, 2, 3}
	a = append(a[:i], a[i+1:]...) // 删除中间1个元素
	a = append(a[:i], a[i+n:]...) // 删除中间N个元素
}
func delete_example6(i, n int) {
	var a = []int{1, 2, 3, 4, 5, 6}
	a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	a = a[:i+copy(a[i:], a[i+n:])] // 删除中间N个元素
}

// 删除开头的元素和删除尾部的元素都可以认为是删除中间元素操作的特殊情况。
//
// 切片内存技巧
//
// 在本节开头的数组部分我们提到过有类似[0]int的空数组，空数组一般很少用到。但是对于切片来说，len为0但是cap容量不为0的切片则是非常有用的特性。
// 当然，如果len和cap都为0的话，则变成一个真正的空切片，虽然它并不是一个nil值的切片。
// 在判断一个切片是否为空时，一般通过len获取切片的长度来判断，一般很少将切片和nil值做直接的比较。
//
// 比如下面的TrimSpace函数用于删除[]byte中的空格。函数实现利用了0长切片的特性，实现高效而且简洁。
func TrimSpace(s []byte) []byte {
	b := s[:0]
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}

// 其实类似的根据过滤条件原地删除切片元素的算法都可以采用类似的方式处理（因为是删除操作不会出现内存不足的情形）：
func Filter(s []byte, fn func(x byte) bool) []byte {
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}
	return b
}

// 切片高效操作的要点是要降低内存分配的次数，尽量保证append操作不会超出cap的容量，降低触发内存分配的次数和每次分配内存大小。
// 避免切片内存泄漏
// 如前面所说，切片操作并不会复制底层的数据。底层的数组会被保存在内存中，直到它不再被引用。
// 但是有时候可能会因为一个小的内存引用而导致底层整个数组处于被使用的状态，这会延迟自动内存回收器对底层数组的回收。
// 例如，FindPhoneNumber函数加载整个文件到内存，然后搜索第一个出现的电话号码，最后结果以切片方式返回。
func FindPhoneNumber(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return regexp.MustCompile("[0-9]+").Find(b)
}

// 这段代码返回的[]byte指向保存整个文件的数组。因为切片引用了整个原始数组，导致自动垃圾回收器不能及时释放底层数组的空间。
// 一个小的需求可能导致需要长时间保存整个文件数据。这虽然这并不是传统意义上的内存泄漏，但是可能会拖慢系统的整体性能
// 要修复这个问题，可以将感兴趣的数据复制到一个新的切片中（数据的传值是Go语言编程的一个哲学，虽然传值有一定的代价，但是换取的好处是切断了对原始数据的依赖）：

func FindPhoneNumber2(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = regexp.MustCompile("[0-9]+").Find(b)
	return append([]byte{}, b...)
}

type ex struct {
	a, b int
}

// 类似的问题，在删除切片元素时可能会遇到。
// 假设切片里存放的是指针对象，那么下面删除末尾的元素后，被删除的元素依然被切片底层数组引用，从而导致不能及时被自动垃圾回收器回收（这要依赖回收器的实现方式）：
func delete_example7() {
	var exs []*ex
	exs = exs[:len(exs)-1] // 被删除的最后一个元素依然被引用, 可能导致GC操作被阻碍

	// 保险的方式是先将需要自动内存回收的元素设置为nil，保证自动回收器可以发现需要回收的对象，然后再进行切片的删除操作：
	exs[len(exs)-1] = nil  // GC回收最后一个元素内存
	exs = exs[:len(exs)-1] // 从切片删除最后一个元素
	// 当然，如果切片存在的周期很短的话，可以不用刻意处理这个问题。因为如果切片本身已经可以被GC回收的话，切片对应的每个元素自然也就是可以被回收的了。
}

// 切片类型强制转换
//
// 为了安全，当两个切片类型[]T和[]Y的底层原始切片类型不同时，Go语言是无法直接转换类型的。
// 不过安全都是有一定代价的，有时候这种转换是有它的价值的——可以简化编码或者是提升代码的性能。
// 比如在64位系统上，需要对一个[]float64切片进行高速排序，我们可以将它强制转为[]int整数切片，然后以整数的方式进行排序
// （因为float64遵循IEEE754浮点数标准特性，当浮点数有序时对应的整数也必然是有序的）。
//
// 下面的代码通过两种方法将[]float64类型的切片转换为[]int类型的切片：

var af = []float64{4, 2, 5, 7, 2, 1, 88, 1}

func SortFloat64FastV1(a []float64) {
	// 强制类型转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	// 以int方式给float64排序
	sort.Ints(b)
}

func SortFloat64FastV2(a []float64) {
	// 通过 reflect.SliceHeader 更新切片头部信息实现转换
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr

	// 以int方式给float64排序
	sort.Ints(c)
}

// 第一种强制转换是先将切片数据的开始地址转换为一个较大的数组的指针，然后对数组指针对应的数组重新做切片操作。
// 中间需要unsafe.Pointer来连接两个不同类型的指针传递。
// 需要注意的是，Go语言实现中非0大小数组的长度不得超过2GB，
// 因此需要针对数组元素的类型大小计算数组的最大长度范围（[]uint8最大2GB，[]uint16最大1GB，以此类推，但是[]struct{}数组的长度可以超过2GB）。
//
// 第二种转换操作是分别取到两个不同类型的切片头信息指针，
// 任何类型的切片头部信息底层都是对应reflect.SliceHeader结构，
// 然后通过更新结构体方式来更新切片信息，从而实现a对应的[]float64切片到c对应的[]int类型切片的转换。
// 通过基准测试，我们可以发现用sort.Ints对转换后的[]int排序的性能要比用sort.Float64s排序的性能好一点。
// 不过需要注意的是，这个方法可行的前提是要保证[]float64中没有NaN和Inf等非规范的浮点数（因为浮点数中NaN不可排序，正0和负0相等，但是整数中没有这类情形）。

func main() {
	var a = []int{1, 2, 3}
	// append_example1()
	// append_example2()
	// append_example3(1,4444)
	append_example4(2, 7878)
	InsertOne(&a, 1, 23232)
	fmt.Println(a)
	a = []int{1, 2, 3}
	Insert(&a, 2, []int{4, 5, 6})
	fmt.Println(a)
}
