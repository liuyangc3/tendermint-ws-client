package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/types"
)

var c *client.HTTP
var serverRPCAddr string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tm-wscli",
	Short: "Tendermint websocket client",
	Long:  ``,
	Run:   runClient,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&serverRPCAddr, "server", "s", "tcp://localhost:26657", "Tendermint RPC address.")
	rootCmd.MarkFlagRequired("server")
}

// VoteCounter count vote
type VoteCounter struct {
	count   int
	missing bool
}

func newVoteCounter() *VoteCounter {
	return &VoteCounter{0, true}
}

func (v *VoteCounter) reset() {
	v.count = 0
	v.missing = true
}

func runClient(cmd *cobra.Command, args []string) {
	c = client.NewHTTP(serverRPCAddr, "/websocket")
	err := c.Start()
	if err != nil {
		panic(err)
	}
	defer c.Stop()

	// Get validator status
	status, err := c.Status()
	if err != nil {
		panic(err)
	}

	// Subscribe events
	ctx, cancel := context.WithCancel(context.Background())
	events, err := c.Subscribe(ctx, "tendermint-ws-client", "tm.event = 'Vote'")
	if err != nil {
		panic(err)
	}
	defer cancel()

	// Check events
	var lastHeight int64 = 0
	prevote := newVoteCounter()
	precommit := newVoteCounter()

	for e := range events {
		vote := e.Data.(types.EventDataVote).Vote

		if lastHeight == 0 {
			lastHeight = vote.Height
		}

		if vote.Height > lastHeight {
			// new Block Height begain
			if !precommit.missing {
				fmt.Println("not missing precommit")
			}
			if !prevote.missing {
				fmt.Println("not missing prevote")
			}
			// reset all varibles
			prevote.reset()
			precommit.reset()
		}

		switch vote.Type {
		case types.PrevoteType:
			if bytes.Compare(status.ValidatorInfo.Address, vote.ValidatorAddress) == 0 {
				prevote.missing = false
			}
			prevote.count++
		case types.PrecommitType:
			if bytes.Compare(status.ValidatorInfo.Address, vote.ValidatorAddress) == 0 {
				precommit.missing = false
			}
			precommit.count++
		}
	}
}
