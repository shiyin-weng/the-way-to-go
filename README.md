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

- chapter 9: package
  - sync.Mutex is able to only allow once process accessing data
	```golang
		import "sync"

		type Info struct {
			mu sync.Mutex
			Str string
		}

		func Update(info *Info) (
			info.mu.Lock()
			info.Str = "update"
			info.mu.Unlock()
		)
	```
  - shared buffer
	```golang
		type SyncedBuffer struct {
			lock sync.Mutex
			buffer bytes.Buffer
		}
	```

- chapter 10: struct & method
  - inherit in Go
	```golang
		type Engine struct {
			Start()
			Stop()
		}

		type Car struct {
			Engine
		}

		func (c *Car) Drive() {
			c.Start()
			c.Stop()
		}
	```

- chapter 11: interface & reflection
  - polymorphism
	```golang
	type Shaper interface {
		Area() float32
	}
	type Square struct {
		side float32
	}
	func (sq *Square) Area() float32 {
		return sq.side * sq.side
	}
	type Rectangle struct {
		length, width float32
	}
	func (r Rectangle) Area() float32 {
		return r.length * r.width
	}

	func main() {
		r := Rectangle{5, 3}
		q := &Square{5}
		shapes := []Shaper{r, q}
		for n := range shapes {
			fmt.Println("Shape details: ", shapes[n])
			fmt.Println("Area of this shape is: ", shapes[n].Area())
		}
	}
	```
  - type assertion
	```golang
	if v, ok := varI.(T); ok {
		// do something
		return
	}
	```
  - type-switch
	```golang
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
	```