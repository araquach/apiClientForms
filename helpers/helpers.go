package helpers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func RespondWithError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func ResponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

type Error struct {
	Message string `json:"message"`
}

func DateFormatYearSlash(d string) (f string) {
	s := strings.Split(d, "/")
	f = "20" + s[2] + "-" + s[1] + "-" + s[0]
	return
}

func DateFormatYearDash(d string) (f string) {
	s := strings.Split(d, "-")
	f = "20" + s[2] + "-" + s[1] + "-" + s[0]
	return
}

func DateFormatSlash(d string) (f string) {
	s := strings.Split(d, "/")
	f = s[2] + "-" + s[1] + "-" + s[0]
	return
}

// To remove the unwanted space from staff names from Phorest
func TrimSpace(n string) string {
	var res string

	split := strings.Split(n, " ")

	if len(split) > 2 {
		trimmed := strings.TrimSpace(split[0])
		toJoin := []string{trimmed, split[2]}
		res = strings.Join(toJoin, " ")
	} else {
		res = n
	}
	return res
}
