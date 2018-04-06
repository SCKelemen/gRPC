package snowflake

import (
	"fmt"
	"os/user"
)

type Snowflake struct {
}

func GenerateSnowflake() (Snowflake, error) {
	username, err := user.Current()
	if err != nil {
		fmt.Println("oh shit bro")
	}
}
