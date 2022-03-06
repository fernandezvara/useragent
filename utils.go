package useragent

import "strings"

func (u *UserAgent) includes(item string) bool {
	_, ok := u.partsMap[item]
	return ok
}

func (u *UserAgent) getVersionString(search string) string {

	var (
		index, index2 int
		rest          string
	)

	if index = strings.Index(u.input, search); index == -1 {
		return ""
	}

	index += len(search) // starts: index + lenght
	rest = u.input[index:]
	if index2 = strings.Index(rest, " "); index2 == -1 {
		return ""
	}

	index2 += index // ends: start + length
	return cleanVersion(u.input[index:index2])

}

func partsToMap(input string) map[string]string {

	var (
		arr   []string
		parts = make(map[string]string)
	)

	arr = strings.FieldsFunc(input, splitFunc)

	for _, part := range arr {
		// has version?
		if strings.Contains(part, "/") {
			this := strings.Split(part, "/")
			if len(this) == 2 {
				parts = set(parts, this[0], cleanVersion(this[1]))
			}
		} else {
			parts = set(parts, part, "")
		}

	}

	return parts
}

// set verifies if the part already exists and if exists but is blank then sets it
func set(data map[string]string, key, value string) map[string]string {

	if _, ok := data[key]; ok {
		// set it only if blank
		if data[key] == "" {
			data[key] = value
		}
	} else {
		data[key] = value
	}

	return data
}

// split on each parenthesis, ; or space
func splitFunc(r rune) bool {
	return strings.ContainsRune(" ();", r)
}

func cleanVersionFunc(r rune) bool {
	return strings.ContainsRune("0123456789._", r)
}

func cleanVersion(input string) (output string) {
	for _, char := range input {
		if cleanVersionFunc(char) {
			// if version contains _ instead of dots
			if string(char) == "_" {
				output += string(".")
			} else {
				output += string(char)
			}
		}
	}
	output = removeTrailingDots(output)
	return
}

func removeTrailingDots(input string) string {

	for {
		if strings.HasPrefix(input, ".") {
			input = input[1:]
		} else {
			break
		}
	}

	for {
		if strings.HasSuffix(input, ".") {
			input = input[:len(input)-1]
		} else {
			break
		}
	}

	return input
}
