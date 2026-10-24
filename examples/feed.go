//go:build ignore

// Run this file alone with: go run examples/feed.go
package main
import("context";"fmt";"os";"net/url";sdk "github.com/jsonscraper/tiktok-api-go")
func main(){c,e:=sdk.NewTikTok(os.Getenv("JSONSCRAPER_TIKTOK_LICENSE_KEY"),sdk.Options{});if e!=nil{panic(e)};r,e:=c.Request(context.Background(),"/getUserFeed",url.Values{"query":{"therock"}, "count":{"10"}, "max_cursor":{""}, "cache_timeout":{"60"}});if e!=nil{panic(e)};fmt.Println(string(r.Body))}
