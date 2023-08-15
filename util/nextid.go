package util

import (
	"errors"
	"github.com/bwmarrin/snowflake"
)

func GetNextId() (num int64) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		err = errors.New("ID generate failed")
		return
	}
	return node.Generate().Int64()
}
