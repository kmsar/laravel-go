package drivers

//import (
//	"errors"
//	"fmt"
//	_ "github.com/ClickHouse/clickhouse-go/v2"
//	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/supports-master/logs"
//
//	//"github.com/goal-web/contracts"
//	//"github.com/goal-web/supports/logs"
//	//"github.com/goal-web/supports/utils"
//	"github.com/jmoiron/sqlx"
//	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
//	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEvent"
//	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
//	"github.com/laravel-go-version/v2/pkg/Illuminate/Database/support"
//	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"
//
//	"strings"
//)
//
//type Clickhouse struct {
//	support.Executor
//}
//
//func (c Clickhouse) Begin() (IDatabase.DBTx, error) {
//	return nil, errors.New("begin is not supported")
//}
//
//func (c Clickhouse) Transaction(f func(executor IDatabase.SqlExecutor) error) error {
//	return errors.New("transaction is not supported")
//}
//
//
//
//func ClickHouseConnector(config Support.Fields, events IEvent.EventDispatcher) IDatabase.DBConnection {
//	var dsn = Field.GetStringField(config, "dsn")
//	if dsn == "" {
//		address := config["address"].([]string)
//		dsn = fmt.Sprintf("tcp://%s?debug=%s",
//			address[0],
//			//config["database"].(string),
//			config["debug"].(string),
//		)
//		if config["username"] != "" {
//			dsn = fmt.Sprintf("%s&username=%s",
//				dsn,
//				config["username"].(string),
//			)
//		}
//		if config["database"] != "" {
//			dsn = fmt.Sprintf("%s&database=%s",
//				dsn,
//				config["database"].(string),
//			)
//		}
//		if config["password"] != "" {
//			dsn = fmt.Sprintf("%s&password=%s",
//				dsn,
//				config["password"].(string),
//			)
//		}
//		if len(address) > 1 {
//			dsn = fmt.Sprintf("%s&alt_hosts=%s", dsn, strings.Join(address[1:], ","))
//		}
//	}
//	db, err := sqlx.Connect("clickhouse", dsn)
//	if err != nil {
//		logs.WithError(err).WithField("dsn", dsn).WithField("config", config).Fatal("clickhouse 数据库初始化失败")
//	}
//	db.SetMaxOpenConns(Field.GetIntField(config, "max_connections"))
//	db.SetMaxIdleConns(Field.GetIntField(config, "max_idles"))
//	return &Clickhouse{
//		support.NewExecutor(db, events, dollarNParamBindWrapper),
//	}
//}
