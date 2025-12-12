package parsinglogfiles

import "regexp"

// ^ - anchors to the beginning of the string
// \[ - matches a literal [ character (escaped because [ has special meaning in regex)
// (TRC|DBG|INF|WRN|ERR|FTL) - matches one of the six valid log levels
// \] - matches a literal ] character
func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

// < - matches a literal < character
// [~*=-]* - matches zero or more of the characters ~, *, =, or -
// > - matches a literal > character
func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

// (?i) - case-insensitive flag
// " - matches a literal double quote
// .* - matches zero or more of any character
// password - matches the literal string "password"
// .* - matches zero or more of any character
// " - matches a literal double quote
func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

// end-of-line - matches the literal string "end-of-line"
// \d+ - matches one or more digits
func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

// User - matches the literal string "User"
// \s+ - matches one or more whitespace characters
// (\S+) - captures one or more non-whitespace characters (the username)
func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	result := make([]string, len(lines))
	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			result[i] = "[USR] " + match[1] + " " + line
		} else {
			result[i] = line
		}
	}
	return result
}
