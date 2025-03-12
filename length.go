package main

import (
	"fmt"
	"strings"
)
  func lengthofLastWord(a string) int {
   words:= strings.Fields(a)
      if len(words)==0{
		return 0
	  }
	  return len(words[len(words)-1])
	  
 }
   func main (){
	   a:= "Hello World"
	   result:= (lengthofLastWord(a))
	   fmt.Println(result)
   }