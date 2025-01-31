package sluggify

import (
	"regexp"
	"strings"
	"testing"
)

func slugify(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9-]+`)
	s = re.ReplaceAllString(s, "-")
	s = strings.ToLower(s)
	s = strings.Trim(s, "-")

	return s
}

// конец решения
func Test(t *testing.T) {
	const phrase = "Go Is Awesome!"
	const want = "go-is-awesome"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}
