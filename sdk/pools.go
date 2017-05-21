package sdk

import (
	"fmt"
	"encoding/json"
)

type Pool struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ModifyDate string `json:"modifydate"`
	WarnPct string `json:"warnpct"`
	SafePct string `json:"safepct"`
	MDiskGrp string `json:"safepct"`
}

type Pools struct {
	Pools []Pool  `json:"result"`
}

func (conn ActConnection) GetDiskPools(sessionId string) ([]Pool, error) {
	url := fmt.Sprintf("%s/info/lsdiskpool?sessionid=%s", conn.BaseUrl(), sessionId)
	response, err := conn.httpGet(url, nil)
	if err != nil {
		return []Pool{}, err
	}
	bytes := []byte(response)
	var pools Pools
	error := json.Unmarshal(bytes, &pools)
	return pools.Pools, error
}