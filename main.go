package main

import (
	"studygo/array"
	"studygo/cancel"
	"studygo/channel"
	"studygo/client"
	"studygo/csp"
	"studygo/def"
	emptyinterface "studygo/empty_interface"
	"studygo/error"
	"studygo/extension"
	firstres "studygo/first_res"
	"studygo/fun"
	"studygo/groutine"
	"studygo/iface"
	"studygo/loop"
	"studygo/mapgo"
	"studygo/once"
	"studygo/panic"
	"studygo/polymorphism"
	"studygo/pool"
	"studygo/str"
	"studygo/structs"
	syncpool "studygo/sync_pool"
	// "github.com/jinzhu/configor"
)

func main() {
	// fmt.Println("out", configor.Config{})
	// fmt.Println("in", array.Main(123))
	array.Main()
	loop.Main()
	mapgo.Main()
	str.Main()
	fun.Main()
	def.Main()
	structs.Main()
	iface.Main()
	extension.Main()
	emptyinterface.Main()
	polymorphism.Main()
	error.Main()
	panic.Main()
	client.Main()
	groutine.Main()
	csp.Main()
	channel.Main()
	cancel.Main()
	once.Main()
	firstres.Main()
	pool.Main()
	syncpool.Main()
}
