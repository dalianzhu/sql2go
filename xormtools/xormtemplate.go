package xormtools

import (
	"bytes"
	"fmt"
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

func GenerateStructMethod(stmt *sqlparser.DDL, dbEngine string) (string, error) {
	d := new(CodeInfo)
	var err error
	d.StrucName, d.PK, d.PKType, d.PKTypeNil, d.Fields, err = GenerateXormStruct(stmt)
	if err != nil {
		return "", err
	}
	d.DbEngine = dbEngine
	tmpl := template.New("t1")
	// tmpl = tmpl.Funcs(MapFuncs)
	tmpl = template.Must(tmpl.Parse(DefaultTemplate))

	var doc bytes.Buffer
	tmpl.Execute(&doc, d)
	return string(doc.Bytes()), nil
}

func GenerateXormStruct(stmt *sqlparser.DDL) (structName string,
	pk string, pkType string, pkTypeNil string, fields []*Field, err error) {
	tableName := stmt.NewName.Name.String()
	structName = snakeCaseToCamel(tableName)
	primary := getPrimarys(stmt.TableSpec)

	pk = snakeCaseToCamel(primary[0])

	for _, col := range stmt.TableSpec.Columns {
		columnType := col.Type.Type
		goType, ok := sqlType2GoMap[columnType]
		if !ok {
			err = fmt.Errorf("sql %v type is not supported", columnType)
			return
		}

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
