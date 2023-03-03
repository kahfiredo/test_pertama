package main

import "fmt"

func palindromo(word string) bool {

  
  for i:=0; i < len(word)/2; i++ {

    if (word[i] != word[len(word) -1 - i] ) {
      return false
    }
        
  }
  
  return true
}

func main() {
  // result := palindromo("kasur ini rusak")
  // fmt.Println(result)

  var str string
  fmt.Printf("\n->Silahkan ketik tanpa spasi : \n")
  fmt.Scan(&str)
  fmt.Println("input : ", str)
    result := palindromo(str)
    if result == true {
    fmt.Println("result : Palindrome")
  } else {
    fmt.Println("result : Bukan Palindrome")
  }
}