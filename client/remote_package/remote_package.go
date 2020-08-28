package remote

import (
	cm "github.com/easierway/concurrent_map"
)

// Main is a function
func Main() {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	// fmt.Println(m.Get(cm.StrKey("key")))
}
