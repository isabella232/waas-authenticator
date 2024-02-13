package rpc

import (
	"context"

	"github.com/0xsequence/waas-authenticator/proto"
)

func (s *RPC) ChainList(ctx context.Context) ([]*proto.Chain, error) {
	chains, err := s.Wallets.ChainList(waasContext(ctx))
	if err != nil {
		return nil, err
	}

	var retChain []*proto.Chain
	for _, chain := range chains {
		retChain = append(retChain, &proto.Chain{
			Id:        chain.Id,
			Name:      chain.Name,
			IsEnabled: chain.IsEnabled,
		})
	}
	return retChain, nil
}
