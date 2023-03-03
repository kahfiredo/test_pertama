package main

import (
	"fmt"
	"sort"
)

func main(){
   fmt.Printf("Masukin jumlah digit angka : ")
   var size int
   fmt.Scanln(&size)
   var arr = make([]int, size)
   for i:=0; i<size; i++ {
      fmt.Printf("Angka %d : ", i)
      fmt.Scanf("%d", &arr[i])
   }
   fmt.Println("input : ", arr)
   sort.Ints(arr)
   fmt.Println("result : ",arr) 
   
}