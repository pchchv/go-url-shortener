package scripts

import (
	"math/rand"
	"strings"

	"github.com/asaskevich/govalidator"
)

func shortener(url string) string {
	if govalidator.IsRequestURL(url) == true {
		if strings.Contains(url, "://") {
			url = url[strings.Index(url, "://")+3:]
		}
		return url[:rand.Intn(4)]
	}
}
