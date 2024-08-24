package app

import (
	"fmt"
	"log/slog"

	"github.com/Tibirlayn/R2Hunter/internal/app/restapi"
	"github.com/Tibirlayn/R2Hunter/internal/config"
	"github.com/Tibirlayn/R2Hunter/storage/mssql"
)

type App struct {
	RestApi *restapi.App
}

func New(log *slog.Logger, address string, cfgdb *config.ConfigDB) *App {

	// инициализировать СУБД: MS SQL
	storage, err := mssql.New(cfgdb)
	if err != nil {
		panic(err)
	}
	fmt.Println(storage)

	accStorage, err := mssql.NewAccountStorage(cfgdb)
	if err != nil {
		panic(err)
	}
	fmt.Println(accStorage)

	batStorage, err := mssql.NewBattleStorage(cfgdb)
	if err != nil {
		panic(err)
	}
	fmt.Println(batStorage)

	bilStorage, err := mssql.NewBillingStorage(cfgdb)
	if err != nil {
		panic(err)
	}
	fmt.Println(bilStorage)

	gamStorage, err := mssql.NewGameStorage(cfgdb)
	if err != nil {
		panic(err)
	}
	fmt.Println(gamStorage)
	
	logStorage, err := mssql.NewLogsStorage(cfgdb)
	if err != nil {
		panic(err)
	}
	fmt.Println(logStorage)

	parStorage, err := mssql.NewParmStorage(cfgdb)
	if err != nil {
		panic(err)
	}
	fmt.Println(parStorage)

	statStorage, err := mssql.NewStatisticsStorage(cfgdb)
	if err != nil {
		panic(err)
	}
	fmt.Println(statStorage)

	

	restapi := restapi.New(log, address)

	return &App{RestApi: restapi}
}

