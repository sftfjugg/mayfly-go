package entity

import "mayfly-go/pkg/model"

type Mongo struct {
	model.Model

	Name               string `orm:"column(name)" json:"name"`
	Uri                string `orm:"column(uri)" json:"uri"`
	SshTunnelMachineId int    `orm:"column(ssh_tunnel_machine_id)" json:"sshTunnelMachineId"` // ssh隧道机器id
	TagId              uint64 `json:"tagId"`
	TagPath            string `json:"tagPath"`
}
