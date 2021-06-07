package xormtools

var DefaultTemplate = `
func (t *{{.StrucName}}) Get() (*{{.StrucName}}, error) {
    ok, err := {{.DbEngine}}.Get(t)
    if err!=nil{
        err = fmt.Errorf("{{.StrucName}} get Error:%v", err)
        log.Error(err)
        return nil, err
    }
    if ok {
        return t, nil
    }
    return nil, errors.New("TSnmpSubsvr not exists")
}

func (t *{{.StrucName}}) GetList() ([]*{{.StrucName}}, error) {
    var CodeInfo []*{{.StrucName}}
    var err error
    err = {{.DbEngine}}.Find(&CodeInfo, t)
    if err != nil{
        err = fmt.Errorf("{{.StrucName}} get list Error:%v", err)
        log.Error(err)
        return nil, err
    }
    return CodeInfo, err 
}

func (t *{{.StrucName}}) Delete() (err error) {
    affected, err := {{.DbEngine}}.Delete(t)
    if err!=nil{
        err = fmt.Errorf("{{.StrucName}} delete Error:%v", err)
        log.Error(err)
        return err
    }
    log.Infof("{{.StrucName}} delete id:%v, affect:%v", t.{{.PK}}, affected)
    return nil
}

func (t *{{.StrucName}}) DeleteWithSession(session *xorm.Session) (err error) {
    affected, err := session.Delete(t)
    if err!=nil{
        err = fmt.Errorf("{{.StrucName}} delete Error:%v", err)
        log.Error(err)
        return err
    }
    log.Infof("{{.StrucName}} delete id:%v, affect:%v", t.{{.PK}}, affected)
    return nil
}


func (t *{{.StrucName}}) Insert() (err error) {
    _, err = {{.DbEngine}}.Insert(t)
    if err != nil{
        err = fmt.Errorf("{{.StrucName}} insert Error:%v", err)
        log.Error(err)
        return err
    }
    return nil
}

func (t *{{.StrucName}}) InsertWithSession(session *xorm.Session) (err error) {
    _, err = session.Insert(t)
    if err != nil{
        err = fmt.Errorf("{{.StrucName}} insert Error:%v", err)
        log.Error(err)
        return err
    }
    return nil
}

func (t *{{.StrucName}}) Modify() (err error) {
    if t.{{.PK}} == {{.PKTypeNil}} {
        err = errors.New("{{.StrucName}} modify Error, pk is nil")
        log.Error(err)
        return err
    }
    _, err = {{.DbEngine}}.ID(t.{{.PK}}).Update(t)
    if err != nil {
        err = fmt.Errorf("{{.StrucName}} modify Error:%v", err)
        log.Error(err)
        return err
    }
    return nil
}

func (t *{{.StrucName}}) ModifyWithSession(session *xorm.Session) (err error) {
    if t.{{.PK}} == {{.PKTypeNil}} {
        err = errors.New("{{.StrucName}} modify with session Error, pk is nil")
        log.Error(err)
        return err
    }
    _, err = session.ID(t.{{.PK}}).Update(t)
    if err != nil {
        err = fmt.Errorf("{{.StrucName}} modify with session Error:%v", err)
        log.Error(err)
        return err
    }
    return nil
}

{{range .Fields}}

{{ if not (eq .Name $.PK) }}
func (t *{{$.StrucName}}) Set{{.Name}}(val {{.Type}}) (err error) {
    if t.{{$.PK}} == {{$.PKTypeNil}} {
        err = errors.New("{{$.StrucName}} set {{.Name}} Error, pk is nil")
        log.Error(err)
        return err
    }

    tp:=new({{$.StrucName}})
    tp.{{.Name}} = val
    _, err = {{$.DbEngine}}.ID(t.{{$.PK}}).Cols("{{.DbName}}").Update(tp)
    if err != nil{
        err = fmt.Errorf("{{$.StrucName}} set {{.Name}} Error:%v", err)
        log.Error(err)
        return err
    }
    return nil 
}
{{end}}

{{end}}
`
