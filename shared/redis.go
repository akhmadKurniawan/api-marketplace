package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func GetDataRedis(ctx context.Context, key string, data interface{}) (interface{}, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	value, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		mars, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		err = client.Set(ctx, key, mars, time.Second*15).Err()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		fmt.Println("Save")
		return data, nil

	} else if err != nil {
		fmt.Printf("error calling redis: %v\n", err)
		return nil, err
	} else {

		err = json.Unmarshal([]byte(value), &data)
		if err != nil {
			return nil, err
		}

		fmt.Println("Done")
		return data, nil
	}
}
