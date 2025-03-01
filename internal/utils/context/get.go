package context

import (
	"context"
	"errors"
	"log"

	consts "example.com/go-ecommerce-backend-api/internal/const"
	"example.com/go-ecommerce-backend-api/internal/utils/cache"
)

type InfoUserUUID struct {
	UserId      uint64
	UserAccount string
}

func GetSubjectUUID(ctx context.Context) (string, error) {
	sUUID, ok := ctx.Value(consts.PAYLOAD_SUBJECT_UUID).(string)
	if !ok {
		return "", errors.New("failed to get subject UUID")
	}
	return sUUID, nil

}

func GetUserIdFromUUID(ctx context.Context) (uint64, error) {
	sUUID, err := GetSubjectUUID(ctx)
	log.Println("sUUID::", sUUID)
	if err != nil {
		return 0, err
	}
	// get infoUser Redis from uuid
	var inforUser InfoUserUUID

	if err := cache.GetCache(ctx, sUUID, &inforUser); err != nil {
		log.Println("err:::", err)
		return 0, err
	}
	log.Println("inforUser:::", inforUser)
	return inforUser.UserId, nil

}
