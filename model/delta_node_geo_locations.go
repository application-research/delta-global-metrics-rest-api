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
{    "id": 29,    "ip": "AoFZjasOXLsPPxMGoCMNahbty",    "country": "eAdelWAlBUZegkeKfJGvqvGuF",    "city": "bELhkxHYWYaUkwlwokyEMCSiW",    "region": "LwtRiOAMdQARxrHrblpBlBUPV",    "zip": "ybMAKDYjSrGUXrUHAsykuqqAs",    "lat": 0.5675549651243351,    "lon": 0.3986043435344051,    "created_at": "2213-10-23T16:18:24.751263651-04:00",    "updated_at": "2026-02-19T10:08:50.904405976-05:00"}



*/

// DeltaNodeGeoLocations struct is a row record of the delta_node_geo_locations table in the estuary database
type DeltaNodeGeoLocations struct {
	//[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:INT8;"`
	//[ 1] ip                                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	IP sql.NullString `gorm:"column:ip;type:TEXT;"`
	//[ 2] country                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Country sql.NullString `gorm:"column:country;type:TEXT;"`
	//[ 3] city                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	City sql.NullString `gorm:"column:city;type:TEXT;"`
	//[ 4] region                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Region sql.NullString `gorm:"column:region;type:TEXT;"`
	//[ 5] zip                                            TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Zip sql.NullString `gorm:"column:zip;type:TEXT;"`
	//[ 6] lat                                            NUMERIC              null: true   primary: false  isArray: false  auto: false  col: NUMERIC         len: -1      default: []
	Lat sql.NullFloat64 `gorm:"column:lat;type:NUMERIC;"`
	//[ 7] lon                                            NUMERIC              null: true   primary: false  isArray: false  auto: false  col: NUMERIC         len: -1      default: []
	Lon sql.NullFloat64 `gorm:"column:lon;type:NUMERIC;"`
	//[ 8] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMPTZ;"`
	//[ 9] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;"`
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
			GoFieldType:        "sql.NullString",
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
			GoFieldType:        "sql.NullString",
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
			GoFieldType:        "sql.NullString",
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
			GoFieldType:        "sql.NullString",
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
			GoFieldType:        "sql.NullString",
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
			GoFieldType:        "sql.NullFloat64",
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
			GoFieldType:        "sql.NullFloat64",
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
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
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
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
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
