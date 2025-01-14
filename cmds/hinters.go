package cmds

import (
	credentialcmds "github.com/ProtoconNet/mitum-credential/cmds"
	currencycmds "github.com/ProtoconNet/mitum-currency/v3/cmds"
	daocmds "github.com/ProtoconNet/mitum-dao/cmds"
	didcmds "github.com/ProtoconNet/mitum-did-registry/cmds"
	nftcmds "github.com/ProtoconNet/mitum-nft/cmds"
	pointcmds "github.com/ProtoconNet/mitum-point/cmds"
	storagecmds "github.com/ProtoconNet/mitum-storage/cmds"
	timestampcmds "github.com/ProtoconNet/mitum-timestamp/cmds"
	tokencmds "github.com/ProtoconNet/mitum-token/cmds"
	"github.com/ProtoconNet/mitum2/util/encoder"
	"github.com/pkg/errors"
)

var Hinters []encoder.DecodeDetail
var SupportedProposalOperationFactHinters []encoder.DecodeDetail

func init() {
	AllHinters := [][]encoder.DecodeDetail{
		currencycmds.Hinters,
		nftcmds.AddedHinters,
		timestampcmds.AddedHinters,
		credentialcmds.AddedHinters,
		tokencmds.AddedHinters,
		pointcmds.AddedHinters,
		daocmds.AddedHinters,
		storagecmds.AddedHinters,
		didcmds.AddedHinters,
	}

	for i := range AllHinters {
		Hinters = append(Hinters, AllHinters[i]...)
	}

	AllSupportedFactHinters := [][]encoder.DecodeDetail{
		currencycmds.SupportedProposalOperationFactHinters,
		nftcmds.AddedSupportedHinters,
		timestampcmds.AddedSupportedHinters,
		credentialcmds.AddedSupportedHinters,
		tokencmds.AddedSupportedHinters,
		pointcmds.AddedSupportedHinters,
		daocmds.AddedSupportedHinters,
		storagecmds.AddedSupportedHinters,
		didcmds.AddedSupportedHinters,
	}

	for i := range AllSupportedFactHinters {
		SupportedProposalOperationFactHinters = append(SupportedProposalOperationFactHinters, AllSupportedFactHinters[i]...)
	}
}

func LoadHinters(encs *encoder.Encoders) error {
	for i := range Hinters {
		if err := encs.AddDetail(Hinters[i]); err != nil {
			return errors.Wrap(err, "add hinter to encoder")
		}
	}

	for i := range SupportedProposalOperationFactHinters {
		if err := encs.AddDetail(SupportedProposalOperationFactHinters[i]); err != nil {
			return errors.Wrap(err, "add supported proposal operation fact hinter to encoder")
		}
	}

	return nil
}
