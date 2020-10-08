package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

// var document js.Value

// func init() {
// 	document = js.Global().Get("document")
// }

// func getElementByID(elem string) js.Value {
// 	return document.Call("getElementById", elem)
// }

func main() {
	fmt.Println("Allahu akbar, Go web assembly works!")
	//js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("printMessage", js.FuncOf(printMessage))
	js.Global().Set("sum", js.FuncOf(sum))
	//js.Global().Get("document").Call("getElementById", "btnAdd").Call("addEventListener", "click", add)

	// var cb js.Func
	// cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	fmt.Println("button clicked")
	// 	//cb.Release() // release the function if the button will not be clicked again

	// 	alert := js.Global().Get("alert")
	// 	alert.Invoke("Hello")

	// 	js.Global().Call("alert", "Insufficient Crypsys")
	// 	return nil
	// })

	//js.Global().Get("document").Call("getElementById", "btnAdd").Call("addEventListener", "click", cb)
	js.Global().Get("document").Call("getElementById", "btnAdd").Call("addEventListener", "click", js.FuncOf(getNDisplay))

	<-make(chan bool)
}

func printMessage(this js.Value, inputs []js.Value) interface{} {

	//callback := inputs[len(inputs)-1:][0]
	message := inputs[0].String()

	//callback.Invoke(js.Null(), "Did you say "+message)

	return "Did you say " + message
}

func getByTagName(_ js.Value, _ []js.Value) interface{} {

	//inputElms := js.Global().Get("document").Call("getElementsByTagName", "input")
	inputElms := js.Global().Get("document").Call("getElementById", "val1")
	fmt.Println("tagsE:", inputElms, inputElms.JSValue())

	//jsonData, err := json.Marshal(inputElms)
	//fmt.Println("err:", err)
	//fmt.Println("json:", jsonData)
	//inputElms.Invoke(inputElms)

	return nil
}

func getNDisplay(_ js.Value, _ []js.Value) interface{} {

	val1 := js.Global().Get("document").Call("getElementById", "val1").Get("value")
	val2 := js.Global().Get("document").Call("getElementById", "val2").Get("value")

	fmt.Println(val1, val2)
	//fmt.Printf("%T, %T", val1, val2)
	//res := js.ValueOf(args[0].Int() + args[1].Int())
	//fmt.Println("result:", args)

	// for i, v := range args {
	// 	fmt.Println(i, v.Truthy())
	// 	fmt.Println(i, v.JSValue())
	// }

	int1, _ := strconv.Atoi(val1.String())
	int2, _ := strconv.Atoi(val2.String())
	result := int1 + int2
	fmt.Println("result:", result)

	js.Global().Call("alert", result)

	return nil

}

// func addOld(i []js.Value) {

// 	input := js.ValueOf(i[0].Int())
// 	js.Global().Set("output", input)
// 	fmt.Println(input)

// }

//function passes by id from js ==> add('val1', 'val2');
func add(_ js.Value, _ []js.Value) interface{} {

	value1 := js.Global().Get("document").Call("getElementById", "val1").Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", "val2").Get("value").String()
	//js.Global().Set("output", value1+value2)
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	return int1 + int2
}

//function passes by value from js ==> sum(5, 5);
func sum(this js.Value, args []js.Value) interface{} {

	value1 := args[0].Int()
	value2 := args[1].Int()
	//return js.ValueOf(i[0].Int() + i[1].Int())
	sum := value1 + value2

	return js.ValueOf(sum)
}

// exposing to JS
//js.Global().Set("add", js.FuncOf(add))
//
//js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)
