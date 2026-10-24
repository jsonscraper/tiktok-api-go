//go:build ignore

// Get user feed by profile URL or nickname. GET /getUserFeed
// Run this standalone example: go run examples/endpoints/get_user_feed.go --help
package main
import("context";"encoding/json";"errors";"flag";"fmt";"net/url";"os";"strings";"time";sdk "github.com/jsonscraper/tiktok-api-go")
var (
 timeout=flag.Duration("timeout",30*time.Second,"Request timeout, e.g. 30s")
 retries=flag.Int("retries",0,"Optional retries, 0 to 5; may consume quota")
 output=flag.String("output","","Save JSON or binary response to a file")
 query=flag.String("query","therock","API query parameter: query")
 max_cursor=flag.String("max_cursor","","API query parameter: max_cursor")
 count=flag.String("count","","API query parameter: count")
 region=flag.String("region","","API query parameter: region")
 with_pinned_posts=flag.String("with_pinned_posts","","API query parameter: with_pinned_posts")
 sort_type=flag.String("sort_type","","API query parameter: sort_type")
 cache_timeout=flag.String("cache_timeout","60","API query parameter: cache_timeout")
)
func run() error {
 flag.Parse()
 key:=strings.TrimSpace(os.Getenv("JSONSCRAPER_TIKTOK_LICENSE_KEY"));if key==""{return errors.New("set JSONSCRAPER_TIKTOK_LICENSE_KEY before making a licensed request")}
 if *timeout<=0||*retries<0||*retries>5{return errors.New("use a positive timeout and 0 to 5 retries")}
 client,err:=sdk.NewTikTok(key,sdk.Options{Timeout:*timeout,Retries:*retries});if err!=nil{return err}
 ctx,cancel:=context.WithTimeout(context.Background(),*timeout*time.Duration(*retries+1)+time.Minute);defer cancel()
 query:=url.Values{"query":{*query}, "max_cursor":{*max_cursor}, "count":{*count}, "region":{*region}, "with_pinned_posts":{*with_pinned_posts}, "sort_type":{*sort_type}, "cache_timeout":{*cache_timeout}}
 supplied:=map[string]bool{};flag.Visit(func(f *flag.Flag){supplied[f.Name]=true})
 if !supplied["max_cursor"]{query.Del("max_cursor")}; if !supplied["count"]{query.Del("count")}; if !supplied["region"]{query.Del("region")}; if !supplied["with_pinned_posts"]{query.Del("with_pinned_posts")}; if !supplied["sort_type"]{query.Del("sort_type")}
 response,err:=client.Request(ctx,"/getUserFeed",query);if err!=nil{return err}
 if !strings.Contains(response.Header.Get("Content-Type"),"json"){
  if *output==""{return errors.New("binary response: provide --output=filename")}
  return os.WriteFile(*output,response.Body,0600)
 }
 var envelope map[string]any
 if err=response.Decode(&envelope);err!=nil{return fmt.Errorf("invalid JSON response: %w",err)}
 if envelope["status"]!="ok"||envelope["rate_limit_reached"]==true{return errors.New("API application error or quota limit; inspect the response envelope in your application")}
 rendered,err:=json.MarshalIndent(envelope,"","  ");if err!=nil{return err}
 if *output!=""{return os.WriteFile(*output,append(rendered,'\n'),0600)}
 fmt.Println(string(rendered));return nil
}
func main(){if err:=run();err!=nil{var apiError *sdk.APIError;if errors.As(err,&apiError){fmt.Fprintf(os.Stderr,"HTTP %d; Retry-After: %s\n",apiError.Status,apiError.RetryAfter)}else{fmt.Fprintln(os.Stderr,err)};os.Exit(1)}}
