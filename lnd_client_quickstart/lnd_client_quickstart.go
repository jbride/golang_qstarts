package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/lightninglabs/lndclient"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/lnrpc/verrpc"
	"github.com/lightningnetwork/lnd/lntest/wait"
)

const (
	defaultServerTimeout  = 10 * time.Second
	defaultConnectTimeout = 15 * time.Second
	defaultStartupTimeout = 5 * time.Second
)

var (
	clientOptions []lndclient.BasicClientOption
	basicClient   lnrpc.LightningClient
	grpcServices  *lndclient.GrpcLndServices
	bCtx          context.Context

	host    string = "localhost"
	network string = "regtest"
	macFile string = "admin.macaroon"
	tlsPath string = ""
	macPath string
	macData []byte
)

func setUpLNDClients() error {
	clientOptions = append(clientOptions, lndclient.MacaroonData(
		hex.EncodeToString(macData),
	))
	clientOptions = append(
		clientOptions, lndclient.MacFilename(filepath.Base(macPath)),
	)

	// The main RPC listener of lnd might need some time to start, it could
	// be that we run into a connection refused a few times. We use the
	// basic client connection to find out if the RPC server is started yet
	// because that doesn't do anything else than just connect. We'll check
	// if lnd is also ready to be used in the next step.
	fmt.Printf("setUpLNDClients() Connecting basic lnd client\n")
	err := wait.NoError(func() error {
		// Create an lnd client now that we have the full configuration.
		// We'll need a basic client and a full client because not all
		// subservers have the same requirements.
		var err error
		basicClient, err = lndclient.NewBasicClient(
			host, tlsPath, filepath.Dir(macPath), string(network),
			clientOptions...,
		)
		return err
	}, defaultStartupTimeout)
	if err != nil {
		return fmt.Errorf("could not create basic LND Client: %v", err)
	}

	lndInfo, err := basicClient.GetInfo(bCtx, &lnrpc.GetInfoRequest{})
	if err != nil {
		if !lndclient.IsUnlockError(err) {
			return fmt.Errorf("error querying remote "+"node : %v", err)
		}

	} else {
		fmt.Printf("Node version = %s  and alias = %s\n", lndInfo.GetVersion(), lndInfo.GetAlias())
	}

	channelGraph, err := basicClient.DescribeGraph(bCtx, &lnrpc.ChannelGraphRequest{})
	if err != nil {
		return fmt.Errorf("ERROR: DescribeGraph() %v", err)
	}

	prettyJsonNodes, err := json.MarshalIndent(channelGraph.Nodes, "", "  ")
	if err != nil {
		return fmt.Errorf("ERROR: pretty print graph json: %v", err)
	}
	fmt.Printf("channel edges: %s\t nodes: \n%s\n", channelGraph.Edges, prettyJsonNodes)

	grpcServices, err = lndclient.NewLndServices(
		&lndclient.LndServicesConfig{
			LndAddress:         host,
			Network:            lndclient.Network(network),
			CustomMacaroonPath: macPath,
			RPCTimeout:         defaultConnectTimeout,
			TLSPath:            tlsPath,
			CheckVersion: &verrpc.Version{
				AppMajor: 0,
				AppMinor: 13,
			},
		})
	if err != nil {
		return fmt.Errorf("ERROR: newLndServices: %v", err)
	}

	return err
}

func subscribeToEvents() error {

	_, channelErr, err := grpcServices.Client.SubscribeChannelEvents(bCtx)

	if err != nil {
		return fmt.Errorf("ERROR: subscribeChannelEvents %v", err)
	}
	if channelErr != nil {
		return fmt.Errorf("CHANNEL ERROR: subscribeChannelEvents %v", channelErr)
	}

	return err
}

func main() {
	fmt.Printf("main() Starting lnd_client_quickstart\n")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	bCtx = context.Background()
	macPath = homeDir + "/.lnd/data/chain/bitcoin/" + network + "/" + macFile
	err = setUpLNDClients()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = subscribeToEvents()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
