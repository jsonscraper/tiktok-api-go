//go:build ignore

// Run this file alone with: go run examples/video.go
package main
import("context";"fmt";"os";"net/url";sdk "github.com/jsonscraper/tiktok-api-go")
func main(){c,e:=sdk.NewTikTok(os.Getenv("JSONSCRAPER_TIKTOK_LICENSE_KEY"),sdk.Options{});if e!=nil{panic(e)};r,e:=c.Request(context.Background(),"/getVideoByID",url.Values{"video_id":{"7106855913906081070"}, "cache_timeout":{"60"}});if e!=nil{panic(e)};fmt.Println(string(r.Body))}
