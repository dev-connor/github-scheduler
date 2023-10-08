package util

import "github.com/google/uuid"

func GetId(domain string) string {
	return domain + "-" + uuid.New().String()
}
