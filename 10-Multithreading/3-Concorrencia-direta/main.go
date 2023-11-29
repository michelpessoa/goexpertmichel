package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	// Declalção para Lock da variável compartilhada
	//m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Lock a utilização da variável
		//m.Lock()
		//number++
		atomic.AddUint64(&number, 1)
		// unlock a variável depois da utilização
		//m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa página %d vezes", number)))
	})

	http.ListenAndServe(":3000", nil)
}
