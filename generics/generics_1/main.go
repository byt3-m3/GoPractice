package main

import (
	"fmt"
)

type Router struct {
	hostname string
}

type Switch struct {
	hostname string
}

type Firewall struct {
	hostname string
}

func (f Firewall) GetHostname() string {
	return f.hostname
}

func (r Router) GetHostname() string {
	return r.hostname
}

func (s Switch) GetHostname() string {
	return s.hostname
}

type HasHostName interface {
	GetHostname() string
}

type NetworkDevices interface {
	Router | Switch
	HasHostName
}

func updateHostname[device NetworkDevices](dev device) {
	fmt.Println(fmt.Printf("%T", dev))
	fmt.Println(dev.GetHostname())

}

func main() {
	r := Router{hostname: "r1"}
	sw := Switch{hostname: "sw1"}
	fw := Firewall{hostname: "fw1"}

	updateHostname(r)
	updateHostname(sw)
	updateHostname(fw)
}
