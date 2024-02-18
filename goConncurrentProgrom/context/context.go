package conn_context

import (
	"fmt"
	"sync"
	"time"
)

func ContextCancel() {
	// ctx, cacleFunc := context.WithCancel(context.Background())

	mutexStruct := struct {
		sync.RWMutex
		Data  map[string]interface{}
	}{
		Data: map[string]interface{}{},
	}

	v := make(chan struct{}, 1)

	v <- struct{}{}

	for i := 0; i < 10; i++ {
		go func (i int) {
			defer mutexStruct.Unlock()
			mutexStruct.Lock()
			mutexStruct.Data[fmt.Sprintf("s_%d", i)] = i
		}(i)
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer mutexStruct.RUnlock()
			mutexStruct.RLock()
			v := mutexStruct.Data[fmt.Sprintf("s_%d", i)]
			fmt.Println("v", v)
		}(i)
	}

	time.Sleep(time.Second * 3)
}