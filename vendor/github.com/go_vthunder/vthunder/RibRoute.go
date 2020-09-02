package go_vthunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type AutoGenerated struct {
	Rib Rib `json:"rib"`
}
type IPNexthopLif struct {
	DescriptionNexthopLif string `json:"description-nexthop-lif,omitempty"`
	Lif                   int    `json:"lif,omitempty"`
}
type IPNexthopIpv4 struct {
	DescriptionNexthopIP string `json:"description-nexthop-ip,omitempty"`
	IPNextHop            string `json:"ip-next-hop,omitempty"`
	DistanceNexthopIP    int    `json:"distance-nexthop-ip,omitempty"`
}
type IPNexthopTunnel struct {
	Tunnel                   int    `json:"tunnel,omitempty"`
	IPNextHopTunnel          string `json:"ip-next-hop-tunnel,omitempty"`
	DistanceNexthopTunnel    int    `json:"distance-nexthop-tunnel,omitempty"`
	DescriptionNexthopTunnel string `json:"description-nexthop-tunnel,omitempty"`
}
type IPNexthopPartition struct {
	PartitionName               string `json:"partition-name,omitempty"`
	VridNumInPartition          int    `json:"vrid-num-in-partition,omitempty"`
	DescriptionNexthopPartition string `json:"description-nexthop-partition,omitempty"`
	DescriptionPartitionVrid    string `json:"description-partition-vrid,omitempty"`
}
type RibInst struct {
	Lif             []IPNexthopLif       `json:"ip-nexthop-lif,omitempty"`
	IPNextHop       []IPNexthopIpv4      `json:"ip-nexthop-ipv4,omitempty"`
	UUID            string               `json:"uuid,omitempty"`
	IPDestAddr      string               `json:"ip-dest-addr,omitempty"`
	IPNextHopTunnel []IPNexthopTunnel    `json:"ip-nexthop-tunnel,omitempty"`
	PartitionName   []IPNexthopPartition `json:"ip-nexthop-partition,omitempty"`
	IPMask          string               `json:"ip-mask,omitempty"`
	Instance        string               `json:"-"`
}

type Rib struct {
	UUID RibInst `json:"rib,omitempty"`
}

func GetRibRoute(id string, host string, instance string) (*Rib, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/route/rib/"+instance, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var r Rib
		erro := json.Unmarshal(data, &r)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(r)
			logger.Println("[INFO] GET REQ RES..........................", r)
			return &r, nil
		}
	}
}

func PostRibRoute(id string, r Rib, host string, instance string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	logger.Println("[INFO] Headers - " + headers["Accept"] + "," + headers["Content-Type"] + "," + headers["Authorization"])

	payloadBytes, err := json.Marshal(r)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}

	url := "https://" + host + "/axapi/v3/ip/route/rib"
	logger.Println("The HTTP request url \n", url)

	resp, err := DoHttp("POST", url, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var rb Rib
		erro := json.Unmarshal(data, &rb)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}

func PutRibRoute(id string, instance string, r Rib, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(r)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}

	url := "https://" + host + "/axapi/v3/ip/route/rib/" + instance
	logger.Println("The HTTP request URL\n", url)

	resp, err := DoHttp("PUT", url, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var rb Rib
		erro := json.Unmarshal(data, &rb)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}

func DeleteRibRoute(id string, instance string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/ip/route/rib/"+instance, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var rb Rib
		erro := json.Unmarshal(data, &rb)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return err
		} else {
			fmt.Print(rb)
			logger.Println("[INFO] GET REQ RES..........................", rb)
		}
	}
	return nil
}
