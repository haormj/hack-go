package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var idx uint64 = 0
	result := make(map[string]int)
	count := 1
	for {
		fmt.Println(count)
		keys, cursor, err := rdb.Scan(idx, "", 1000).Result()
		if err != nil {
			log.Fatalln(err)
		}
		for _, k := range keys {
			prefix := ""
			splits := strings.Split(k, ":")
			switch len(splits) {
			case 1:
				prefix = "noprefix"
			default:
				prefix = strings.Join(splits[:len(splits)-1], ":")
			}
			v, ok := result[prefix]
			if ok {
				v++
				result[prefix] = v
			} else {
				result[prefix] = 1
			}
		}
		if cursor == 0 {
			break
		}
		idx = cursor
		count++
	}

	for k, v := range result {
		fmt.Printf("%50s %d\n", k, v)
	}
}
