package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


Table: content_deal_proposal_parameters_logs
[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
[ 1] content                                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 2] label                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] duration                                       INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 4] start_epoch                                    INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 5] end_epoch                                      INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 6] transfer_params                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 7] remove_unsealed_copy                           BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[ 8] skip_ip_ni_announce                            BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[ 9] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[10] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[11] requesting_api_key                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[12] system_content_deal_proposal_parameters_id     INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[13] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[14] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[15] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 11,    "content": 28,    "label": "gbgXFsMvnaClNRAqhtQdtFhNQ",    "duration": 43,    "startEpoch": 93,    "endEpoch": 15,    "transferParams": "hDsSoFfCmVRZJFXIPJYBAxGLK",    "removeUnsealedCopy": false,    "skipIpNiAnnounce": true,    "nodeInfo": "idaYRkrLBhLFHYQAZbeFIkcXi",    "requesterInfo": "IwXjAiusvSYuTSiLFiYOJTTCO",    "requestingApiKey": "TfZrGwdaEivZEPgRZouqkMoVs",    "systemContentDealProposalParametersId": 54,    "createdAt": "2051-07-18T14:58:51.12809116-04:00",    "updatedAt": "2196-07-15T21:02:52.682828298-04:00",    "deltaNodeUuid": "vXwPAuNtmYAUodvoDsZMFPAjG"}



*/

// ContentDealProposalParametersLogs struct is a row record of the content_deal_proposal_parameters_logs table in the estuary database
type ContentDealProposalParametersLogs struct {
	//[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:INT8;"`
	//[ 1] content                                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	Content null.Int `gorm:"column:content;type:INT8;"`
	//[ 2] label                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Label null.String `gorm:"column:label;type:TEXT;"`
	//[ 3] duration                                       INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	Duration null.Int `gorm:"column:duration;type:INT8;"`
	//[ 4] start_epoch                                    INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	StartEpoch null.Int `gorm:"column:start_epoch;type:INT8;"`
	//[ 5] end_epoch                                      INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	EndEpoch null.Int `gorm:"column:end_epoch;type:INT8;"`
	//[ 6] transfer_params                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	TransferParams null.String `gorm:"column:transfer_params;type:TEXT;"`
	//[ 7] remove_unsealed_copy                           BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	RemoveUnsealedCopy null.Int `gorm:"column:remove_unsealed_copy;type:BOOL;"`
	//[ 8] skip_ip_ni_announce                            BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	SkipIPNiAnnounce null.Int `gorm:"column:skip_ip_ni_announce;type:BOOL;"`
	//[ 9] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	NodeInfo null.String `gorm:"column:node_info;type:TEXT;"`
	//[10] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequesterInfo null.String `gorm:"column:requester_info;type:TEXT;"`
	//[11] requesting_api_key                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequestingAPIKey null.String `gorm:"column:requesting_api_key;type:TEXT;"`
	//[12] system_content_deal_proposal_parameters_id     INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	SystemContentDealProposalParametersID null.Int `gorm:"column:system_content_deal_proposal_parameters_id;type:INT8;"`
	//[13] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:TIMESTAMPTZ;"`
	//[14] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;"`
	//[15] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DeltaNodeUUID null.String `gorm:"column:delta_node_uuid;type:TEXT;"`
}

var content_deal_proposal_parameters_logsTableInfo = &TableInfo{
	Name: "content_deal_proposal_parameters_logs",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "content",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "Content",
			GoFieldType:        "null.Int",
			JSONFieldName:      "content",
			ProtobufFieldName:  "content",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "label",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Label",
			GoFieldType:        "null.String",
			JSONFieldName:      "label",
			ProtobufFieldName:  "label",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "duration",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "Duration",
			GoFieldType:        "null.Int",
			JSONFieldName:      "duration",
			ProtobufFieldName:  "duration",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "start_epoch",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "StartEpoch",
			GoFieldType:        "null.Int",
			JSONFieldName:      "startEpoch",
			ProtobufFieldName:  "start_epoch",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "end_epoch",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "EndEpoch",
			GoFieldType:        "null.Int",
			JSONFieldName:      "endEpoch",
			ProtobufFieldName:  "end_epoch",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "transfer_params",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "TransferParams",
			GoFieldType:        "null.String",
			JSONFieldName:      "transferParams",
			ProtobufFieldName:  "transfer_params",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "remove_unsealed_copy",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "RemoveUnsealedCopy",
			GoFieldType:        "null.Int",
			JSONFieldName:      "removeUnsealedCopy",
			ProtobufFieldName:  "remove_unsealed_copy",
			ProtobufType:       "bool",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "skip_ip_ni_announce",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "SkipIPNiAnnounce",
			GoFieldType:        "null.Int",
			JSONFieldName:      "skipIpNiAnnounce",
			ProtobufFieldName:  "skip_ip_ni_announce",
			ProtobufType:       "bool",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "node_info",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "NodeInfo",
			GoFieldType:        "null.String",
			JSONFieldName:      "nodeInfo",
			ProtobufFieldName:  "node_info",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "requester_info",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "RequesterInfo",
			GoFieldType:        "null.String",
			JSONFieldName:      "requesterInfo",
			ProtobufFieldName:  "requester_info",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "requesting_api_key",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "RequestingAPIKey",
			GoFieldType:        "null.String",
			JSONFieldName:      "requestingApiKey",
			ProtobufFieldName:  "requesting_api_key",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "system_content_deal_proposal_parameters_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "SystemContentDealProposalParametersID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "systemContentDealProposalParametersId",
			ProtobufFieldName:  "system_content_deal_proposal_parameters_id",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		{
			Index:              13,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "createdAt",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        14,
		},

		{
			Index:              14,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "updatedAt",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        15,
		},

		{
			Index:              15,
			Name:               "delta_node_uuid",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "DeltaNodeUUID",
			GoFieldType:        "null.String",
			JSONFieldName:      "deltaNodeUuid",
			ProtobufFieldName:  "delta_node_uuid",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ContentDealProposalParametersLogs) TableName() string {
	return "content_deal_proposal_parameters_logs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ContentDealProposalParametersLogs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ContentDealProposalParametersLogs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ContentDealProposalParametersLogs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ContentDealProposalParametersLogs) TableInfo() *TableInfo {
	return content_deal_proposal_parameters_logsTableInfo
}
