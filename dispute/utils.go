// Copyright 2018-2020 Celer Network

package dispute

import (
	"bytes"

	"github.com/celer-network/goCeler/chain"
	"github.com/celer-network/goCeler/ctype"
	"github.com/celer-network/goCeler/entity"
	"github.com/celer-network/goCeler/rpc"
	"github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/golang/protobuf/proto"
)

func SigSortedSimplexState(state *rpc.SignedSimplexState) (*chain.SignedSimplexState, error) {
	var signedState chain.SignedSimplexState
	signedState.SimplexState = make([]byte, len(state.SimplexState))
	copy(signedState.SimplexState, state.SimplexState)
	peerFrom, err := eth.RecoverSigner(state.SimplexState, state.SigOfPeerFrom)
	if err != nil {
		return nil, err
	}
	peerTo, err := eth.RecoverSigner(state.SimplexState, state.SigOfPeerTo)
	if err != nil {
		return nil, err
	}
	if bytes.Compare(peerFrom.Bytes(), peerTo.Bytes()) < 0 {
		signedState.Sigs = append(signedState.Sigs, state.SigOfPeerFrom)
		signedState.Sigs = append(signedState.Sigs, state.SigOfPeerTo)
	} else {
		signedState.Sigs = append(signedState.Sigs, state.SigOfPeerTo)
		signedState.Sigs = append(signedState.Sigs, state.SigOfPeerFrom)
	}
	return &signedState, nil
}

func PrintSignedSimplexState(state *chain.SignedSimplexState) {
	log.Infoln("-- Print Simplex State")
	log.Infof("---- state bytes %x", state.SimplexState)
	var simplex entity.SimplexPaymentChannel
	if proto.Unmarshal(state.SimplexState, &simplex) != nil {
		log.Errorf("unmarshal err for simplex: %x", state.SimplexState)
	}
	log.Infoln("---- channel ID", ctype.Bytes2Cid(simplex.ChannelId).Hex())
	log.Infoln("---- peer from", ctype.Bytes2Hex(simplex.PeerFrom))
	log.Infoln("---- seq num", simplex.SeqNum)
	log.Infoln("---- token transfer addr", ctype.Bytes2Hex(simplex.TransferToPeer.Token.TokenAddress))
	log.Infoln("---- token transfer amount", ctype.Bytes2Hex(simplex.TransferToPeer.Receiver.Amt))
	log.Infoln("---- pending pay IDs", simplex.PendingPayIds)
	log.Infoln("---- last resolve deadline", simplex.LastPayResolveDeadline)
	for _, sig := range state.Sigs {
		signer, err := eth.RecoverSigner(state.SimplexState, sig)
		if err != nil {
			log.Error(err)
		} else {
			log.Infoln("---- signer", ctype.Addr2Hex(signer))
		}
	}
}
