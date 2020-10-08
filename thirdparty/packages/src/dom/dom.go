package dom

import (
    "syscall/js"
)

var document js.Value

func init() {
    document = js.Global().Get("document")
}

func getElementById(elem string) js.Value {
    return document.Call("getElementById", elem)
}

func getElementValue(elem string, value string) js.Value {
    return getElementById(elem).Get(value)
}

func Hide(elem string) {
    getElementValue(elem, "style").Call("setProperty", "display", "none")
}

func Show(elem string) {
    getElementValue(elem, "style").Call("setProperty", "display", "block")
}

func GetString(elem string, value string) string {
    return getElementValue(elem, value).String()
}

func SetValue(elem string, key string, value string) {
    getElementById(elem).Set(key, value)
}

func AddClass(elem string, class string) {
    getElementValue(elem, "classList").Call("add", class)
}

func RemoveClass(elem string, class string) {
    classList := getElementValue(elem, "classList")
    if (classList.Call("contains", class).Bool()) {
        classList.Call("remove", class)
    }
}

/* Standard Go Version
func wrapGoFunction(fn func()) func(js.Value, []js.Value) interface {} {
    return func(_ js.Value, _ []js.Value) interface {} {
        fn()
        return nil
    }
}

func AddEventListener(elem string, event string, fn func()) {
    getElementById(elem).Call("addEventListener", event, js.FuncOf(wrapGoFunction(fn)))
}
*/

// Tinygo version
func AddEventListener(elem string, event string, fn func(js.Value, []js.Value) interface{}) {
    getElementById(elem).Call("addEventListener", event, js.FuncOf(fn))
}
