Getting an object from `dependencyContainer` requires to instantiate a empty struct of `type T` just to use reflection.
One way around is looping through the `dependencyContainer` and check if the `type T` is in the `dependencyContainer` and return it.

```go
// main.go

func GetV2[T any]() *T {
	for k, v := range dependencyContainer {
		unsafePointer := unsafe.Pointer(v)
		pointer := (*T)(unsafePointer)

		metadata := metadata.New[T](*pointer)

		if metadata == k {
			return pointer
		}
	}
	return nil
}
```

After benmarking the two functions, the result is as follows:

```bash
goos: linux
goarch: amd64
pkg: github.com/KafkaWannaFly/singoton/test/performance_test
cpu: 11th Gen Intel(R) Core(TM) i7-11850H @ 2.50GHz
BenchmarkGetBigStruct-16    	  116650	     10451 ns/op	   71808 B/op	      11 allocs/op
BenchmarkGetByLoop-16       	    8750	    139659 ns/op	  876090 B/op	     134 allocs/op
```

For now, reflection is the best way to get an object from `dependencyContainer`.