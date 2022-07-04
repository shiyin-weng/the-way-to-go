《Go入门指南》
===================
## Fork from https://github.com/unknwon/the-way-to-go_ZH_CN for Study Purpose Only

## 开始阅读

- [目录](eBook/directory.md)


## Notes

- chapter 6: function
  - defer: function or not, implementation: close file, unlock mutual locker...
  - function: closure

- chapter 7: array and slice
  - concepts of array and slice
  	- array: [5]int{1, 2, 3, 4, 5} or [5]int{1, 2} // Fixed length, can not use append(); pass copy as function input
  	- slice: arr := make([]int, 10) // dynamic length, pass reference as function input
  - new vs make
	- new(T) assign memory, ptr:nil, len:0, cap:0, return pointer
	- make(T) ptr:[0]int, len:0, cap:0, return value
  - insert item into slice: tricky in golang // empty slice or out of range case
	- reference: /demo/slice.go

- chapter 8: map // dict in Python, hash or HashTable
  - define:
  	- var mp map[string]int
  	- mp := make(map[string]int)


