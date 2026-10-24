//go:build ignore

// Search for music by keyword. GET /searchMusic
// Run this standalone example: go run examples/endpoints/search_music.go --help
package main
import("context";"encoding/json";"errors";"flag";"fmt";neturl "net/url";"os";"strings";"time";sdk "github.com/jsonscraper/tiktok-api-go")
var (
 timeout=flag.Duration("timeout",30*time.Second,"Request timeout, e.g. 30s")
 retries=flag.Int("retries",0,"Optional retries, 0 to 5; may consume quota")
 output=flag.String("output","","Save JSON or binary response to a file")
 keyword=flag.String("keyword","apple","API query parameter: keyword")
 cursor=flag.String("cursor","","API query parameter: cursor")
 count=flag.String("count","","API query parameter: count")
 cache_timeout=flag.String("cache_timeout","60","API query parameter: cache_timeout")
 filter_by=flag.String("filter_by","","API query parameter: filter_by")
 sort_type=flag.String("sort_type","","API query parameter: sort_type")
 region=flag.String("region","","API query parameter: region")
)
func run() error {
 flag.Parse()
 key:=strings.TrimSpace(os.Getenv("JSONSCRAPER_TIKTOK_LICENSE_KEY"));if key==""{return errors.New("set JSONSCRAPER_TIKTOK_LICENSE_KEY before making a licensed request")}
 if *timeout<=0||*retries<0||*retries>5{return errors.New("use a positive timeout and 0 to 5 retries")}
 client,err:=sdk.NewTikTok(key,sdk.Options{Timeout:*timeout,Retries:*retries});if err!=nil{return err}
 ctx,cancel:=context.WithTimeout(context.Background(),*timeout*time.Duration(*retries+1)+time.Minute);defer cancel()
 query:=neturl.Values{"keyword":{*keyword}, "cursor":{*cursor}, "count":{*count}, "cache_timeout":{*cache_timeout}, "filter_by":{*filter_by}, "sort_type":{*sort_type}, "region":{*region}}
 supplied:=map[string]bool{};flag.Visit(func(f *flag.Flag){supplied[f.Name]=true})
 if !supplied["cursor"]{query.Del("cursor")}; if !supplied["count"]{query.Del("count")}; if !supplied["filter_by"]{query.Del("filter_by")}; if !supplied["sort_type"]{query.Del("sort_type")}; if !supplied["region"]{query.Del("region")}
 response,err:=client.Request(ctx,"/searchMusic",query);if err!=nil{return err}
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
