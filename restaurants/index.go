package restaurants

import (
	"io/ioutil"
	"reflect"
	"sync"
)

var Wg sync.WaitGroup

func Init() {
	fun := Funcs()
	files, _ := ioutil.ReadDir("./restaurants")
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
