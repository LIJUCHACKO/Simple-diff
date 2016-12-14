package main
import (
  "fmt"
)



func matchstring (old_word string,new_word string ) ([]int,[]int) {
	i:=-1
	j:=-1
        markeroldword:=[]int{}
	markernewword:=[]int{}
 nextchar :
	i=i+1
	j=j+1
	if i==len(old_word) && j==len(new_word) {
		goto last
	} else if i>len(old_word)-1 || j>len(new_word)-1 {
		markeroldword=append(markeroldword,i)
		markernewword=append(markernewword,j)
		markeroldword=append(markeroldword,len(old_word)-1)
		markernewword=append(markernewword,len(new_word)-1)
		goto last
	}
	if old_word[i]==new_word[j] {
		goto  nextchar 
	} else {
		markeroldword=append(markeroldword,i)
		markernewword=append(markernewword,j)
		valueofi:=-1
		valueofj:=-1
		down:=1
		attempted:=true
	nextsquare :
		//searching diagonally
		for k:=i;k<=i+down;k++ {
			attempted=false
			if k<len(old_word)-1 && (i+j+down-k)<len(new_word)-1 {
			      attempted=true    
			      if old_word[k]==new_word[i+j+down-k] {
				      valueofi=k
				      valueofj=i+j+down-k  
			      } 
			}
		}
		if valueofi==-1 && valueofj==-1 && !attempted {
			markeroldword=append(markeroldword,len(old_word)-1)
			markernewword=append(markernewword,len(new_word)-1)
			goto last
		}
		down=down+1
		if valueofi==-1 && valueofj==-1 && attempted {
			goto nextsquare
		}
		i=valueofi
		j=valueofj
		if markeroldword[len(markeroldword)-1]<=i-1 {
			markeroldword=append(markeroldword,i-1)
		} else {
		        markeroldword=append(markeroldword,-1)
		}
		if markernewword[len(markernewword)-1]<=j-1 {
			markernewword=append(markernewword,j-1)
		} else {
			markernewword=append(markernewword,-1)
		}
		goto nextchar
	}
last :  
	//result format- [start1,end1,start2,end2 ...........]
	 return markeroldword,markernewword
	 
}

func printstringwithmarker (word string,marker []int ) {
        fmt.Printf("\n ")
	currentposition:=0
	fmt.Println(word)
	for i:=0;i<len(marker);i=i+2 {
	      for j:=currentposition;j<=marker[i];j++{
		    fmt.Printf(" ")
	      }
	      currentposition=marker[i]+1
	      if marker[i+1]>=marker[i]{
		      for j:=currentposition;j<=marker[i+1]+1;j++{
			    fmt.Printf("-")
		      }
		    currentposition=marker[i+1]+2
	      } else {
		      fmt.Printf("^")
		      currentposition++
	      }
	}
}

func main() {
	old_word:="testing my program"
	new_word:="test on my program"
	
	markeroldword,markernewword:=matchstring(old_word,new_word)
	
	//////RESULT///////
	fmt.Printf("old word marker-: ") 
	fmt.Println(markeroldword)
	fmt.Printf("new word marker-: ") 
	fmt.Println(markernewword)

	printstringwithmarker (old_word,markeroldword)
	printstringwithmarker (new_word,markernewword)
	fmt.Printf("\n_ is a change \n^ is a insertion \n") 
		
}