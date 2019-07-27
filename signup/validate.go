
package signup

import (
    "regexp"
)

func validateUsername(user string) string {
    if m, _ := regexp.MatchString("[a-zA-Z0-9._]+[@][a-z0-9.-]+", user); !m {
	return "not valid"
    }
    return "ok"
}

func validatePassword(pwd string) string {
    if len(pwd) < 6 || len(pwd) > 20 {
	return "not valid"
    }
    return "ok"
}

func validateName(name string) string {
    if m, _ := regexp.MatchString("[a-ZA-Z ]", name); !m {
	return "not valid"
    }
    return "ok"
}

func validateCheck(option string) string {
    if option != "agree" {
	return "not valid"
    }
    return"ok"
}
