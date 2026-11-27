package main

import (
	"fmt"

	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

func main() {
	out, _ := yaml.Marshal(map[string]string{"lang": language.English.String()})
	fmt.Print(string(out))
}
