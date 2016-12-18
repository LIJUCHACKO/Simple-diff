package main
import (
  "fmt"
)

func matchstring (old_word string,new_word string ) ([]int,[]int) {
	markeroldword:=[]int{}
	markernewword:=[]int{}
         matching:= [][] int{}
        old_word=old_word+" "
	new_word=new_word+" "
	for i:=0;i<=len(old_word)+1;i++ {
		row:=[]int{}
		for j:=0;j<=len(new_word)+1;j++ {
			  
			  if i< len(old_word) && j<len(new_word) {
			        if old_word[i] == new_word[j] {
				      row=append(row,1)
				} else {
				     row=append(row,0)
				}
			  } else {
				row=append(row,0)
			  }
		}
		matching=append(matching,row)
	}


	for j:=-len(new_word);j<=len(new_word);j++ {
		for i:=len(old_word)-1;i>=0;i-- {
			if j+i<len(new_word) && j+i>=0 {
				if matching[i][i+j] > 0 && matching[i+1][i+j+1] > 0 {
				      matching[i][i+j] = matching[i][i+j] + matching[i+1][i+j+1]
				}
			}
		}
	}

	for j:=-len(new_word);j<=len(new_word);j++ {
	        maxvalue:=0
		for i:=0;i<len(old_word);i++ {
			if j+i<len(new_word) && j+i>=0 {
				if matching[i][i+j] > 0  {
				      if maxvalue < matching[i][i+j] {
					    maxvalue = matching[i][i+j]
				      }
				      matching[i][i+j] = maxvalue
				} else if matching[i][i+j] ==0 {
				      maxvalue = 0
				}
			}
		}
	}
	
	
	
	skip_i_till:=-1
	opened:=false
	if  matching[0][0] ==0 {
		markeroldword=append(markeroldword,0)
		markernewword=append(markernewword,0)
		opened=true
	}
	lasty:=0
	for i:=0;i<len(old_word);i++ {
	        
		if i >= skip_i_till {
			maxj:=0
			maxvalue:=0
			for j:=0;j<len(new_word);j++ {
				if matching[i][j] > maxvalue {
				      maxj =  j
				      maxvalue = matching[i][j]
				}
			}
			if maxvalue>0  {
			      if lasty>maxj {
				    maxvalue=0
			      }
			      for k:=0;k<len(old_word);k++ {
				    if maxvalue < matching[k][maxj] {
					maxvalue=0
				    }
			      }
			}
			if maxvalue==0  {
			} else {
				lasty=maxj
				skip_i_till = i + maxvalue
				if i > 0 && opened {
					markeroldword=append(markeroldword, i-1)
					markernewword=append(markernewword, maxj-1)
					
					opened=false
				}
				//starting 
				if (i + maxvalue) < len(old_word) && (maxj + maxvalue) < len(new_word) && maxvalue>0 && !opened{
					markeroldword=append(markeroldword, i + maxvalue)
					markernewword=append(markernewword, maxj + maxvalue)
					
					opened=true
				}
				
			}
			
		}
	}
	if opened {
		markeroldword=append(markeroldword,len(old_word)-2)
	        markernewword=append(markernewword,len(new_word)-2)
	}
	//result format [start1,end1,start2,end2 ...........]
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
