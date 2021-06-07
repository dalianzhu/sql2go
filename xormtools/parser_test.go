package xormtools

import (
	"log"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	sql := `CREATE TABLE t_test_table (
		id bigint(11) NOT NULL AUTO_INCREMENT,
		creator varchar(128) NOT NULL COMMENT '这是creator',
		secretId int NOT NULL default 1 COMMENT '这是secret',
		dataType varchar(32) NOT NULL default "" COMMENT '数据类型，snmp/syslog等',
		callbackTopic varchar(1024) NOT NULL default "" COMMENT '如果是推送类型，此值有',
		info TEXT,
		createTime datetime NOT NULL default CURRENT_TIMESTAMP COMMENT '创建时间',
		lastUpdateTimestamp datetime NOT NULL,
		PRIMARY KEY (id),
		UNIQUE s (secretId),
		KEY creator (creator)
	  ) ENGINE=InnoDB AUTO_INCREMENT=161 DEFAULT CHARSET=utf8 COMMENT='订阅表';`

	parser := NewParser(new(XormConverter))
	ret, err := parser.Parse(sql, "model", "globalEngine")
	log.Printf("ret:%v, err:%v", ret, err)
}

func TestParser_Run(t *testing.T) {
	sql := `CREATE TABLE t_hiro_dal_template (
		id bigint(11) NOT NULL AUTO_INCREMENT,
		templateName varchar(128) NOT NULL COMMENT '模版的名字，比如netstatus.gtpl',
		content text COMMENT '这是模版的内容，是一个go template',
		lastModify datetime COMMENT '这是最后的修改时间',
        primary key (id),
		KEY templateName (templateName)
	  ) ENGINE=InnoDB AUTO_INCREMENT=161 DEFAULT CHARSET=utf8 COMMENT='hiro模版表';`

	parser := NewParser(new(XormConverter))
	ret, err := parser.Parse(sql, "model", "globalEngine")
	log.Printf("ret:%v, err:%v", ret, err)
}
