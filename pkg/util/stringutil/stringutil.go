package stringutil

import (
	"database/sql"
	"github.com/lucitez/later/pkg/util/wrappers"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func NullIfBlank(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

func PointerFromString(str string) *string {
	return &str
}

func RandomNInt(n int) string {
	rand.Seed(time.Now().Unix())

	nInt := ""

	for i := 0; i < n; i++ {
		nInt += strconv.Itoa(rand.Intn(10))
	}

	return nInt
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToNullString(str string) wrappers.NullString {
	return wrappers.NullString{
		sql.NullString{
			String: str,
			Valid:  true,
		},
	}
}
