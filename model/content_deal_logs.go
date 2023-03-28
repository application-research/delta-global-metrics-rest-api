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


Table: content_deal_logs
[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
[ 1] content                                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 2] prop_cid                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] deal_uuid                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] miner                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] deal_id                                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 6] failed                                         BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[ 7] verified                                       BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[ 8] slashed                                        BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[ 9] failed_at                                      TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[10] dt_chan                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[11] transfer_started                               TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[12] transfer_finished                              TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[13] on_chain_at                                    TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[14] sealed_at                                      TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[15] last_message                                   TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[16] deal_protocol_version                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[17] miner_version                                  TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[18] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[19] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[20] requesting_api_key                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[21] system_content_deal_id                         INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[22] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[23] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[24] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 99,    "content": 55,    "propCid": "xknsQTytaZWgYojHSDkHVfPHa",    "dealUuid": "GxKYowuykjCRKgEIOgeQqqsen",    "miner": "OInOCDbSbQxplmMthtVnPierV",    "dealId": 85,    "failed": true,    "verified": false,    "slashed": false,    "failedAt": "2104-07-28T14:23:06.964124494-04:00",    "dtChan": "AcIPDCIUYkjdwjTvidljlPmEx",    "transferStarted": "2255-11-25T21:46:23.259616019-05:00",    "transferFinished": "2185-12-19T08:50:07.455413034-05:00",    "onChainAt": "2029-02-05T01:46:10.623356801-05:00",    "sealedAt": "2028-02-10T19:11:30.660121509-05:00",    "lastMessage": "vNFrivKDIRlSgyxmVrQZtcNPp",    "dealProtocolVersion": "SrvIxXavVsTojSXZWVUyHfVsG",    "minerVersion": "dZasDhnmXPWxsTJTkkjqsXekw",    "nodeInfo": "eaPQgRleeUStfJjFeLaqbMGAl",    "requesterInfo": "IPnKALNplIsAFbfNNEJjxEEbV",    "requestingApiKey": "GfxZduGGhytByYTANrfWicIOx",    "systemContentDealId": 68,    "createdAt": "2236-06-18T16:44:22.565027633-04:00",    "updatedAt": "2108-07-24T20:06:59.490583119-04:00",    "deltaNodeUuid": "DBoqKOqeVRvCTTBXExXyHoKfB"}



*/

// ContentDealLogs struct is a row record of the content_deal_logs table in the estuary database
type ContentDealLogs struct {
	//[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:INT8;"`
	//[ 1] content                                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	Content null.Int `gorm:"column:content;type:INT8;"`
	//[ 2] prop_cid                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	PropCid null.String `gorm:"column:prop_cid;type:TEXT;"`
	//[ 3] deal_uuid                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DealUUID null.String `gorm:"column:deal_uuid;type:TEXT;"`
	//[ 4] miner                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Miner null.String `gorm:"column:miner;type:TEXT;"`
	//[ 5] deal_id                                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	DealID null.Int `gorm:"column:deal_id;type:INT8;"`
	//[ 6] failed                                         BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	Failed null.Int `gorm:"column:failed;type:BOOL;"`
	//[ 7] verified                                       BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	Verified null.Int `gorm:"column:verified;type:BOOL;"`
	//[ 8] slashed                                        BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	Slashed null.Int `gorm:"column:slashed;type:BOOL;"`
	//[ 9] failed_at                                      TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	FailedAt null.Time `gorm:"column:failed_at;type:TIMESTAMPTZ;"`
	//[10] dt_chan                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DtChan null.String `gorm:"column:dt_chan;type:TEXT;"`
	//[11] transfer_started                               TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	TransferStarted null.Time `gorm:"column:transfer_started;type:TIMESTAMPTZ;"`
	//[12] transfer_finished                              TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	TransferFinished null.Time `gorm:"column:transfer_finished;type:TIMESTAMPTZ;"`
	//[13] on_chain_at                                    TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	OnChainAt null.Time `gorm:"column:on_chain_at;type:TIMESTAMPTZ;"`
	//[14] sealed_at                                      TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	SealedAt null.Time `gorm:"column:sealed_at;type:TIMESTAMPTZ;"`
	//[15] last_message                                   TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	LastMessage null.String `gorm:"column:last_message;type:TEXT;"`
	//[16] deal_protocol_version                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DealProtocolVersion null.String `gorm:"column:deal_protocol_version;type:TEXT;"`
	//[17] miner_version                                  TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	MinerVersion null.String `gorm:"column:miner_version;type:TEXT;"`
	//[18] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	NodeInfo null.String `gorm:"column:node_info;type:TEXT;"`
	//[19] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequesterInfo null.String `gorm:"column:requester_info;type:TEXT;"`
	//[20] requesting_api_key                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequestingAPIKey null.String `gorm:"column:requesting_api_key;type:TEXT;"`
	//[21] system_content_deal_id                         INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	SystemContentDealID null.Int `gorm:"column:system_content_deal_id;type:INT8;"`
	//[22] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:TIMESTAMPTZ;"`
	//[23] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;"`
	//[24] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DeltaNodeUUID null.String `gorm:"column:delta_node_uuid;type:TEXT;"`
}

var content_deal_logsTableInfo = &TableInfo{
	Name: "content_deal_logs",
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
			Name:               "prop_cid",
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
			GoFieldName:        "PropCid",
			GoFieldType:        "null.String",
			JSONFieldName:      "propCid",
			ProtobufFieldName:  "prop_cid",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "deal_uuid",
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
			GoFieldName:        "DealUUID",
			GoFieldType:        "null.String",
			JSONFieldName:      "dealUuid",
			ProtobufFieldName:  "deal_uuid",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "miner",
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
			GoFieldName:        "Miner",
			GoFieldType:        "null.String",
			JSONFieldName:      "miner",
			ProtobufFieldName:  "miner",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "deal_id",
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
			GoFieldName:        "DealID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "dealId",
			ProtobufFieldName:  "deal_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "failed",
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
			GoFieldName:        "Failed",
			GoFieldType:        "null.Int",
			JSONFieldName:      "failed",
			ProtobufFieldName:  "failed",
			ProtobufType:       "bool",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "verified",
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
			GoFieldName:        "Verified",
			GoFieldType:        "null.Int",
			JSONFieldName:      "verified",
			ProtobufFieldName:  "verified",
			ProtobufType:       "bool",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "slashed",
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
			GoFieldName:        "Slashed",
			GoFieldType:        "null.Int",
			JSONFieldName:      "slashed",
			ProtobufFieldName:  "slashed",
			ProtobufType:       "bool",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "failed_at",
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
			GoFieldName:        "FailedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "failedAt",
			ProtobufFieldName:  "failed_at",
			ProtobufType:       "uint64",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "dt_chan",
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
			GoFieldName:        "DtChan",
			GoFieldType:        "null.String",
			JSONFieldName:      "dtChan",
			ProtobufFieldName:  "dt_chan",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "transfer_started",
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
			GoFieldName:        "TransferStarted",
			GoFieldType:        "null.Time",
			JSONFieldName:      "transferStarted",
			ProtobufFieldName:  "transfer_started",
			ProtobufType:       "uint64",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "transfer_finished",
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
			GoFieldName:        "TransferFinished",
			GoFieldType:        "null.Time",
			JSONFieldName:      "transferFinished",
			ProtobufFieldName:  "transfer_finished",
			ProtobufType:       "uint64",
			ProtobufPos:        13,
		},

		{
			Index:              13,
			Name:               "on_chain_at",
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
			GoFieldName:        "OnChainAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "onChainAt",
			ProtobufFieldName:  "on_chain_at",
			ProtobufType:       "uint64",
			ProtobufPos:        14,
		},

		{
			Index:              14,
			Name:               "sealed_at",
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
			GoFieldName:        "SealedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "sealedAt",
			ProtobufFieldName:  "sealed_at",
			ProtobufType:       "uint64",
			ProtobufPos:        15,
		},

		{
			Index:              15,
			Name:               "last_message",
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
			GoFieldName:        "LastMessage",
			GoFieldType:        "null.String",
			JSONFieldName:      "lastMessage",
			ProtobufFieldName:  "last_message",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		{
			Index:              16,
			Name:               "deal_protocol_version",
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
			GoFieldName:        "DealProtocolVersion",
			GoFieldType:        "null.String",
			JSONFieldName:      "dealProtocolVersion",
			ProtobufFieldName:  "deal_protocol_version",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		{
			Index:              17,
			Name:               "miner_version",
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
			GoFieldName:        "MinerVersion",
			GoFieldType:        "null.String",
			JSONFieldName:      "minerVersion",
			ProtobufFieldName:  "miner_version",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		{
			Index:              18,
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
			ProtobufPos:        19,
		},

		{
			Index:              19,
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
			ProtobufPos:        20,
		},

		{
			Index:              20,
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
			ProtobufPos:        21,
		},

		{
			Index:              21,
			Name:               "system_content_deal_id",
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
			GoFieldName:        "SystemContentDealID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "systemContentDealId",
			ProtobufFieldName:  "system_content_deal_id",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		{
			Index:              22,
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
			ProtobufPos:        23,
		},

		{
			Index:              23,
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
			ProtobufPos:        24,
		},

		{
			Index:              24,
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
			ProtobufPos:        25,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ContentDealLogs) TableName() string {
	return "content_deal_logs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ContentDealLogs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ContentDealLogs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ContentDealLogs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ContentDealLogs) TableInfo() *TableInfo {
	return content_deal_logsTableInfo
}
