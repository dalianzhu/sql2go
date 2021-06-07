package xormtools

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/xwb1989/sqlparser"
)

type CodeInfo struct {
	StrucName string
	DbEngine  string
	PK        string
	PKType    string
	Fields    []*Field
	PKTypeNil string
}

type Field struct {
	Name   string
	Type   string
	DbName string
}

func GenerateStructInfo(stmt *sqlparser.DDL, dbEngine string) string {
	d := new(CodeInfo)
	d.StrucName, d.PK, d.PKType, d.PKTypeNil, d.Fields = GetInfo(stmt)
	d.DbEngine = dbEngine
	tmpl := template.New("t1")
	// tmpl = tmpl.Funcs(MapFuncs)
	tmpl = template.Must(tmpl.Parse(DefaultTemplate))

	var doc bytes.Buffer
	tmpl.Execute(&doc, d)
	return string(doc.Bytes())
}

func GetInfo(stmt *sqlparser.DDL) (structName string, pk string, pkType string, pkTypeNil string, fields []*Field) {
	tableName := stmt.NewName.Name.String()
	structName = snakeCaseToCamel(tableName)
	primary := getPrimary(stmt.TableSpec)

	pk = snakeCaseToCamel(primary[0])

	for _, col := range stmt.TableSpec.Columns {
		columnType := col.Type.Type
		goType := sqlTypeMap[columnType]

		if col.Type.Unsigned {
			columnType += " unsigned"
		}
		colName := snakeCaseToCamel(col.Name.String())
		if colName == pk {
			pkType = goType
			pkTypeNil = getNilVal(pkType)
		} else {
			fields = append(fields, &Field{
				Name:   colName,
				Type:   goType,
				DbName: col.Name.String(),
			})
		}
	}
	return
}

func getNilVal(goType string) string {
	if strings.HasPrefix(goType, "int") {
		return "0"
	}
	if goType == "time.Time" {
		return "nil"
	}
	return `""`
}
