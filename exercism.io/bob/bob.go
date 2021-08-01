/*Package bob implements ackadaisical teenager.

In conversation, his responses are very limited.
Bob answers 'Sure.' if you ask him a question, such as "How are you?".
He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
He says 'Fine. Be that way!' if you address him without actually saying anything.
He answers 'Whatever.' to anything else.
*/
package bob

import (
	"strings"
)

// Hey takes remark and return bob responce to conversation.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case remark == "":
		return "Fine. Be that way!"
	case isYelling(remark) && isQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case isYelling(remark):
		return "Whoa, chill out!"
	case isQuestion(remark):
		return "Sure."
	default:
		return "Whatever."
	}

}

// isQuestion checks if string ends with ?
func isQuestion(s string) bool {
	return strings.HasSuffix(s, "?")
}

// isYelling checks if string is valid YELL
func isYelling(s string) bool {
	nonYeling := true

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			return false
		}

		if v >= 'A' && v <= 'Z' {
			nonYeling = false
		}
	}

	if nonYeling {
		return false
	}

	return true
}
