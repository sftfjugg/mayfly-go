package entity

import "mayfly-go/pkg/model"

type MachineQuery struct {
	model.Model
	Name               string `json:"name"`
	Ip                 string `json:"ip"`         // IP地址
	Username           string `json:"username"`   // 用户名
	AuthMethod         int8   `json:"authMethod"` // 授权认证方式
	Password           string `json:"-"`
	Port               int    `json:"port"`               // 端口号
	Status             int8   `json:"status"`             // 状态 1:启用；2:停用
	Remark             string `json:"remark"`             // 备注
	SshTunnelMachineId uint64 `json:"sshTunnelMachineId"` // ssh隧道机器id
	EnableRecorder     int8   `json:"enableRecorder"`     // 是否启用终端回放记录

	TagId       uint64
	TagPathLike string
	TagIds      []uint64
}
