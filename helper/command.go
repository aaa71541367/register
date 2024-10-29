package helper

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"github.com/shillbie/register/LineThrift"
)

func MarshalMentionees(msg *LineThrift.Message, prefix ...string) (*LineThrift.Message, []string) {
	target := []string{}
	res := mentionMsg{}
	utf8Str := []rune(msg.Text)
	if msg != nil && msg.ContentMetadata != nil {
		json.Unmarshal([]byte(msg.ContentMetadata["MENTION"]), &res)
		if len(res.MENTIONEES) > 0 {
			index := 0
			midx := 0
			for no, mentionee := range res.MENTIONEES {
				if index+1 < len(utf8Str) && utf8Str[index] == ' ' {
					index += 1
				}
				s, _ := strconv.Atoi(mentionee.S)
				if s == index {
					target = append(target, mentionee.M)
					e, _ := strconv.Atoi(mentionee.E)
					index = e
					midx = no + 1
				}
			}
			if index+1 < len(utf8Str) && utf8Str[index] == ' ' {
				index += 1
			}
			res.MENTIONEES = res.MENTIONEES[midx:]
			if index > len(utf8Str) {
				index = len(utf8Str)
			}
			msg.Text = string(utf8Str[index:])
			bData, _ := json.Marshal(res)
			msg.ContentMetadata["MENTION"] = string(bData)
		}
		if len(prefix) > 0 {
			for _, c := range prefix {
				if strings.HasPrefix(msg.Text, c) && len(msg.Text) > len(c) {
					msg.Text = msg.Text[len(c):]
					return msg, target
				}
			}
			msg.Text = ""
		}
	}
	return msg, target
}

func FormatCommand(text, format string, prefix ...string) ([]string, bool) {
	format = strings.ReplaceAll(format, "[GID]", "(c[a-f0-9]{32})")
	format = strings.ReplaceAll(format, "[MID]", "(u[a-f0-9]{32})")
	format = strings.ReplaceAll(format, "[TOKEN]", "(u[a-f0-9]{32}:[a-zA-Z0-9=]{22,}..[a-zA-Z0-9+/]{27}=)")
	format = strings.ReplaceAll(format, "[SP]", "[\\s\\n:_]")
	format = strings.ReplaceAll(format, "[ALL]", "([\\s\\S]+)")
	format = strings.ReplaceAll(format, "[LINE]", "(.*)")
	format = strings.ReplaceAll(format, "[TXT]", "([^\\s\\n:_]{0,})")
	r, err := regexp.Compile(format)
	if err != nil {
		return nil, false
	}
	if len(prefix) == 0 {
		if r.MatchString(text) {
			matchs := r.FindAllStringSubmatch(text, -1)
			res := []string{}
			for _, match := range matchs {
				res = append(res, match[1:]...)
			}
			return res, true
		} else {
			return nil, false
		}
	}
	for _, p := range prefix {
		if strings.HasPrefix(strings.ToLower(text), strings.ToLower(p)) {
			if r.MatchString(text[len(p):]) {
				matchs := r.FindAllStringSubmatch(text[len(p):], -1)
				res := []string{}
				for _, match := range matchs {
					res = append(res, match[1:]...)
				}
				return res, true
			}
		}
	}
	return nil, false
}
