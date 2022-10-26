package main

import (
	"fmt"
	"reflect"
)

type router struct {
	hostname string
	mgmtIp   string
}

func (r *router) GetRouter() *router {
	return r
}

type networkSwitch struct {
	hostname string
}

func (s *networkSwitch) GetSwitch() *networkSwitch {
	return s
}

type isRouter interface {
	GetRouter() *router
}

type isSwitch interface {
	GetSwitch() *networkSwitch
}

func handleDevice(node interface{}) {
	reflectedType := reflect.TypeOf(node)
	reflectedValue := reflect.ValueOf(node)
	fmt.Println(reflectedValue.Elem())
	fmt.Println("refelected type", reflectedType)
	fmt.Println("refelected value", reflectedValue)
	reflectedTypeName := reflectedType.Name() // Will not work for pointers
	reflectedTypeKind := reflectedType.Kind()
	fmt.Println("reflectedTypeName", reflectedTypeName)
	fmt.Println("reflectedTypeKind", reflectedTypeKind)
	//fmt.Println(reflectedType.NumField()) // Only used for structs
	if reflectedTypeKind.String() == "ptr" {
		fmt.Println("Pointer Detected!!")
		fmt.Println(reflectedType.Elem())          // Only used for pointers
		fmt.Println(reflectedType.Elem().Name())   // Only used for pointers
		fmt.Println(reflectedType.Elem().Field(0)) // Only used for pointers
	}

}

func main() {
	r1 := router{hostname: "r1", mgmtIp: "192.168.1.1"}
	//sw1 := networkSwitch{hostname: "sw1"}

	handleDevice(&r1)

	//handleDevice(sw1)
}
