package mainз

import (
	"Cash/storage"
	"fmt"
	"time"
)

type Cache interface {
	Get(key string) (value any, ok bool)
	Set(key string, value any, ttl time.Duration)
}

func main() {

	key := "Egor"
	val := 777
	t := 3 * time.Second

	myMap := make(map[string]any)

	st := storage.New(myMap)

	st.Set(key, val, t)

	time.Sleep(2 * time.Second)

	fmt.Println(st.Get(key))

}
