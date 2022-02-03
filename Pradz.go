package Pradz

import (
	"strconv"
	"strings"
)

type PradzTable struct {
	Elements []string
	Index int
	FloatVal int
}

func Init() PradzTable {
	table := PradzTable{make([]string,16),0,2}
	return table
}

func AddElement(table PradzTable, args ...interface{}) {
	if table.Index<len(table.Elements){
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
		table.Elements[table.Index]=addelem
	}
}

func RenderTable(table PradzTable) string {
	str:=""
	w:=0
	for idx, elm := range table.Elements {
		Use(idx)
		if len(elm)>w { w = len(elm) }
	}
	//w+=4
	// Now, we know maximal width + len("| ") + len(" |")
	str+="┏"+strings.Repeat("━",w+2)+"┓"+"\n" // Head
	for idx, elm := range table.Elements {
		Use(idx)
		str+="┃ "+elm+strings.Repeat(" ",w-len(elm))+" ┃"+"\n"
	} // Body
	str+="┗"+strings.Repeat("━",w+2)+"┛"+"\n" // Footer
	return str
}

func SetFloatVal(table PradzTable, val int) {
	table.FloatVal = val
}

func Use(vals ...interface{}) {
    for _, val := range vals {
        _ = val
    }
}
