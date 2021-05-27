package main

import (
	"fmt"
	"os"
	"strings"
	"flag"
	"io/ioutil"
	"path/filepath"
)

var arr []string

func permute(clist[] string) []string{
	var temp[] string
	sym := [5]string{"","@",".","_","-"}
	for _,i := range arr {
		for _,j := range clist {
			for _,k := range sym {
				if i != "" && j != "" {
					temp = append(temp,j+k+i,i+k+j)
				}
			}
		}
	}
	return temp
}

func addnum(a []string) []string{
	var temp []string
	sym := [5]string{"","@",".","_","-"}
	for _,i := range a {
		for _, k := range sym {
			temp = append(temp, i+k+"1", i+k+"12", i+k+"123", i+k+"1234", i+k+"12345")
			for j := 2000; j <= 2021; j++ {
				temp = append(temp, i+k+fmt.Sprint(j))
			}
		}
	}
	return temp
}

func main() {
	word := flag.String("w", "", "Enter your word")
	wlist := flag.String("cl", "","custom list for combination")
	minlen := flag.Int("l",6,"min length of password")
	cc := flag.String("cc","","camelCase verison of company (default: camelCase on the middle character eg: comPany)")
	flag.Parse()

	if *word != "" {
		arr = append(arr, *word)
		a := string((*word)[0])
		arr = append(arr, strings.ToUpper(a)+(*word)[1:])
		arr = append(arr, strings.ToUpper(*word))

		if *cc == "" {
			b := string((*word)[len(*word)/2])
			arr = append(arr, (*word)[:len(*word)/2]+strings.ToUpper(b)+(*word)[len(*word)/2+1:])
		} else {
			arr = append(arr, *cc) // 4 added
		}

		var common *os.File
		if *wlist != "" {
			wl1, err := os.Open(*wlist)
			if err != nil {
				cwd, err := os.Getwd()
				wl2, err := os.Open(filepath.Join(cwd,*wlist))
				if err != nil {
					panic(err)
				} else {
					common = wl2
				}
				defer wl2.Close()
			} else {
				common = wl1
			}
			defer wl1.Close()

		c, err := ioutil.ReadAll(common)
		clist := strings.Split(strings.ReplaceAll(string(c),"\r\n","\n"),"\n")

		temp := permute(clist) 

		arr = append(arr,temp...) // 12*5*4*2 added

		arr = append(arr, clist...) // 12 added
		} else {
			
			clist := []string{"admin","Admin","ADMIN","administrator","Administrator","ADMINISTRATOR","dev","Dev","DEV","password","Password","PASSWORD"}
			temp := permute(clist)
			
			arr = append(arr,temp...) // 12*5*4*2 added
			
			arr = append(arr, clist...) // 12 added
		}

		temp := addnum(arr)

		arr = append(arr,temp...) // total*5*27 added (approx. 66960)

		file, err := os.Create(*word + ".txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		wfile, err := os.OpenFile(*word+".txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer wfile.Close()

		for i := range arr {
			if len(arr[i]) >= *minlen {
				if _, err := wfile.WriteString(arr[i] + "\n"); err != nil {
					panic(err)
				}
			}
		}
	} else {
		flag.Usage()
	}
}
