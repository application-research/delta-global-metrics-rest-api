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


Table: delta_node_geo_locations
[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
[ 1] ip                                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] country                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] city                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] region                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] zip                                            TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] lat                                            NUMERIC              null: true   primary: false  isArray: false  auto: false  col: NUMERIC         len: -1      default: []
[ 7] lon                                            NUMERIC              null: true   primary: false  isArray: false  auto: false  col: NUMERIC         len: -1      default: []
[ 8] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[ 9] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 89,    "ip": "uwpBtXBntBWgkCKraXSpkctDe",    "country": "GHPtnoTwpFebPPVvcspCrMqZE",    "city": "kNBasfOCYAWmojkCMuiBaimiw",    "region": "vbJjsoGjvIDUSYIYsbvtGKygD",    "zip": "VRnvGNQBsiJRbbUgODZEJDkcB",    "lat": 0.2523771892476505,    "lon": 0.8991950175159545,    "createdAt": "2205-06-16T18:43:47.606371433-04:00",    "updatedAt": "2166-09-03T02:07:18.062428741-04:00"}



*/

// DeltaNodeGeoLocations struct is a row record of the delta_node_geo_locations table in the estuary database
type DeltaNodeGeoLocations struct {
	//[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:INT8;"`
	//[ 1] ip                                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	IP null.String `gorm:"column:ip;type:TEXT;"`
	//[ 2] country                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Country null.String `gorm:"column:country;type:TEXT;"`
	//[ 3] city                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	City null.String `gorm:"column:city;type:TEXT;"`
	//[ 4] region                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Region null.String `gorm:"column:region;type:TEXT;"`
	//[ 5] zip                                            TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Zip null.String `gorm:"column:zip;type:TEXT;"`
	//[ 6] lat                                            NUMERIC              null: true   primary: false  isArray: false  auto: false  col: NUMERIC         len: -1      default: []
	Lat null.Float `gorm:"column:lat;type:NUMERIC;"`
	//[ 7] lon                                            NUMERIC              null: true   primary: false  isArray: false  auto: false  col: NUMERIC         len: -1      default: []
	Lon null.Float `gorm:"column:lon;type:NUMERIC;"`
	//[ 8] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:TIMESTAMPTZ;"`
	//[ 9] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;"`
}

var delta_node_geo_locationsTableInfo = &TableInfo{
	Name: "delta_node_geo_locations",
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
			Name:               "ip",
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
			GoFieldName:        "IP",
			GoFieldType:        "null.String",
			JSONFieldName:      "ip",
			ProtobufFieldName:  "ip",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "country",
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
			GoFieldName:        "Country",
			GoFieldType:        "null.String",
			JSONFieldName:      "country",
			ProtobufFieldName:  "country",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "city",
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
			GoFieldName:        "City",
			GoFieldType:        "null.String",
			JSONFieldName:      "city",
			ProtobufFieldName:  "city",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "region",
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
			GoFieldName:        "Region",
			GoFieldType:        "null.String",
			JSONFieldName:      "region",
			ProtobufFieldName:  "region",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "zip",
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
			GoFieldName:        "Zip",
			GoFieldType:        "null.String",
			JSONFieldName:      "zip",
			ProtobufFieldName:  "zip",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "lat",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "NUMERIC",
			DatabaseTypePretty: "NUMERIC",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "NUMERIC",
			ColumnLength:       -1,
			GoFieldName:        "Lat",
			GoFieldType:        "null.Float",
			JSONFieldName:      "lat",
			ProtobufFieldName:  "lat",
			ProtobufType:       "float",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "lon",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "NUMERIC",
			DatabaseTypePretty: "NUMERIC",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "NUMERIC",
			ColumnLength:       -1,
			GoFieldName:        "Lon",
			GoFieldType:        "null.Float",
			JSONFieldName:      "lon",
			ProtobufFieldName:  "lon",
			ProtobufType:       "float",
			ProtobufPos:        8,
		},

		{
			Index:              8,
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
			ProtobufPos:        9,
		},

		{
			Index:              9,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *DeltaNodeGeoLocations) TableName() string {
	return "delta_node_geo_locations"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DeltaNodeGeoLocations) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DeltaNodeGeoLocations) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DeltaNodeGeoLocations) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *DeltaNodeGeoLocations) TableInfo() *TableInfo {
	return delta_node_geo_locationsTableInfo
}
