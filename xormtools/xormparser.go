package xormtools

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/xwb1989/sqlparser"
)

type XormTag struct {
	pk         string
	autoincr   string
	notNull    string
	defaultVal string
	comment    string
	ttype      string
	field      string
}

func (x *XormTag) String() string {
	return fmt.Sprintf("`json:\"%v\" xorm:\"%v%v%v%v%v%v '%v'\"`",
		x.field, x.pk, x.autoincr, x.notNull, x.defaultVal, x.comment, x.ttype, x.field)
}

func defaultVal(s string) string {
	intVal, err := strconv.Atoi(s)
	if err == nil {
		return fmt.Sprintf("default %v ", intVal)
	}
	if s == "current_timestamp" {
		return fmt.Sprintf("default 'CURRENT_TIMESTAMP' ")
	}

	if s == "" {
		return "default '' "
	}
	return fmt.Sprintf("default '%s' ", s)
}

func getXormTag(col *sqlparser.ColumnDefinition) *XormTag {
	// `json:"id" xorm:"pk autoincr BIGINT(11) 'id'"`
	// `json:"subid" xorm:"not null default 0 BIGINT(11) 'subId'`
	ret := new(XormTag)
	ret.field = col.Name.String()
	ret.pk = ""
	if col.Type.Autoincrement {
		ret.autoincr = "autoincr "
	}
	if col.Type.NotNull {
		ret.notNull = "not null "
	}
	if col.Type.Default != nil {
		ret.defaultVal = defaultVal(string(col.Type.Default.Val))
	}

	ttype := strings.ToUpper(col.Type.Type)
	log.Printf("field:%v %v", ret.field, col.Type.Length)
	if col.Type.Length != nil {
		ttype = fmt.Sprintf("%s(%s)", ttype, col.Type.Length.Val)
	}
	ret.ttype = ttype

	if col.Type.Comment != nil {
		ret.comment = fmt.Sprintf("comment('%s') ", col.Type.Comment.Val)
	}
	return ret
}

type XormConverter struct {
}

func (x *XormConverter) StmtToGo(stmt *sqlparser.DDL, pkgName string, dbEngine string) (string, error) {
	builder := strings.Builder{}
	tableName := stmt.NewName.Name.String()

	header := fmt.Sprintf("package %s\n", pkgName)
	// import time package
	headerPkg := "import (\n" +
		"\t\"time\"\n" +
		")\n\n"
	importTime := false

	structName := snakeCaseToCamel(tableName)
	structStart := fmt.Sprintf("type %s struct { \n", structName)
	builder.WriteString(structStart)

	primary := getPrimary(stmt.TableSpec)

	codeInfo := new(CodeInfo)
	if len(primary) == 0 {
		return "", fmt.Errorf("no primary key")
	}
	codeInfo.PK = primary[0]
	codeInfo.DbEngine = dbEngine

	for _, col := range stmt.TableSpec.Columns {
		columnType := col.Type.Type

		if col.Type.Unsigned {
			columnType += " unsigned"
		}

		goType := sqlTypeMap[columnType]
		if goType == "time.Time" {
			importTime = true
		}

		builder.WriteString("\t")
		builder.WriteString(snakeCaseToCamel(col.Name.String()))
		builder.WriteString("\t")
		builder.WriteString(goType)
		builder.WriteString("\t")
		xormTag := getXormTag(col)
		if stringIn(col.Name.String(), primary) {
			xormTag.pk = "pk "
		}
		builder.WriteString(xormTag.String() + "\n")
	}
	builder.WriteString("}\n")

	// struct info
	structInfo := GenerateStructInfo(stmt, dbEngine)

	if importTime {
		return header + headerPkg + builder.String() + structInfo, nil
	}
	return header + builder.String() + structInfo, nil
}
func snakeCaseToCamel(str string) string {
	builder := strings.Builder{}
	index := 0
	if str[0] >= 'a' && str[0] <= 'z' {
		builder.WriteByte(str[0] - ('a' - 'A'))
		index = 1
	}
	for i := index; i < len(str); i++ {
		if str[i] == '_' && i+1 < len(str) {
			if str[i+1] >= 'a' && str[i+1] <= 'z' {
				builder.WriteByte(str[i+1] - ('a' - 'A'))
				i++
				continue
			}
		}
		builder.WriteByte(str[i])
	}
	return builder.String()
}

func getPrimary(table *sqlparser.TableSpec) []string {
	ret := make([]string, 0)
	for _, index := range table.Indexes {
		if index.Info.Primary {
			for _, col := range index.Columns {
				ret = append(ret, col.Column.String())
			}
		}
	}
	return ret
}

func stringIn(s string, arr []string) bool {
	for _, item := range arr {
		if item == s {
			return true
		}
	}
	return false
}

var sqlTypeMap = map[string]string{
	"int":                "int",
	"integer":            "int",
	"tinyint":            "int8",
	"smallint":           "int16",
	"mediumint":          "int32",
	"bigint":             "int",
	"int unsigned":       "uint",
	"integer unsigned":   "uint",
	"tinyint unsigned":   "uint8",
	"smallint unsigned":  "uint16",
	"mediumint unsigned": "uint32",
	"bigint unsigned":    "uint64",
	"bit":                "byte",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time",
	"datetime":           "time.Time",
	"timestamp":          "time.Time",
	"time":               "time.Time",
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}
