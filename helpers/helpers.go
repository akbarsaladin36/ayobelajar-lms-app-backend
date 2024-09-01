package helpers

import "github.com/google/uuid"

func ConvertToUUID(data string) (newUUIDString string) {
	namespace := uuid.NameSpaceDNS
	customUUID := uuid.NewSHA1(namespace, []byte(data))
	uuidString := customUUID.String()

	return uuidString
}
