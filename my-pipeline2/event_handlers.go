package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/Zettablock/zsource/dao/ethereum"
	"github.com/Zettablock/zsource/utils"
)

func HandlePrintLog(log ethereum.Log, deps *utils.Deps) (bool, error) {
	deps.Logger.Info("print log", "blockNumber", log.BlockNumber)
	slog.Info("print log", "blockNumber", log.BlockNumber)
	deps.Logger.Info("print log", "logIndex", log.LogIndex)
	slog.Info("print log", "logIndex", log.LogIndex)

	fmt.Println("new new new new new new new new new")

	time.Sleep(5 * time.Second)

	return true, nil
}
