package Pradz

import (
	"strconv"
	"strings"
)

type PradzEntry struct {
	Entry string
	Added bool
}

type PradzTable struct {
	Elements []PradzEntry
	Index int
	FloatVal int
}

func (table *PradzTable)Init() {
	table.Elements = make([]PradzEntry,16)
	table.Index    = 0
	table.FloatVal = 2
}

func (table *PradzTable)AddElement(args ...interface{}) {
	if table.Index<cap(table.Elements){
		addelem := ""
		for idx, elm := range args {
			Use(idx)
			switch elm.(type) {
				case int64:
					addelem+=strconv.Itoa(elm.(int))
				case int32:
					addelem+=strconv.Itoa(elm.(int))
				case int:
					addelem+=strconv.Itoa(elm.(int))
				case float32:
					addelem+=strconv.FormatFloat(float64(elm.(float32)),'f',table.FloatVal,32)
				case float64:
					addelem+=strconv.FormatFloat(elm.(float64),'f',table.FloatVal,64)
				case string:
					addelem+=elm.(string)
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
	}
}

func (table *PradzTable)Render() string {
	str:=""
	w:=0
	for idx, elm := range table.Elements {
		Use(idx)
		if len(elm.Entry)>w { w = len(elm.Entry) }
	}
	//w+=4
	// Now, we know maximal width + len("| ") + len(" |")
	str+="┏"+strings.Repeat("━",w+2)+"┓"+"\n" // Head
	
	for idx, elm := range table.Elements {
		Use(idx)
		if elm.Added {
			str+="┃ "+elm.Entry+strings.Repeat(" ",w-len(elm.Entry))+" ┃"+"\n"
		}
	} // Body
	
	str+="┗"+strings.Repeat("━",w+2)+"┛"+"\n" // Footer
	return str
}

func (table *PradzTable)SetFloatVal(val int) {
	table.FloatVal = val
}

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
