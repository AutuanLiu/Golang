package main

// 引入了 encoding/base64 包并使用名称 b64代替默认的 base64。
import b64 "encoding/base64"
import "fmt"

func main() {
	// 将要编解码的字符串。
	data := "abc123!?$*&()'-=@~"
	// 编码需要使用 []byte 类型的参数，所以要将字符串转成此类型。
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	// 解码可能会返回错误，如果不确定输入信息格式是否正确，你就需要进行错误检查了。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()
	// 使用 URL 兼容的 base64 格式进行编解码。
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
