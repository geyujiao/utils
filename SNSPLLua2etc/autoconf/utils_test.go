package autoconf

import "testing"

func TestGetAllLua(t *testing.T) {
	// 现网L1 L2 lua
	// GetAllLua("xianwangL1.txt", "resultL1.txt")
	// GetAllLua("xianwangL2.txt", "resultL2.txt")
}

func TestGetNoExistLua(t *testing.T) {
	// GetNoExistLua("resultL1.txt", "resultL2.txt", "gitlua.txt", "l1.txt", "l2.txt", "l-git.txt", "l-common.txt")
	// // GetNoExistLua2("result/l1.txt", "result/l2.txt", "result/l-git.txt", "result/l1.txt", "result/l2.txt", "result/l-git.txt", "result/l-common.txt")

}

func TestBuild(t *testing.T){
	result := Build()
	println(result)
}