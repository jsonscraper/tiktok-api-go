//go:build ignore

package main
import (
 "context"
 "fmt"
 "os"
 "net/url"
 sdk "github.com/jsonscraper/tiktok-api-go"
)
func main() {
 client, err := sdk.NewTikTok(os.Getenv("JSONSCRAPER_TIKTOK_LICENSE_KEY"), sdk.Options{})
 if err != nil { panic(err) }
 response, err := client.GetUserInfo(context.Background(), url.Values{"query":{"therock"}, "cache_timeout":{"60"}})
 if err != nil { panic(err) }
 fmt.Println(string(response.Body))
}
