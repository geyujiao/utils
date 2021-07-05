package study

func c10k()  {

	// 进程&线程(apych)



	// 协程---goroutine : go fun(){}   channel----多个goroutine间数据通讯

}
// 并发&并行
func bingfa()  {
	// 区别
	// 并发：同一时刻，两个任务交替运行
	// 并行：同一时刻，两个任务同时运行


}

type Foo struct {
	C string
}

type Bar struct {
	Foo
}

func (f *Foo) Set(c string)  {

}

func (b *Bar)SetBar()  {
	b.C = ""
}