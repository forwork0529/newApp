package proxyStructs

type AppConfig struct{
	Sources []string `json:"rss"`
	ReqPer  int      `json:"request_period"`
}
