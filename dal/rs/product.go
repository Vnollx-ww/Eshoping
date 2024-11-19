package rs

import (
	"Eshop/kitex_gen/product"
	"context"
	"encoding/json"
	"log"
)

func GetProductListInfo(ctx context.Context, cachekey string) ([]*product.Product, error) {
	cacheddata, err := client.Get(ctx, cachekey).Result()
	if err == nil && cacheddata != "" {
		var cachedProductList []*product.Product
		err := json.Unmarshal([]byte(cacheddata), &cachedProductList)
		if err != nil {
			log.Println("缓存反序列化失败:", err)
			client.Del(ctx, cachekey)
			return nil, err
		}
		return cachedProductList, nil
	}
	return nil, err
}
