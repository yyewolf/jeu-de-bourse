package ids

import "github.com/bwmarrin/snowflake"

func Parse(id string) (snowflake.ID, error) {
	return snowflake.ParseString(id)
}
