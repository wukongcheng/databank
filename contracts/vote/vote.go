package vote

//go:generate abigen --sol contract/vote.sol --pkg contract --out contract/vote.go

import (
	"github.com/wukongcheng/databank/accounts/abi/bind"
	"github.com/wukongcheng/databank/common"
	"github.com/wukongcheng/databank/contracts/vote/contract"
	"github.com/wukongcheng/databank/core/types"
	"github.com/wukongcheng/databank/accounts"
	"math/big"
)

var (
	MainNetAddress = common.HexToAddress("0x314159265dD8dbb310642f98f50C066173C1259b")
	TestNetAddress = common.HexToAddress("0x2d06568d53bb8fb5d26057b1c706fc38afdab06e")
)

type Vote struct {
	*contract.BallotManagerSession
	contractBackend bind.ContractBackend
}

func NewVote(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*Vote, error) {
	vote, err := contract.NewBallotManager(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &Vote{
		&contract.BallotManagerSession{
			Contract:     vote,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func GetVote(accMng *accounts.Manager, backend bind.ContractBackend, address common.Address, nonce uint64) (*Vote, error) {

	account := accounts.Account{Address: address}
	wallet, err := accMng.Find(account)
	if err != nil {
		return nil, err
	}

	transactOpts, err := wallet.NewUnlockedKeyedTransactor(account,nonce)
	if err != nil {
		return nil, err
	}

	contract, err := NewVote(transactOpts, TestNetAddress, backend)
	if err != nil {
		return nil, err
	}

	return contract, nil
}

func DeployVote(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *Vote, error) {
	voteAddr, _, _, err := contract.DeployBallotManager(transactOpts, contractBackend)
	if err != nil {
		return voteAddr, nil, err
	}

	xcdata, err := NewVote(transactOpts, voteAddr, contractBackend)
	if err != nil {
		return voteAddr, nil, err
	}

	return voteAddr, xcdata, nil
}

func (self *Vote) NewBallot(ballotname string, proposalNames [][32]byte) (*types.Transaction, error) {
	return self.Contract.NewBallot(&self.TransactOpts, ballotname, proposalNames)
}

func (self *Vote) Vote(ballotId *big.Int, proposal *big.Int) (*types.Transaction, error) {
	return self.Contract.Vote(&self.TransactOpts, ballotId, proposal)
}

func (self *Vote) GetWinningProposal(ballotId *big.Int) (*big.Int, error) {
	return self.Contract.WinningProposal(&self.CallOpts, ballotId);
}

func (self *Vote) GetWinnerName(ballotId *big.Int) ([32]byte, error) {
	return self.Contract.WinnerName(&self.CallOpts, ballotId);
}

