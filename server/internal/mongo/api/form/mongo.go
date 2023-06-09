package form

type Mongo struct {
	Id                 uint64
	Uri                string `binding:"required" json:"uri"`
	SshTunnelMachineId int    `json:"sshTunnelMachineId"` // ssh隧道机器id
	Name               string `binding:"required" json:"name"`
	TagId              uint64 `binding:"required" json:"tagId"`
	TagPath            string `json:"tagPath"`
}

type MongoCommand struct {
	Database   string                 `binding:"required" json:"database"`
	Collection string                 `binding:"required" json:"collection"`
	Filter     map[string]interface{} `json:"filter"`
}

type MongoRunCommand struct {
	Database string                 `binding:"required" json:"database"`
	Command  map[string]interface{} `json:"command"`
}

type MongoFindCommand struct {
	MongoCommand
	Sort  map[string]interface{} `json:"sort"`
	Skip  int64
	Limit int64
}

type MongoUpdateByIdCommand struct {
	MongoCommand
	DocId  interface{}            `binding:"required" json:"docId"`
	Update map[string]interface{} `json:"update"`
}

type MongoInsertCommand struct {
	MongoCommand
	Doc map[string]interface{} `json:"doc"`
}
