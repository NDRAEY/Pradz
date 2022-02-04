package Pradz

import (
	"strconv"
	"strings"
)

var LUC string = "┏"
var RUC string = "┓"
var HL  string = "━"
var VL  string = "┃"
var LLC string = "┗"
var RLC string = "┛"
var VERTICAL int = 5
var HORIZONTAL int = 6

// Types start

type PradzEntry struct {
	Entry string
	Added bool
}

type PradzTable struct {
	Elements []PradzEntry
	Index int
	FloatVal int
}

type PradzFixedTable struct {
	Elements []PradzEntry
	Index int
	FloatVal int
	Width int
	Height int
}
/*
type PradzTablet struct {
	Tables []PradzFixedTable
	Index int
	Width int
	Height int
}
*/

// Types end

func (table *PradzTable)Init() {
	table.Elements = make([]PradzEntry,16)
	table.Index    = 0
	table.FloatVal = 2
}

func (table *PradzFixedTable)Init(w int, h int) {
	table.Elements = make([]PradzEntry,16)
	table.Index = 0
	table.FloatVal = 2
	table.Width = w
	table.Height = h
}
/*
func (tablet *PradzTablet)Init(w int, h int) {
	tablet.Tables = make([]PradzFixedTable,1)
	tablet.Index = 0
	tablet.Width = w
	tablet.Height = h
}
*/
func (table *PradzTable)AddElement(args ...interface{}) {
	if table.Index<cap(table.Elements){
		addelem := ""
		for idx, elm := range args {
			Use(idx)
			switch elm.(type) {
				case int64:   addelem+=strconv.Itoa(elm.(int))
				case int32:   addelem+=strconv.Itoa(elm.(int))
				case int:     addelem+=strconv.Itoa(elm.(int))
				case float32: addelem+=strconv.FormatFloat(float64(elm.(float32)),'f',table.FloatVal,32)
				case float64: addelem+=strconv.FormatFloat(elm.(float64),'f',table.FloatVal,64)
				case string:  addelem+=elm.(string)
			}
		}
		table.Elements[table.Index].Entry=addelem
		table.Elements[table.Index].Added=true
		table.Index++
	}else{
		table.ResizeImpl(cap(table.Elements)+16)
		print("Resized table to: ")
		print(cap(table.Elements))
		print("\n")
		table.AddElement(args)
	}
}

func (table *PradzFixedTable)AddElement(args ...interface{}) {
	if table.Index<cap(table.Elements){
		addelem := ""
		for idx, elm := range args {
			Use(idx)
			switch elm.(type) {
				case int64:   addelem+=strconv.Itoa(elm.(int))
				case int32:   addelem+=strconv.Itoa(elm.(int))
				case int:     addelem+=strconv.Itoa(elm.(int))
				case float32: addelem+=strconv.FormatFloat(float64(elm.(float32)),'f',table.FloatVal,32)
				case float64: addelem+=strconv.FormatFloat(elm.(float64),'f',table.FloatVal,64)
				case string:  addelem+=elm.(string)
			}
		}
		if len(addelem)-4>table.Width {
			addelem=addelem[:len(addelem)-4]
		}
		table.Elements[table.Index].Entry=addelem
		table.Elements[table.Index].Added=true
		table.Index++
	}
}
/*
func (tablet *PradzTablet)AddTable(table PradzFixedTable) {
	if tablet.Index<cap(tablet.Tables) {
		tablet.Tables[tablet.Index] = table
		tablet.Index++
	}else{
		tablet.ResizeImpl(cap(tablet.Tables)+6)
		print("Resized tablet to: ")
		print(cap(tablet.Tables))
		print("\n")
	}
}
*/

func (table *PradzTable)Render() string {
	str:=""; w:=0
	for idx, elm := range table.Elements {
		Use(idx); if len(elm.Entry)>w { w = len(elm.Entry) }
	}
	str+="┏"+strings.Repeat("━",w+2)+"┓"+"\n"
	for idx, elm := range table.Elements {
		Use(idx); if elm.Added { str+="┃ "+elm.Entry+strings.Repeat(" ",w-len(elm.Entry))+" ┃"+"\n" }
	}
	str+="┗"+strings.Repeat("━",w+2)+"┛"+"\n"
	return str
}

func (table *PradzFixedTable)Render() string {
	str:=""; w:=0
	for idx, elm := range table.Elements {
		Use(idx); if len(elm.Entry)>w { w = len(elm.Entry) }
	}
	str+="┏"+strings.Repeat("━",w+2)+"┓"+"\n"
	for idx, elm := range table.Elements {
		Use(idx); str+="┃ "+elm.Entry+strings.Repeat(" ",w-len(elm.Entry))+" ┃"+"\n"
	}
	str+="┗"+strings.Repeat("━",w+2)+"┛"+"\n"
	return str
}
/*
func (tablet *PradzTablet)Render() string {
	str := GenRect(tablet.Width, tablet.Height)
	x:=1
	y:=0
	for idx, elm := range tablet.Tables {
		elmstr := GenRect(elm.Width, elm.Height)
		Use(idx)
		str = NLCopyAt(str,elmstr,x,y)
	}
	return str
}
*/
func (table *PradzTable)SetFloatVal(val int) { table.FloatVal = val }
func (table *PradzFixedTable)SetFloatVal(val int) { table.FloatVal = val }

func Use(vals ...interface{}) {
    for _, val := range vals {
        _ = val
    }
}

func (table *PradzTable)ResizeImpl(size int) {
	newarr := make([]PradzEntry, size)
	for idx, elm := range table.Elements {
		newarr[idx] = elm
	}
	table.Elements = newarr
}

func (table *PradzFixedTable)ResizeImpl(size int) {
	newarr := make([]PradzEntry, size)
	for idx, elm := range table.Elements {
		newarr[idx] = elm
	}
	table.Elements = newarr
}
/*
func (tablet *PradzTablet)ResizeImpl(size int) {
	newarr := make([]PradzFixedTable,size)
	for idx, elm := range tablet.Tables {
		newarr[idx] = elm
	}
	tablet.Tables = newarr
}
*/

func GenRect(w int, h int) string {
	str := LUC+strings.Repeat(HL,w-1)+RUC+"\n"
	for i:=0;i<h-2;i++ {
		str+=VL+strings.Repeat(" ",w-1)+VL+"\n"
	}
	str += LLC+strings.Repeat(HL,w-1)+RLC
	return str
}
/*
func NLCopyAt(orig string, str string, x int, y int) string {
	ostr := strings.Split(orig,"\n")
	sstr := strings.Split(str,"\n")
	tstr := ""
	fidx := 0

	for idx, elm := range sstr {
		fmt.Printf("Index: %d Cap: %d\n",idx,cap(sstr))
		if idx<len(sstr)-1 {
			tstr += CopyAt(ostr[idx],string(elm),x)+"\n"
		}else{
			tstr += string(elm)+"\n"
		}
		fidx = idx
	}

	for a, b := range ostr[fidx:] {
		Use(a)
		tstr += string(b)+"\n"
	}
	return tstr
}
*/

func CopyAt(orig string, str string, at int) string {
	strk := []byte(orig)
	/*
	if len(orig)<=len(str) {
		for i:=0;i<len(orig)-len(str);i++ {
			strk = append(strk[:], " "[:]...)
		}
	}*/
	for i:=int(at); i<at+len(str); i++ {
		strk[i]=str[i-at]
	}
	nstr := string(strk[:])
	return nstr
}
