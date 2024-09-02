package main

import (
	"log/slog"
	"time"

	"example-project/dao"
	"github.com/Zettablock/zsource/dao/ethereum"
	"github.com/Zettablock/zsource/utils"
)

func HandlePrintBlock(blockNumber ethereum.Block, deps *utils.Deps) (bool, error) {
	deps.Logger.Info("print block", "blockNumber", blockNumber.Number)
	slog.Info("print block", "blockNumber", blockNumber.Number)

	block := dao.BlockRPC{
		Number:     blockNumber.Number,
		Hash:       blockNumber.Hash,
		ParentHash: blockNumber.ParentHash,
	}

	deps.DestinationDB.Save(&block)

	time.Sleep(5 * time.Second)

	return true, nil
}
