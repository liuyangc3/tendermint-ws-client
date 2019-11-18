package main

import (
	"github.com/liuyangc3/tendermint-ws-client/cmd"
)

func main() {

	cmd.Execute()

	// 	var rpcAddr string
	// 	flag.StringVar(&rpcAddr, "rpc-addr", "tcp://0.0.0.0:26657", "Tendermint RPC Websocket address")
	// 	flag.Usage = func() {
	// 		fmt.Println(`Tendermint websock client subscribes events
	// Usage:
	// 	tm-wsclient [-rpc-addr="tcp://localhost:26657"]
	// Examples:
	// 	# monitor single instance
	// 	tm-wsclient tcp://localhost:26657`)
	// 		fmt.Println("Flags:")
	// 		flag.PrintDefaults()
	// 	}

	// 	flag.Parse()

	// 	if flag.NArg() == 0 {
	// 		flag.Usage()
	// 		os.Exit(1)
	// 	}

	// 	client := client.NewHTTP(rpcAddr, "/websocket")
	// 	err := client.Start()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer client.Stop()

	// 	status, err := client.Status()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// see below for all event types
	// 	// https://godoc.org/github.com/tendermint/tendermint/types#pkg-constants
	// 	ctx := context.Background()
	// 	subscriber := "tendermint-websocket-client"
	// 	query := "tm.event = 'Vote'"
	// 	txs, err := client.Subscribe(ctx, subscriber, query)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer client.Unsubscribe(ctx, subscriber, query)

	// 	// EventDataVote
	// 	// type Vote struct {
	// 	// 	Type             SignedMsgType `json:"type"`
	// 	// 	Height           int64         `json:"height"`
	// 	// 	Round            int           `json:"round"`
	// 	// 	BlockID          BlockID       `json:"block_id"` // zero if vote is nil.
	// 	// 	Timestamp        time.Time     `json:"timestamp"`
	// 	// 	ValidatorAddress Address       `json:"validator_address"`
	// 	// 	ValidatorIndex   int           `json:"validator_index"`
	// 	// 	Signature        []byte        `json:"signature"`
	// 	// }

	// 	// Vote.Type SignedMsgType
	// 	// PrevoteType   SignedMsgType = 0x01
	// 	// PrecommitType SignedMsgType = 0x02

	// 	var lastHeight int64 = 0
	// 	var prevote int
	// 	var precommit int
	// 	var lastTime time.Time
	// 	var validatorMissPrevote bool = true
	// 	var validatorMissPrecommit bool = true

	// 	for e := range txs {
	// 		vote := e.Data.(types.EventDataVote).Vote
	// 		if vote.Height > lastHeight {
	// 			delta := vote.Timestamp.Sub(lastTime)

	// 			send_data("tendermint")

	// 			lastHeight = vote.Height
	// 			prevote = 0
	// 			precommit = 0

	// 		}
	// 		lastTime = vote.Timestamp

	// 		switch vote.Type {
	// 		case types.PrevoteType:
	// 			if bytes.Compare(status.ValidatorInfo.Address, vote.ValidatorAddress) == 0 {
	// 				validatorMissPrevote = false
	// 			}
	// 			prevote++
	// 		case types.PrecommitType:
	// 			if bytes.Compare(status.ValidatorInfo.Address, vote.ValidatorAddress) == 0 {
	// 				validatorMissPrecommit = false
	// 			}
	// 			precommit++
	// 		}
	// 	}
}
