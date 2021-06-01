/**
 * @Author : gaopeihan
 * @ClassName : main.go
 * @Date : 2021/6/1 14:46
 */
package main

import (
	"fmt"
	"net"
	"sort"
)

//非并发
//func main(){
//	for i:=8000;i<8100;i++{
//		address := fmt.Sprintf("127.0.0.1:%d",i)
//		conn,err := net.Dial("tcp",address)
//		if err != nil{
//			fmt.Printf("%s 关闭了\n",address)
//			continue
//		}
//		conn.Close()
//		fmt.Printf("%s 打开了！！\n",address)
//	}
//}

//并发
//func main(){
//	s1 := time.Now()
//	var wg sync.WaitGroup
//	for i:=1;i<65535;i++{
//		wg.Add(1)
//		go func(j int) {
//			defer wg.Done()
//			address := fmt.Sprintf("127.0.0.1:%d",j)
//			conn,err := net.Dial("tcp",address)
//			if err != nil{
//				fmt.Printf("%s 关闭了\n",address)
//				return
//			}
//			conn.Close()
//			fmt.Printf("%s 打开了！！\n",address)
//		}(i)
//
//	}
//	wg.Wait()
//	elapsed := time.Since(s1).Seconds()
//	fmt.Printf("%f\n",elapsed)
//}

//Worker池
func worker(ports chan int,results chan int){
	for p := range ports{
		address := fmt.Sprintf("127.0.0.1:%d",p)
		conn,err := net.Dial("tcp",address)
		if err!=nil{
			results <- p
			fmt.Printf("%d关闭了\n",p)
			continue
		}
		_ = conn.Close()
		fmt.Printf("%d打开了\n", p)
	}
}

func main(){

	ports := make(chan int,100)
	results := make(chan int)
	var openports []int
	//开启100个worker
	for i:=0;i<cap(ports);i++{
		go worker(ports,results)
	}
	//生产任务

	go func() {
		for j:=1;j<1024;j++{
			ports <- j
		}
	}()

	for i:=1;i<1024;i++ {
		r := <- results
		if r != 0{
			openports = append(openports,r)
		}
	}
	fmt.Println("1111111111111111111111111")
	close(ports)
	close(results)
	sort.Ints(openports)

	for _,i := range openports{
		fmt.Println(i)
	}
}