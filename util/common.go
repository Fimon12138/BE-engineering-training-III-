package util

import (
	uuid "github.com/nu7hatch/gouuid"
	"regexp"
	"tickethub_service/util/enum"
	"tickethub_service/util/log"
	"time"
)

func CheckTelephone(telephone string) bool {
	if len(telephone) != enum.LENOFTELEPHONE {
		return false
	}

	for _, char := range telephone {
		if char <= '0' || char >= '9' {
			return false
		}
	}

	return true
}

func CheckEmaliAddress(address string) (bool, error) {
	match, err := regexp.MatchString(enum.EMAILADDRESSREGEXP, address)
	return match, err
}

func ContainsString(slice []string, value string) bool {
	for _, elem := range slice {
		if elem == value {
			return true
		}
	}
	return false
}

func ContainsStringAllowZero(slice []string, value string) bool {
	if value == "" {
		return true
	}

	return ContainsString(slice, value)
}

func GetTimeFromString(str string) (time.Time, error) {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Time{}, err
	}
	tm, err := time.ParseInLocation("2006-01-02 15:04:05", str, location)
	if err != nil {
		return time.Time{}, err
	}
	return tm, nil
}

func GetDisplayTime(tm time.Time) string {
	return tm.Format("2006-01-02 15:04:05")
}

func NewUUIDString(prefix string) string {
	uuidStr := ""

	newID, err := uuid.NewV4()
	if err != nil {
		log.Errorf("Generate uuid failed %v", err)
		return uuidStr
	}
	uuidStr = newID.String()
	uuidStr = prefix + "-" + uuidStr
	return uuidStr
}
