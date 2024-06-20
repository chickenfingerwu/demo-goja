package main

import (
	"github.com/dop251/goja"
	"os"
)

func readStrFrom(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func callingFromGo() {
	vm := goja.New()

	// callingFromGo.js
	fileName := "src/js/callingFromGo.js"
	str := readStrFrom(fileName)

	multiply := func(a, b int) int {
		return a * b
	}
	err := vm.Set("multiply", multiply)
	if err != nil {
		panic(err)
	}
	v, err := vm.RunString(str)
	if err != nil {
		panic(err)
	}
	println("calling from Go multiply(): ", v.Export().(int64))

}

func callingFromJs() {
	vm := goja.New()

	code := readStrFrom("src/js/callingFromJs.js")
	_, err := vm.RunString(code)
	if err != nil {
		panic(err)
	}
	var minus func(int, int) int
	err = vm.ExportTo(vm.Get("minus"), &minus)
	if err != nil {
		panic(err)
	}

	result := minus(10, 5)
	println("calling from Js minus(): ", result)
}

func demo() {
	vm := goja.New()

	// demo.js
	code := readStrFrom("src/js/demo.js")
	v, err := vm.RunString(code)
	if err != nil {
		panic(err)
	}

	result := v.Export().(int64)
	println("demo run simple sum(): ", result)

}

func main() {
	demo()
	callingFromJs()
	callingFromGo()
}
