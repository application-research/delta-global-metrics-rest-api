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


Table: piece_commitment_logs
[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
[ 1] cid                                            TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] piece                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] size                                           INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 4] padded_piece_size                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 5] un_padded_piece_size                           INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 6] status                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 7] last_message                                   TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 8] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 9] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[10] requesting_api_key                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[11] system_content_piece_commitment_id             INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[12] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[13] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[14] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 96,    "cid": "lSpleRdPmtqSgPIPaupwBTuyl",    "piece": "rdrggiHeNdpKHBfFylyCggASX",    "size": 54,    "paddedPieceSize": 34,    "unPaddedPieceSize": 67,    "status": "rZcNQmBiCIqntDBaEJOEBDeUU",    "lastMessage": "nMdhoVjfwPSQpWiVwprTHtQDv",    "nodeInfo": "hgMKEJrQUrKFCsxmSNhmKrfns",    "requesterInfo": "gjlGNNmhGjgIZlGkkRLInnaIl",    "requestingApiKey": "OdFJvYFetTPCvOnNtqEEefDtu",    "systemContentPieceCommitmentId": 41,    "createdAt": "2263-03-20T02:46:34.0114278-04:00",    "updatedAt": "2233-07-09T11:09:29.842592668-04:00",    "deltaNodeUuid": "BHSEdeJWRcPaiTAwASdFCrYor"}



*/

// PieceCommitmentLogs struct is a row record of the piece_commitment_logs table in the estuary database
type PieceCommitmentLogs struct {
	//[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:INT8;"`
	//[ 1] cid                                            TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Cid null.String `gorm:"column:cid;type:TEXT;"`
	//[ 2] piece                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Piece null.String `gorm:"column:piece;type:TEXT;"`
	//[ 3] size                                           INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	Size null.Int `gorm:"column:size;type:INT8;"`
	//[ 4] padded_piece_size                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	PaddedPieceSize null.Int `gorm:"column:padded_piece_size;type:INT8;"`
	//[ 5] un_padded_piece_size                           INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	UnPaddedPieceSize null.Int `gorm:"column:un_padded_piece_size;type:INT8;"`
	//[ 6] status                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Status null.String `gorm:"column:status;type:TEXT;"`
	//[ 7] last_message                                   TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	LastMessage null.String `gorm:"column:last_message;type:TEXT;"`
	//[ 8] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	NodeInfo null.String `gorm:"column:node_info;type:TEXT;"`
	//[ 9] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequesterInfo null.String `gorm:"column:requester_info;type:TEXT;"`
	//[10] requesting_api_key                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequestingAPIKey null.String `gorm:"column:requesting_api_key;type:TEXT;"`
	//[11] system_content_piece_commitment_id             INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	SystemContentPieceCommitmentID null.Int `gorm:"column:system_content_piece_commitment_id;type:INT8;"`
	//[12] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:TIMESTAMPTZ;"`
	//[13] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;"`
	//[14] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DeltaNodeUUID null.String `gorm:"column:delta_node_uuid;type:TEXT;"`
}

var piece_commitment_logsTableInfo = &TableInfo{
	Name: "piece_commitment_logs",
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
			Name:               "cid",
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
			GoFieldName:        "Cid",
			GoFieldType:        "null.String",
			JSONFieldName:      "cid",
			ProtobufFieldName:  "cid",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "piece",
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
			GoFieldName:        "Piece",
			GoFieldType:        "null.String",
			JSONFieldName:      "piece",
			ProtobufFieldName:  "piece",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "size",
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
			GoFieldName:        "Size",
			GoFieldType:        "null.Int",
			JSONFieldName:      "size",
			ProtobufFieldName:  "size",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "padded_piece_size",
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
			GoFieldName:        "PaddedPieceSize",
			GoFieldType:        "null.Int",
			JSONFieldName:      "paddedPieceSize",
			ProtobufFieldName:  "padded_piece_size",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "un_padded_piece_size",
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
			GoFieldName:        "UnPaddedPieceSize",
			GoFieldType:        "null.Int",
			JSONFieldName:      "unPaddedPieceSize",
			ProtobufFieldName:  "un_padded_piece_size",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "status",
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
			GoFieldName:        "Status",
			GoFieldType:        "null.String",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
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
			ProtobufPos:        8,
		},

		{
			Index:              8,
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
			ProtobufPos:        9,
		},

		{
			Index:              9,
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
			ProtobufPos:        10,
		},

		{
			Index:              10,
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
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "system_content_piece_commitment_id",
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
			GoFieldName:        "SystemContentPieceCommitmentID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "systemContentPieceCommitmentId",
			ProtobufFieldName:  "system_content_piece_commitment_id",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		{
			Index:              12,
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
			ProtobufPos:        13,
		},

		{
			Index:              13,
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
			ProtobufPos:        14,
		},

		{
			Index:              14,
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
			ProtobufPos:        15,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PieceCommitmentLogs) TableName() string {
	return "piece_commitment_logs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PieceCommitmentLogs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PieceCommitmentLogs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PieceCommitmentLogs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PieceCommitmentLogs) TableInfo() *TableInfo {
	return piece_commitment_logsTableInfo
}
