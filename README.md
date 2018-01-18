# Go Tutorial
## Cơ bản về cú pháp trong Go
* Hello World trong Go

	
	```
	package main
	
	import "fmt"
	
	func main() {
		fmt.Println("Hello các bạn")
	}
	
	```
	Kết quả:
	``
	Hello các bạn
	``
	
	Chú ý: Go có hỗ trợ mặc định ``unicode``

* Vòng lặp

	Cú pháp vòng lặp for:

	```
	for initialization; condition; post {
    	// zero or more statements
	}

	```
	Ví dụ:
	
	```
	for i = 1; i <= len(os.Args); i++ {
		fmt.Println(Args[i])
	}
	```
	Có thể dùng vòng for cho range:
	
	```
	for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
	```
	Trong câu lệnh ``for range`` với vòng lặp phía trên, ``arg`` là biến chạy trong mảng ``os.Args[1:]``. lặp theo một phần mảng với ``for range``:
	
	```
	for _, arg := range os.Args[2:] {
		//
	}
	```
	
* Điều kiện trong Go
	
	```
	if condition {
		//
	} else if {
		//
	} else {
		// 
	}
	```
	Ví dụ:
	
	```
	if n > 10 {
		fmt.Println("2 chữ số")
	} else if n == 10 {
		fmt.Printlb("Nhỏ nhất có 2 chữ số")
	} else {
		
	}
	```