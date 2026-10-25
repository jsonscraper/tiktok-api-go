// Package jsonscraper provides TikTok Scraper API clients.
package jsonscraper

import (
 "context"
 "encoding/json"
 "errors"
 "fmt"
 "io"
 "net/http"
 "net/url"
 "regexp"
 "strconv"
 "strings"
 "time"
)

type APIError struct { Status int; RetryAfter string }
func (e *APIError) Error() string { return fmt.Sprintf("jsonscraper API returned HTTP %d",e.Status) }
type Options struct { BaseURL string; Timeout time.Duration; Retries int; CacheTimeout *int; HTTPClient *http.Client }
type Client struct { key string; baseURL string; retries int; cacheTimeout *int; http *http.Client }
type Response struct { Status int; Header http.Header; Body []byte }
// Decode preserves JSON numbers; decode IDs into strings or json.Number, not float64.
func (r *Response) Decode(v any) error { d:=json.NewDecoder(strings.NewReader(string(r.Body)));d.UseNumber();return d.Decode(v) }
func New(licenseKey string, options Options) (*Client,error) {
 base:=options.BaseURL
 if strings.TrimSpace(licenseKey)=="" {return nil,errors.New("a license key is required")}
 if base=="" {base="https://tiktok.evelode.com"}
 u,err:=url.Parse(base)
 if err!=nil || u.Host=="" || u.User!=nil || u.RawQuery!="" || u.Fragment!="" || (u.Scheme!="https" && !(u.Scheme=="http" && (u.Hostname()=="localhost" || u.Hostname()=="127.0.0.1" || u.Hostname()=="::1"))) {return nil,errors.New("use an HTTPS base URL")}
 if options.Retries<0 || options.Retries>5 || options.Timeout<0 {return nil,errors.New("invalid timeout or retries")}
 var hc http.Client
 if options.HTTPClient!=nil {hc=*options.HTTPClient}
 if options.Timeout>0 {hc.Timeout=options.Timeout} else if hc.Timeout==0 {hc.Timeout=30*time.Second}
 hc.CheckRedirect=func(*http.Request,[]*http.Request)error{return http.ErrUseLastResponse}
 return &Client{licenseKey,strings.TrimRight(base,"/"),options.Retries,options.CacheTimeout,&hc},nil
}
var pathPattern=regexp.MustCompile(`^/[A-Za-z0-9_/-]+$`)
func(c *Client) Request(ctx context.Context,path string,params url.Values)(*Response,error){
 if !pathPattern.MatchString(path)||strings.Contains(path,"//")||strings.Contains(path,"..") {return nil,errors.New("use an API path, not a URL")}
 q:=make(url.Values);for k,v:=range params {q[k]=append([]string(nil),v...)}
 if c.cacheTimeout!=nil && !q.Has("cache_timeout") {q.Set("cache_timeout",strconv.Itoa(*c.cacheTimeout))}
 q.Set("license_key",c.key)
 for attempt:=0;attempt<=c.retries;attempt++ {
  req,err:=http.NewRequestWithContext(ctx,"GET",c.baseURL+path+"?"+q.Encode(),nil);if err!=nil{return nil,errors.New("invalid API request")}
  req.Header.Set("Accept","application/json");req.Header.Set("User-Agent","jsonscraper-go/0.1.1")
  res,err:=c.http.Do(req);if err!=nil {return nil,errors.New("unable to reach jsonscraper API")}
  body,err:=io.ReadAll(res.Body);res.Body.Close();if err!=nil{return nil,errors.New("unable to read API response")}
  if res.StatusCode>=200 && res.StatusCode<300 {return &Response{res.StatusCode,res.Header,body},nil}
  if attempt==c.retries || (res.StatusCode!=429 && res.StatusCode!=502 && res.StatusCode!=503 && res.StatusCode!=504) {return nil,&APIError{res.StatusCode,res.Header.Get("Retry-After")}}
  delay:=time.Duration(1<<attempt)*time.Second
  if n,e:=strconv.Atoi(res.Header.Get("Retry-After"));e==nil && n>=0 {if n>60{n=60};delay=time.Duration(n)*time.Second}
  timer:=time.NewTimer(delay);select {case <-ctx.Done():timer.Stop();return nil,ctx.Err();case <-timer.C:}
 }
 return nil,errors.New("request failed")
}
