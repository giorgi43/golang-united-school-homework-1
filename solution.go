package solution

import "github.com/kyokomi/emoji/v2"

func GetMessage() string {
	res := emoji.Sprint("Hello :world_map:!")
	return res
}
