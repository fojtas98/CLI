package restaurants

import (
	"embed"
	"reflect"
	"sync"
)

//go:embed *.go
var restarantsFs embed.FS
var Wg sync.WaitGroup

func Init() {
	fun := Funcs()
	files, err := restarantsFs.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		name := f.Name()
		name = name[:len(name)-3]
		if name != "map" && name != "index" {
			Wg.Add(1)
			go reflect.ValueOf(fun[name]).Call([]reflect.Value{})

		}

	}
	Wg.Wait()
}
