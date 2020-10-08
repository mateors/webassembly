package main

import (
    "dom"
    "strconv"
    "syscall/js"
)

/* Go version
func performCalculation() {
    firstNumber, firstNumberErr := strconv.Atoi(dom.GetString("first-number", "value"))
    secondNumber, secondNumberErr := strconv.Atoi(dom.GetString("second-number", "value"))

    if (firstNumberErr == nil && secondNumberErr == nil) {
        dom.SetValue("result", "value", strconv.Itoa(firstNumber + secondNumber))
        dom.RemoveClass("result", "error")
    } else {
        dom.SetValue("result", "value", "ERR")
        dom.AddClass("result", "error")
    }
}
*/

// Tinygo version
func performCalculation(_ js.Value, _ []js.Value) interface{} {
    firstNumber, firstNumberErr := strconv.Atoi(dom.GetString("first-number", "value"))
    secondNumber, secondNumberErr := strconv.Atoi(dom.GetString("second-number", "value"))

    if (firstNumberErr == nil && secondNumberErr == nil) {
        dom.SetValue("result", "value", strconv.Itoa(firstNumber + secondNumber))
        dom.RemoveClass("result", "error")
    } else {
        dom.SetValue("result", "value", "ERR")
        dom.AddClass("result", "error")
    }

    return nil
}

func main() {
    dom.Hide("loading")
    dom.Show("calc")

    dom.SetValue("first-number", "value", "0")
    dom.SetValue("second-number", "value", "0")
    dom.SetValue("result", "value", "0")

    dom.AddEventListener("first-number", "input", performCalculation)
    dom.AddEventListener("second-number", "input", performCalculation)

    ch := make(chan struct{})
    <-ch
}
