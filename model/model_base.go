package model

import "fmt"

// Action CRUD actions
type Action int32

var (
	// Create action when record is created
	Create = Action(0)

	// RetrieveOne action when a record is retrieved from db
	RetrieveOne = Action(1)

	// RetrieveMany action when record(s) are retrieved from db
	RetrieveMany = Action(2)

	// Update action when record is updated in db
	Update = Action(3)

	// Delete action when record is deleted in db
	Delete = Action(4)

	// FetchDDL action when fetching ddl info from db
	FetchDDL = Action(5)

	tables map[string]*TableInfo
)

func init() {
	tables = make(map[string]*TableInfo)

	tables["content_deal_logs"] = content_deal_logsTableInfo
	tables["content_deal_proposal_logs"] = content_deal_proposal_logsTableInfo
	tables["content_deal_proposal_parameters_logs"] = content_deal_proposal_parameters_logsTableInfo
	tables["content_logs"] = content_logsTableInfo
	tables["content_miner_logs"] = content_miner_logsTableInfo
	tables["content_wallet_logs"] = content_wallet_logsTableInfo
	tables["delta_node_geo_locations"] = delta_node_geo_locationsTableInfo
	tables["delta_startup_logs"] = delta_startup_logsTableInfo
	tables["instance_meta_logs"] = instance_meta_logsTableInfo
	tables["log_events"] = log_eventsTableInfo
	tables["piece_commitment_logs"] = piece_commitment_logsTableInfo
	tables["wallet_logs"] = wallet_logsTableInfo
}

// String describe the action
func (i Action) String() string {
	switch i {
	case Create:
		return "Create"
	case RetrieveOne:
		return "RetrieveOne"
	case RetrieveMany:
		return "RetrieveMany"
	case Update:
		return "Update"
	case Delete:
		return "Delete"
	case FetchDDL:
		return "FetchDDL"
	default:
		return fmt.Sprintf("unknown action: %d", int(i))
	}
}

// Model interface methods for database structs generated
type Model interface {
	TableName() string
	BeforeSave() error
	Prepare()
	Validate(action Action) error
	TableInfo() *TableInfo
}

// TableInfo describes a table in the database
type TableInfo struct {
	Name    string        `json:"name"`
	Columns []*ColumnInfo `json:"columns"`
}

// ColumnInfo describes a column in the database table
type ColumnInfo struct {
	Index              int    `json:"index"`
	GoFieldName        string `json:"goFieldName"`
	GoFieldType        string `json:"goFieldType"`
	JSONFieldName      string `json:"jsonFieldName"`
	ProtobufFieldName  string `json:"protobufFieldName"`
	ProtobufType       string `json:"protobufFieldType"`
	ProtobufPos        int    `json:"protobufFieldPos"`
	Comment            string `json:"comment"`
	Notes              string `json:"notes"`
	Name               string `json:"name"`
	Nullable           bool   `json:"isNullable"`
	DatabaseTypeName   string `json:"databaseTypeName"`
	DatabaseTypePretty string `json:"databaseTypePretty"`
	IsPrimaryKey       bool   `json:"isPrimaryKey"`
	IsAutoIncrement    bool   `json:"isAutoIncrement"`
	IsArray            bool   `json:"isArray"`
	ColumnType         string `json:"columnType"`
	ColumnLength       int64  `json:"columnLength"`
	DefaultValue       string `json:"defaultValue"`
}

// GetTableInfo retrieve TableInfo for a table
func GetTableInfo(name string) (*TableInfo, bool) {
	val, ok := tables[name]
	return val, ok
}
