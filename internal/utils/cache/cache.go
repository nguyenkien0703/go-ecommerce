package cache

import (
	"context"
	"encoding/json"
	"example.com/go-ecommerce-backend-api/global"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func GetCache(ctx context.Context, key string, obj interface{}) error {
	rs, err := global.Rdb.Get(ctx, key).Result()
	// loi tra ve tu redis se co hai loai
	if err == redis.Nil { //1. loi ko có key nay ton tai trong redis
		return fmt.Errorf("key %s not found", key)
	} else if err != nil { // 2. neu ma no loi that su
		return err
	}
	// convert rs json to object
	if err := json.Unmarshal([]byte(rs), obj); err != nil {
		return fmt.Errorf("failed to unmarshal")
	}
	return nil
}
