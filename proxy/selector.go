package proxy

import (
	C "github.com/Dreamacro/clash/constant"
)

type SelectorInfo struct {
	All     []string      `json:"all,omitempty"`
	History []interface{} `json:"history,omitempty"`
	Name    string        `json:"name"`
	Now     string        `json:"now"`
	Type    string        `json:"type"`
}

type Proxy struct {
	Name		string
	Type		C.AdapterType
	Parent		*Proxy
	//Children 	*list.List
	Delay		int16
}

type GroupsList struct {
	Name    string   `json:"name"`
	Proxies []string `json:"proxies"`
	Type    string   `json:"type"`
}

//func (proxy *Proxy) AddChild(child *Proxy) {
//	proxy.Children.PushBack(child)
//}
