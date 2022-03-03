package shared

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (a *API) getDataRedis(ctx context.Context, arg string) error {
	value, err := a.cache.Set(ctx, "ok", arg, 0).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(value)
	json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(json)
	return nil
}

type API struct {
	cache *redis.Client
}

func NewAPI() *API {
	client := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	rdb := redis.NewClient(client)

	return &API{
		cache: rdb,
	}
}
