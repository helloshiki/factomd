package anchor

import (
	"fmt"
	"encoding/json"
)

var anchorCfg = struct {
	Anchor struct {
		ServerECPrivKey     string
		ServerECPublicKey   string
		AnchorChainID       string
		ConfirmationsNeeded int
	}
	Btc struct {
		BTCPubAddr         string
		SendToBTCinSeconds int
		WalletPassphrase   string
		CertHomePath       string
		RpcClientHost      string
		RpcClientEndpoint  string
		RpcClientUser      string
		RpcClientPass      string
		BtcTransFee        float64
		CertHomePathBtcd   string
		RpcBtcdHost        string
	}
}{}

func init() {
	anchorCfg.Anchor.ServerECPrivKey = "397c49e182caa97737c6b394591c614156fbe7998d7bf5d76273961e9fa1edd4"
	anchorCfg.Anchor.ServerECPublicKey = "06ed9e69bfdf85db8aa69820f348d096985bc0b11cc9fc9dcee3b8c68b41dfd5"
	anchorCfg.Anchor.AnchorChainID = "df3ade9eec4b08d5379cc64270c30ea7315d8a8a1a69efe2b98a60ecdd69e604"
	anchorCfg.Anchor.ConfirmationsNeeded = 20
	anchorCfg.Btc.WalletPassphrase                      = "test"
	anchorCfg.Btc.CertHomePath                          = "btcwallet"
	anchorCfg.Btc.RpcClientHost                         = "localhost:18554"
	anchorCfg.Btc.RpcClientEndpoint                     = "ws"
	anchorCfg.Btc.RpcClientUser                         = "myuser"
	anchorCfg.Btc.RpcClientPass                         = "SomeDecentp4ssw0rd"
	anchorCfg.Btc.BtcTransFee                           = 0.001
	anchorCfg.Btc.CertHomePathBtcd                      = "btcd"
	anchorCfg.Btc.RpcBtcdHost                           = "localhost:18556"
	s, _ := json.MarshalIndent(anchorCfg, "", "\t")
	fmt.Println("======== init anchor cfg")
	fmt.Println(string(s))
}
/*
[anchor]
ServerECKey							= 397c49e182caa97737c6b394591c614156fbe7998d7bf5d76273961e9fa1edd406ed9e69bfdf85db8aa69820f348d096985bc0b11cc9fc9dcee3b8c68b41dfd5
AnchorChainID						= df3ade9eec4b08d5379cc64270c30ea7315d8a8a1a69efe2b98a60ecdd69e604
ConfirmationsNeeded					= 20

[btc]
WalletPassphrase					= "btcdpasswd"
CertHomePath				  		= "btcwallet"
RpcClientHost				  		= "localhost:18554"
RpcClientEndpoint					= "ws"
RpcClientUser				  		= "myuser"
RpcClientPass 						= "SomeDecentp4ssw0rd"
BtcTransFee				  			= 0.0001
CertHomePathBtcd					= "btcd"
RpcBtcdHost 			  			= "localhost:18556"
RpcUser								=myuser
RpcPass								=SomeDecentp4ssw0rd
*/


/*
[anchor]
ServerECPrivKey                       = 397c49e182caa97737c6b394591c614156fbe7998d7bf5d76273961e9fa1edd4
ServerECPublicKey                     = 06ed9e69bfdf85db8aa69820f348d096985bc0b11cc9fc9dcee3b8c68b41dfd5
AnchorChainID                         = df3ade9eec4b08d5379cc64270c30ea7315d8a8a1a69efe2b98a60ecdd69e604
ConfirmationsNeeded                   = 20

[btc]
WalletPassphrase                      = "lindasilva"
CertHomePath                          = "btcwallet"
RpcClientHost                         = "localhost:18332"
RpcClientEndpoint                     = "ws"
RpcClientUser                         = "testuser"
RpcClientPass                         = "notarychain"
BtcTransFee                           = 0.000001
CertHomePathBtcd                      = "btcd"
RpcBtcdHost                           = "localhost:18334"
 */