package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strings"
	"strconv"
	"unicode/utf8"
)

type ObjLetterFreq struct{
	letter string
	percent float64
} 

type ObjTransfer struct{
	oldLetter string
	newLetter string
} 

var words [] string
var decryt string
var  letters = "abcdefghijklmnopqrstuvwxyz"
var letterFreq [] ObjLetterFreq

var LstLTransfer [] ObjTransfer

func main() {
	
	decryt = "LQ BF SI DZ LB PK FL QT QC DM QU IN HR PB BE AC QU FD YY QS DZ PT FC XX DO UT OP QD RR OY XH QB EH AE ZU YT CA FQ MQ KH IS SI CU ZB DU XN PL TS DQ KC WM UU EB OC YP PH NT OT PK XN PL QH QB PH OV EB MH BU PT QC OC ZD TU KC XN TG KP HO HR FC EH QC MV DV RX MT SI FA HT OY ER US AB RR XI BQ TW EL BQ QI TM OR FC YD PT UB OP RO HT TH MH CL DY IK MH"
	decryt = strings.ToLower(decryt)
	fmt.Println("3809ICT sub exercise w1")
	fmt.Println(letters)
    LoadDictionary()
	LoadLetterFrequency()
	GenerateStatistics()
	PrintResults()
}

func LoadLetterFrequency(){
	file, err := os.Open("freq.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		if(strings.Contains(scanner.Text(), ",")){
			tmps := strings.Split(scanner.Text(), ",")
			if s, err := strconv.ParseFloat(tmps[1], 64); err == nil {
				AppendLetterFreq(tmps[0],s)
			}
			
		}
		
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func LoadDictionary(){
	file, err := os.Open("words.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		words = append(words,scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func AppendLetterFreq(let string, freq float64){
	l := ObjLetterFreq{letter: let,percent: freq}
	letterFreq = append(letterFreq,l)
}

func CountCharactersInString(str string) int {
    count := 0
    for len(str) > 0 {
        _, size := utf8.DecodeLastRuneInString(str)
        count++
        //fmt.Printf("%c %v\n", r, size)

        str = str[:len(str)-size]
    }
    return count
}

func GenerateStatistics(){
	max := CountCharactersInString(decryt)
	fmt.Println("Count:", max)
	for _, c := range letters{
		//fmt.Println(i, " => ", string(c))
		lcount := strings.Count(decryt, string(c))
		fmt.Println("L Count:" + string(c), lcount)
		per := float64(float64(lcount)/float64(max))*100
		newL := LowestFreq(string(c),per)

		tmpl := ObjTransfer{oldLetter: string(c),newLetter: newL}
		LstLTransfer = append(LstLTransfer,tmpl)
		//fmt.Println(l)
	}
	//fmt.Println("Done: ",complete)
}

func NumDiff(a float64, b float64) float64 {
	if a < b {
	   return b - a
	}
	return a - b
 }

func LowestFreq(l string, p float64) string{
	convert := ""
	lowest := 1000000000.0

	for _, c := range letterFreq{
		tmp := NumDiff(p,c.percent)
		if tmp <= lowest{
			convert = c.letter
			lowest = tmp
		}
	}
	return convert
}

func PrintResults(){
	complete := ""
	for _, l := range decryt{
		for _, c := range LstLTransfer{
			if string(l) == c.oldLetter{
				complete = complete + c.newLetter
				break
			}
		}
	}
	fmt.Println(complete)
}