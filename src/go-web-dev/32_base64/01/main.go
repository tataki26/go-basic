package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// In the cookie value it formally bans the control characters and non-ASCII characters
	// so some characters should be encoded before they are used as cookie value
	s := "Love is but a song to sing Fear's the way we die You can make the mountains ring Or make the angels cry Though the bird is on the wing And you may not know why Come on people now Smile on your brother Everybody get together Try to love one another Right now"

	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	// create base64 encoding of byte slice with setting custom encode standard
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))

	fmt.Println(s)
	fmt.Println(s64)
}
