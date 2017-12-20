package utils

import (
	 "regexp"
)

	func CheckMobile(mobile string) bool {
		if m, _ := regexp.MatchString(`^(1[3|4|5|7|8][0-9]\d{4,8})$`, mobile); !m {
			return false
		}
		return true
	}

	func CheckEmail(email string) bool {
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
			return false
		}
		return true
	}

