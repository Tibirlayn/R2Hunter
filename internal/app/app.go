package app

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/Tibirlayn/R2Hunter/internal/app/restapi"
	"github.com/Tibirlayn/R2Hunter/internal/config"
	"github.com/Tibirlayn/R2Hunter/internal/service/account/auth"
	"github.com/Tibirlayn/R2Hunter/internal/service/account/member"
	"github.com/Tibirlayn/R2Hunter/internal/service/game/pc"
	"github.com/Tibirlayn/R2Hunter/storage/mssql"
)

type App struct {
	RestApi *restapi.App
}

func New(log *slog.Logger, address string, cfgdb *config.ConfigDB, tokenTLL time.Duration) *App {

	// инициализировать СУБД: MS SQL
	accStorage, err := mssql.NewAccountStorage(cfgdb)
	if err != nil {
		panic(err)
	}

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

	authService := auth.New(log, accStorage, accStorage, accStorage, tokenTLL)
	gamService := pc.New(log, gamStorage, tokenTLL)
	memberService := member.New(log, accStorage, authService, gamService, tokenTLL)
	restapi := restapi.New(log, authService, memberService, address)

	return &App{RestApi: restapi}
}

