package api

import (
	"fmt"
	"strings"

	"github.com/WillKopa/boot_dev_chirpy/constants"
)

func validate_chirp(chirp string) (string, error) {
	if len(chirp) > constants.MAX_CHIRP_LENGTH  {
		return "", fmt.Errorf("chirp is too long. must be less than 141 characters")
	}

	clean_chirp := remove_profane(chirp)
	return clean_chirp, nil

}

func remove_profane(text string) string {
	profane_words := map[string]struct{} {
		"kerfuffle": {},
		"sharbert": {},
		"fornax": {},
	}
	cleaned_text := []string{}
	for _, word := range(strings.Split(text, " ")) {
		_, is_profane := profane_words[strings.ToLower(word)]
		if is_profane {
			cleaned_text = append(cleaned_text, "****")
		} else {
			cleaned_text = append(cleaned_text, word)
		}
	}

	return strings.Join(cleaned_text, " ")
}
