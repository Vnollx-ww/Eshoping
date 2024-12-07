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
func GetFriendList(ctx context.Context, cachekey string) ([]*user.FriendInfo, error) {
	// 尝试从缓存获取数据
	cacheuserdata, err := client.Get(ctx, cachekey).Result()
	if err == nil && cacheuserdata != "" {
		var cacheduserfriend []user.FriendInfo
		err := json.Unmarshal([]byte(cacheuserdata), &cacheduserfriend)
		if err != nil {
			// 如果反序列化失败，删除缓存并记录错误
			log.Println("缓存反序列化失败:", err)
			client.Del(ctx, cachekey)
			return nil, err
		}
		var result []*user.FriendInfo
		for i := range cacheduserfriend {
			result = append(result, &cacheduserfriend[i])
		}
		return result, nil
	}
	return nil, err
}
