package db

import (
	"fmt"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

func Select(query string, inputParams map[string]string) {
	say, err := cowsay.Say(
		query,
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
}
