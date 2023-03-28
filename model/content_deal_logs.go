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
{    "id": 78,    "content": 50,    "prop_cid": "wPitHmApMJNqFkbCVCEXLPOgO",    "deal_uuid": "KfpfjKjQsUEdUlCKdpDmZPHSm",    "miner": "oNPSjFMgtvtqevtwSaCronoHs",    "deal_id": 69,    "failed": false,    "verified": false,    "slashed": false,    "failed_at": "2073-03-13T07:58:47.787105952-04:00",    "dt_chan": "gbajoJITYJHctKypowgyQwouy",    "transfer_started": "2288-03-25T04:38:08.461808637-04:00",    "transfer_finished": "2054-06-03T09:38:09.659274379-04:00",    "on_chain_at": "2242-09-05T02:11:35.872053883-04:00",    "sealed_at": "2029-08-07T07:31:54.707466325-04:00",    "last_message": "fJmUPhsRNdLSGkcfdTaGGgnAX",    "deal_protocol_version": "etTRJJaJfuZJNqcywDOanNvWe",    "miner_version": "XjJCulvZcXNXKflGaedJVMgFF",    "node_info": "onYUJQpmPnpMgOHnWwjaWDeUd",    "requester_info": "NGGHKoXWAibSANLnomZsFWOQB",    "requesting_api_key": "njTDXKscmeIOOHGiipNNRcYyH",    "system_content_deal_id": 1,    "created_at": "2226-07-15T00:10:30.581336342-04:00",    "updated_at": "2098-07-14T10:42:36.953905498-04:00",    "delta_node_uuid": "HbTqKrUqmSCnaYUYmRoPtipKX"}



*/

// ContentDealLogs struct is a row record of the content_deal_logs table in the estuary database
type ContentDealLogs struct {
	//[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:INT8;"`
	//[ 1] content                                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	Content sql.NullInt64 `gorm:"column:content;type:INT8;"`
	//[ 2] prop_cid                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	PropCid sql.NullString `gorm:"column:prop_cid;type:TEXT;"`
	//[ 3] deal_uuid                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DealUUID sql.NullString `gorm:"column:deal_uuid;type:TEXT;"`
	//[ 4] miner                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Miner sql.NullString `gorm:"column:miner;type:TEXT;"`
	//[ 5] deal_id                                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	DealID sql.NullInt64 `gorm:"column:deal_id;type:INT8;"`
	//[ 6] failed                                         BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	Failed sql.NullBool `gorm:"column:failed;type:BOOL;"`
	//[ 7] verified                                       BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	Verified sql.NullBool `gorm:"column:verified;type:BOOL;"`
	//[ 8] slashed                                        BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	Slashed sql.NullBool `gorm:"column:slashed;type:BOOL;"`
	//[ 9] failed_at                                      TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	FailedAt time.Time `gorm:"column:failed_at;type:TIMESTAMPTZ;"`
	//[10] dt_chan                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DtChan sql.NullString `gorm:"column:dt_chan;type:TEXT;"`
	//[11] transfer_started                               TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	TransferStarted time.Time `gorm:"column:transfer_started;type:TIMESTAMPTZ;"`
	//[12] transfer_finished                              TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	TransferFinished time.Time `gorm:"column:transfer_finished;type:TIMESTAMPTZ;"`
	//[13] on_chain_at                                    TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	OnChainAt time.Time `gorm:"column:on_chain_at;type:TIMESTAMPTZ;"`
	//[14] sealed_at                                      TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	SealedAt time.Time `gorm:"column:sealed_at;type:TIMESTAMPTZ;"`
	//[15] last_message                                   TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	LastMessage sql.NullString `gorm:"column:last_message;type:TEXT;"`
	//[16] deal_protocol_version                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DealProtocolVersion sql.NullString `gorm:"column:deal_protocol_version;type:TEXT;"`
	//[17] miner_version                                  TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	MinerVersion sql.NullString `gorm:"column:miner_version;type:TEXT;"`
	//[18] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	NodeInfo sql.NullString `gorm:"column:node_info;type:TEXT;"`
	//[19] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequesterInfo sql.NullString `gorm:"column:requester_info;type:TEXT;"`
	//[20] requesting_api_key                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequestingAPIKey sql.NullString `gorm:"column:requesting_api_key;type:TEXT;"`
	//[21] system_content_deal_id                         INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	SystemContentDealID sql.NullInt64 `gorm:"column:system_content_deal_id;type:INT8;"`
	//[22] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMPTZ;"`
	//[23] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;"`
	//[24] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DeltaNodeUUID sql.NullString `gorm:"column:delta_node_uuid;type:TEXT;"`
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
			GoFieldType:        "sql.NullInt64",
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
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "prop_cid",
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
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "deal_uuid",
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
			GoFieldType:        "sql.NullString",
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
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "deal_id",
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
			GoFieldType:        "sql.NullBool",
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
			GoFieldType:        "sql.NullBool",
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
			GoFieldType:        "sql.NullBool",
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
			GoFieldType:        "time.Time",
			JSONFieldName:      "failed_at",
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
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dt_chan",
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
			GoFieldType:        "time.Time",
			JSONFieldName:      "transfer_started",
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
			GoFieldType:        "time.Time",
			JSONFieldName:      "transfer_finished",
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
			GoFieldType:        "time.Time",
			JSONFieldName:      "on_chain_at",
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
			GoFieldType:        "time.Time",
			JSONFieldName:      "sealed_at",
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
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "last_message",
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
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "deal_protocol_version",
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
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "miner_version",
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
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "node_info",
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
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "requester_info",
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
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "requesting_api_key",
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
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "system_content_deal_id",
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
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
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
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
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
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "delta_node_uuid",
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
