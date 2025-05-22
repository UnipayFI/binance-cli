package common

import "strings"

func IsZero(s string) bool {
	if s == "0" {
		return true
	}

	if len(s) > 1 && s[0] == '0' && s[1] == '.' {
		for _, c := range s[2:] {
			if c != '0' {
				return false
			}
		}
		return true
	}
	return false
}

func ParseArgs(args []string) (params map[string]string) {
	params = make(map[string]string, len(args))

	var key string
	for len(args) > 0 {
		arg := strings.TrimLeft(args[0], "-")
		kv := strings.Split(arg, "=")
		if len(kv) == 1 {
			if key == "" {
				key = kv[0]
			} else {
				params[key] = kv[0]
				key = ""
			}
		}
		if len(kv) == 2 {
			params[kv[0]] = kv[1]
		}

		args = args[1:]
	}
	return
}
