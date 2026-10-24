package main
import("context";"encoding/json";"fmt";"os";"strconv";"time";"net/url";sdk "github.com/jsonscraper/tiktok-api-go")
func main(){
 retries,_:=strconv.Atoi(os.Getenv("RETRIES"));seconds,_:=strconv.ParseFloat(os.Getenv("TIMEOUT"),64);if seconds==0{seconds=3}
 c,err:=sdk.New("contract-secret",sdk.Options{BaseURL:os.Getenv("BASE_URL"),Timeout:time.Duration(seconds*float64(time.Second)),Retries:retries})
 if err!=nil{fmt.Print(`{"error":"transport"}`);return}
 r,err:=c.Request(context.Background(),os.Getenv("API_PATH"),url.Values{"video_id":{"a&b?x=1+/Ю"},"license_key":{"override"},"count":{"0"},"flag":{"false"}})
 if err!=nil{if e,ok:=err.(*sdk.APIError);ok{b,_:=json.Marshal(map[string]any{"error":e.Status});fmt.Print(string(b))}else{fmt.Print(`{"error":"transport"}`)};return}
 os.Stdout.Write(r.Body)
}
