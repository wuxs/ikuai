package action

// {"func_name":"acl","action":"edit","param":{ "dst_addr":"","src6_addr":"","dst6_addr":"","src6_mode":0,"dst6_mode":0,"src6_suffix":"","dst6_suffix":"","src6_mac":"","dst6_mac":"","protocol":"tcp+udp","src_port":"","dst_port":"","week":"1234567","ip_type":"4","id":2,"enabled":"yes","comment":"","action":"accept","dir":"forward","ctdir":1,"iinterface":"wan1","ointerface":"lan1","time":"00:00-23:59","src_addr":"192.168.1.0/24,111.183.45.155,1.1.1.1" }}

func NewACLEditAction(acl *ACL) *Action {
	return &Action{
		Action:   "edit",
		FuncName: "acl",
		Param: map[string]interface{}{
			"dst_addr":    acl.DstAddr,
			"src6_addr":   acl.Src6Addr,
			"dst6_addr":   acl.Dst6Addr,
			"src6_mode":   acl.Src6Mode,
			"dst6_mode":   acl.Dst6Mode,
			"src6_suffix": acl.Src6Suffix,
			"dst6_suffix": acl.Dst6Suffix,
			"src6_mac":    acl.Src6Mac,
			"dst6_mac":    acl.Dst6Mac,
			"protocol":    acl.Protocol,
			"src_port":    acl.SrcPort,
			"dst_port":    acl.DstPort,
			"week":        acl.Week,
			"ip_type":     acl.IPType,
			"id":          acl.ID,
			"enabled":     acl.Enabled,
			"comment":     acl.Comment,
			"action":      acl.Action,
			"dir":         acl.Dir,
			"ctdir":       acl.Ctdir,
			"iinterface":  acl.Iinterface,
			"ointerface":  acl.Ointerface,
			"time":        acl.Time,
			"src_addr":    acl.SrcAddr,
		},
	}
}

type ShowACLResult struct {
	Result

	Data struct {
		Data  []ACL `json:"data"`
		Total int   `json:"total"`
	} `json:"Data"`
}

func NewACLShowAction() *Action {
	return &Action{
		Action:   "show",
		FuncName: "acl",
		Param: map[string]interface{}{
			"ORDER":    "",
			"ORDER_BY": "",
			"TYPE":     "total,data",
			"limit":    "0,20",
		},
	}
}
