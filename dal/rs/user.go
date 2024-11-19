package rs

import (
	"Eshop/kitex_gen/user"
	"context"
	"encoding/json"
	"log"
)

func GetUserInfo(ctx context.Context, cachekey string) (*user.User, error) {
	cacheuserdata, err := client.Get(ctx, cachekey).Result()
	if err == nil && cacheuserdata != "" {
		var cacheduser user.User
		err := json.Unmarshal([]byte(cacheuserdata), &cacheduser)
		if err != nil {
			log.Println("缓存反序列化失败:", err)
			client.Del(ctx, cachekey)
			return nil, err
		}
		return &cacheduser, nil
	}
	return nil, err
}
