package action

import "strings"

type Action struct {
	Action   string                 `json:"action"`
	FuncName string                 `json:"func_name"`
	Param    map[string]interface{} `json:"param,omitempty"`
}

type SwitchState string

const (
	SwitchStateUp   SwitchState = "up"
	SwitchStateDown SwitchState = "down"
)

func NewSwitchAction(a *Action, state SwitchState) *Action {
	return &Action{
		Action:   string(state),
		FuncName: a.FuncName,
		Param:    a.Param,
	}
}

// Result 操作结果
type Result struct {
	Result int    `json:"Result"`
	ErrMsg string `json:"ErrMsg"`
}

type DataResult struct {
	Total int                      `json:"total"`
	Data  []map[string]interface{} `json:"data"`
}

// IPGroup IP 组信息
type IPGroup struct {
	Id        int    `json:"id"`
	GroupName string `json:"group_name"`
	AddrPool  string `json:"addr_pool"`
	Comment   string `json:"comment"`
}

func (i *IPGroup) AddIPs(ips []string) {
	i.AddrPool = strings.Join(ips, ",")
}

func (i *IPGroup) AddComments(comments []string) {
	i.Comment = strings.Join(comments, ",")
}

type ACL struct {
	DstAddr    string `json:"dst_addr"`
	Src6Addr   string `json:"src6_addr"`
	Dst6Addr   string `json:"dst6_addr"`
	Src6Mode   int    `json:"src6_mode"`
	Dst6Mode   int    `json:"dst6_mode"`
	Src6Suffix string `json:"src6_suffix"`
	Dst6Suffix string `json:"dst6_suffix"`
	Src6Mac    string `json:"src6_mac"`
	Dst6Mac    string `json:"dst6_mac"`
	Protocol   string `json:"protocol"`
	SrcPort    string `json:"src_port"`
	DstPort    string `json:"dst_port"`
	Week       string `json:"week"`
	IPType     string `json:"ip_type"`
	ID         int    `json:"id"`
	Enabled    string `json:"enabled"`
	Comment    string `json:"comment"`
	Action     string `json:"action"`
	Dir        string `json:"dir"`
	Ctdir      int    `json:"ctdir"`
	Iinterface string `json:"iinterface"`
	Ointerface string `json:"ointerface"`
	Time       string `json:"time"`
	SrcAddr    string `json:"src_addr"`
}

func (i *ACL) SetSrcAddrIPs(ips []string) {
	i.SrcAddr = strings.Join(ips, ",")
}

func (i *ACL) GetSrcAddrIPs() []string {
	return strings.Split(i.SrcAddr, ",")
}

func (i *ACL) AddSrcAddrIPs(ips []string) {
	curiIPs := strings.Split(i.SrcAddr, ",")
	for _, ip := range ips {
		for _, curIP := range curiIPs {
			if curIP == ip {
				return
			}
		}
		curiIPs = append(curiIPs, ip)
	}
	i.SrcAddr = strings.Join(curiIPs, ",")
}

func (i *ACL) DelSrcAddrIPs(ips []string) {
	curiIPs := strings.Split(i.SrcAddr, ",")
	for _, ip := range ips {
		for i, curIP := range curiIPs {
			if curIP == ip {
				curiIPs = append(curiIPs[:i], curiIPs[i+1:]...)
				break
			}
		}
	}
	i.SrcAddr = strings.Join(curiIPs, ",")
}
