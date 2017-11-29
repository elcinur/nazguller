package main
import(
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "encoding/json"
)

type jsonobject struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  struct {
		Bid  float64 `json:"Bid"`
		Ask  float64 `json:"Ask"`
		Last float64 `json:"Last"`
	} `json:"result"`
}

type m_jsonobject struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		MarketName     string  `json:"MarketName"`
		High           float64 `json:"High"`
		Low            float64 `json:"Low"`
		Volume         float64 `json:"Volume"`
		Last           float64 `json:"Last"`
		BaseVolume     float64 `json:"BaseVolume"`
		TimeStamp      string  `json:"TimeStamp"`
		Bid            float64 `json:"Bid"`
		Ask            float64 `json:"Ask"`
		OpenBuyOrders  int     `json:"OpenBuyOrders"`
		OpenSellOrders int     `json:"OpenSellOrders"`
		PrevDay        float64 `json:"PrevDay"`
		Created        string  `json:"Created"`
	} `json:"result"`
}

func main(){

    base_url := "https://bittrex.com/api/v1.1/public/getticker?market="
    //base_url := "https://bittrex.com/api/v1.1/public/getticker?market=BTC-LTC"
    market_type_url :="https://bittrex.com/api/v1.1/public/getmarketsummaries"


    m_resp, m_err := http.Get(market_type_url)
    resp, err := http.Get(base_url)

    if err != nil {
        fmt.Printf("Error getting link")
    }
    if m_err != nil {
        fmt.Printf("Error getting link")
    }

    m_body, m_err := ioutil.ReadAll(m_resp.Body)
    body, err := ioutil.ReadAll(resp.Body)

    if m_err != nil {
        fmt.Printf("Error reading body")
    }
    if err != nil {
        fmt.Printf("Error reading body")
    }
    var json_var jsonobject
    var m_json_var m_jsonobject
    //var data map[string]interface{}
    json.Unmarshal(body,&json_var)
    json.Unmarshal(m_body,&m_json_var)
    fmt.Println(json_var.Result.Last)

    for i:= 0 ; i<len(m_json_var.Result) ; i++ {
        market_name := m_json_var.Result[i].MarketName
        sum_string := []string{base_url,market_name}
        current_url := strings.Join(sum_string,"")

        c_resp, err := http.Get(current_url)
        c_body, err := ioutil.ReadAll(c_resp.Body)
        if err != nil {
            fmt.Printf("Error reading body")
        }
        var c_json_var jsonobject
        json.Unmarshal(c_body,&c_json_var)
        price := c_json_var.Result.Last

        fmt.Printf("%s , %f\n",market_name,price)
    }

    defer resp.Body.Close()
    defer m_resp.Body.Close()

}
