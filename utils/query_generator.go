package utils

import (
	"fmt"
)

func GenerateIlikeCriteria(x string) string {
	return fmt.Sprintf("%s%s%s", "%", x, "%")
}
