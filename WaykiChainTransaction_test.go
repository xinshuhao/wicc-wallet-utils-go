package wiccwallet

import (
	"encoding/hex"
	"encoding/json"
	"github.com/WaykiChain/wicc-wallet-utils-go/commons"
	"io/ioutil"
	"testing"
)

/*
  * 多币种转账交易 ,支持多种币种转账
  * Test nUniversal Coin Transfer Tx
  * fee Minimum 100000 sawi
  * */
func TestSignUCoinTransferTx(t *testing.T) {
	privateKey := "Y6J4aK6Wcs4A3Ex4HXdfjJ6ZsHpNZfjaS4B9w7xqEnmFEYMqQd13"
	var txParam UCoinTransferTxParam
	txParam.FeeSymbol = string(commons.WICC)
	txParam.Fees = 1000000
	txParam.ValidHeight = 3745312
	txParam.SrcRegId = ""
	txParam.Dests=NewDestArr()
	dest1:=Dest{string(commons.WICC),1000000, "0-1"}
	dest2:=Dest{string(commons.WICC),1000000, "whK6Jp4CFqgfB8ZUgXeTy2GnGgwTKWpG4d"}
	txParam.Dests.Add(&dest1)
	txParam.Dests.Add(&dest2)
	txParam.PubKey = "036c5397f3227a1e209952829d249b7ad0f615e43b763ac15e3a6f52627a10df21"
	txParam.Memo = ""
	hash, err := SignUCoinTransferTx(privateKey, &txParam)
	if err != nil {
		t.Error("UCoinTransferTx err: ", err)
	}
	println(hash)
}

/*
  * 多币种合约调用交易 ,支持多种币种转账
  * Test nUniversal Coin Contract Tx
  * fee Minimum 1000000 sawi
  * */
func TestSignUCoinCallContractTx(t *testing.T) {
	privateKey := "Y6J4aK6Wcs4A3Ex4HXdfjJ6ZsHpNZfjaS4B9w7xqEnmFEYMqQd13"
	var txParam UCoinContractTxParam
	txParam.FeeSymbol = string(commons.WICC)
	txParam.CoinSymbol = string(commons.WICC)
	txParam.CoinAmount = 1000000
	txParam.Fees = 1000000
	txParam.ValidHeight = 297449
	txParam.SrcRegId = "0-1"
	txParam.AppId = "0-1"
	txParam.PubKey = "036c5397f3227a1e209952829d249b7ad0f615e43b763ac15e3a6f52627a10df21"
	txParam.ContractHex = "f017"
	hash, err := SignUCoinCallContractTx(privateKey, &txParam)
	if err != nil {
		t.Error("SignUCoinCallContractTx err: ", err)
	}
	println(hash)
}


/*
	多币种部署合约,推荐使用
	fee Minimum: 110000000 sawi
*/
func TestSignUCoinDeployContractTx(t *testing.T) {

	privateKey := "YAa1wFCfFnZ5bt4hg9MDeDevTMd1Nu874Mn83hEXwtfAL2vkQE9f" //"Y9sx4Y8sBAbWDAqAWytYuUnJige3ZPwKDZp1SCDqqRby1YMgRG9c"

	script, err := ioutil.ReadFile("./demo/data/hello.lua")
	if err != nil {
		t.Error("Read contract script file err: ", err)
	}

	var txParam UCoinRegisterContractTxParam
	txParam.ValidHeight = 630314 //20999
	txParam.SrcRegId = "0-1"     //"7849-1"
	txParam.Fees = 110000000
	txParam.Script = script
	txParam.Description = "My hello contract!!!"
	txParam.FeeSymbol = string(commons.WICC)

	hash, err := SignUCoinRegisterContractTx(privateKey, &txParam)
	if err != nil {
		t.Error("UCoinRegisterContractTx err: ", err)
	}
	println(hash)
}

/*
   部署合约 ,不推荐
   fee Minimum :110000000 sawi
*/
func TestSignDeployContractTx(t *testing.T) {

	privateKey := "YAa1wFCfFnZ5bt4hg9MDeDevTMd1Nu874Mn83hEXwtfAL2vkQE9f" //"Y9sx4Y8sBAbWDAqAWytYuUnJige3ZPwKDZp1SCDqqRby1YMgRG9c"

	script, err := ioutil.ReadFile("./demo/data/hello.lua")
	if err != nil {
		t.Error("Read contract script file err: ", err)
	}

	var txParam RegisterContractTxParam
	txParam.ValidHeight = 630314 //20999
	txParam.SrcRegId = "0-1"     //"7849-1"
	txParam.Fees = 110000000
	txParam.Script = script
	txParam.Description = "My hello contract!!!"
	hash, err := SignRegisterContractTx(privateKey, &txParam)
	if err != nil {
		t.Error("SignRegisterContractTx err: ", err)
	}
	println(hash)
}

/*
Asset Create 创建资产

fee Minimum 0.01WICC
make sure account balance >550.01 wicc
发布资产需要扣除550wicc的发布费用
*/
func TestAssetCreateTx(t *testing.T) {
	privateKey := "Y6J4aK6Wcs4A3Ex4HXdfjJ6ZsHpNZfjaS4B9w7xqEnmFEYMqQd13"
	var txParam AssetIssueTxParam
	txParam.FeeSymbol = string(commons.WICC)
	txParam.Fees = 1000000
	txParam.ValidHeight = 11625
	txParam.SrcRegId = "0-1"
	txParam.AssetOwner="0-1"//only support regid
	txParam.AssetTotal=100*100000000
	txParam.AssetSymbol="SSSSSSS"
	txParam.AssetName="SK Token"
	txParam.MinTable=true
	hash, err := SignAssetCreateTx(privateKey, &txParam)
	if err != nil {
		t.Error("SignAssetIssueTx err: ", err)
	}
	println(hash)
}

/*
Asset Update 更新资产
fee Minimum 0.01WICC
make sure account balance >110.01 wicc
发布资产需要扣除110wicc的发布费用
*/
func TestAssetUpdateTx(t *testing.T) {
	privateKey := "Y6J4aK6Wcs4A3Ex4HXdfjJ6ZsHpNZfjaS4B9w7xqEnmFEYMqQd13"
	var txParam AssetUpdateTxParam
	txParam.FeeSymbol = string(commons.WICC)
	txParam.Fees = 1000000
	txParam.ValidHeight = 688709
	txParam.SrcRegId = "0-1"
	txParam.UpdateType=int(commons.ASSET_OWNER_UID)
	txParam.AssetSymbol="STOKEN"
	txParam.AssetOwner="111-1" //only support regid
	//txParam.AssetTotal=100*100000000
	//txParam.AssetName="SK Token"
	hash, err := SignAssetUpdateTx(privateKey, &txParam)
	if err != nil {
		t.Error("AssetUpdateTx err: ", err)
	}
	println(hash)
}


/*
* 投票交易
* Voting transaction
* fee Minimum 0.01 wicc
* */
func TestSignDelegateTx(t *testing.T) {
	privateKey := "Y9sx4Y8sBAbWDAqAWytYuUnJige3ZPwKDZp1SCDqqRby1YMgRG9c"
	var txParams DelegateTxParam
	txParams.ValidHeight = 95728
	txParams.SrcRegId = ""
	txParams.Fees = 1000000
	txParams.PubKey = "03e93e7d870ce6f1c9997076c56fc24e6381c612662cd9a5a59294fac9ba7d21d7"
	txParams.Votes = NewOperVoteFunds()
	pubKey, _ := hex.DecodeString("025a37cb6ec9f63bb17e562865e006f0bafa9afbd8a846bd87fc8ff9e35db1252e")
	vote := OperVoteFund{PubKey: pubKey, VoteValue: 10000}
	txParams.Votes.Add(&vote)

	hash, err := SignDelegateTx(privateKey, &txParams)
	if err != nil {
		t.Error("SignDelegateTx err: ", err)
	}
	println(hash)
}


/*
    * 创建,追加cdp交易
    * Create or append an  cdp transaction
    * fee Minimum 0.01 wicc
    * */
func TestCdpStakeTx(t *testing.T) {
	privateKey := "Y6J4aK6Wcs4A3Ex4HXdfjJ6ZsHpNZfjaS4B9w7xqEnmFEYMqQd13"
	var txParam CdpStakeTxParam
	txParam.CdpTxid = "0b9734e5db3cfa38e76bb273dba4f65a210cc76ca2cf739f3c131d0b24ff89c1"
	txParam.ScoinSymbol = string(commons.WUSD)
	txParam.FeeSymbol = string(commons.WICC)
	AssetSymbol:= string(commons.WICC)
	AssetAmount:= 1000
	model:=AssetModel{int64(AssetAmount),AssetSymbol}
	txParam.Assets=NewCdpAssets()
	txParam.Assets.Add(&model)
	txParam.ScoinMint = 10
	txParam.Fees = 10000000
	txParam.ValidHeight = 283308
	txParam.SrcRegId = "0-1"
	txParam.PubKey = "03e93e7d870ce6f1c9997076c56fc24e6381c612662cd9a5a59294fac9ba7d21d7"

	hash, err := SignCdpStakeTx(privateKey, &txParam)
	if err != nil {
		t.Error("CdpStakeTx err: ", err)
	}
	println(hash)
}

/*
  * 赎回cdp交易
  * Redeem cdp transaction
  * fee Minimum 0.01 wicc
  * */
func TestSignCdpRedeemTx(t *testing.T) {
	privateKey := "Y6J4aK6Wcs4A3Ex4HXdfjJ6ZsHpNZfjaS4B9w7xqEnmFEYMqQd13"
	var txParam CdpRedeemTxParam
	txParam.CdpTxid = "0b9734e5db3cfa38e76bb273dba4f65a210cc76ca2cf739f3c131d0b24ff89c1"
	txParam.FeeSymbol = string(commons.WICC)
	txParam.ScoinsToRepay = 0
	AssetSymbol:= string(commons.WICC)
	AssetAmount:= 1000
	model:=AssetModel{int64(AssetAmount),AssetSymbol}
	txParam.Assets=NewCdpAssets()
	txParam.Assets.Add(&model)
	txParam.Fees = 1000000
	txParam.ValidHeight = 25
	txParam.SrcRegId = "0-1"
	txParam.PubKey = "03e93e7d870ce6f1c9997076c56fc24e6381c612662cd9a5a59294fac9ba7d21d7"

	hash, err := SignCdpRedeemTx(privateKey, &txParam)
	if err != nil {
		t.Error("CdpRedeemTx err: ", err)
	}
	println(hash)
}

/*
  * 清算cdp交易
  * Liquidate cdp transaction
  * fee Minimum 0.01 wicc
  * */
func TestSignCdpLiquidateTx(t *testing.T) {
	privateKey := "Y6J4aK6Wcs4A3Ex4HXdfjJ6ZsHpNZfjaS4B9w7xqEnmFEYMqQd13"
	var txParam CdpLiquidateTxParam
	txParam.CdpTxid = "0b9734e5db3cfa38e76bb273dba4f65a210cc76ca2cf739f3c131d0b24ff89c1"
	txParam.FeeSymbol = string(commons.WICC)
	txParam.ScoinsLiquidate = 100
	txParam.Fees = 1000000
	txParam.ValidHeight = 25
	txParam.SrcRegId = "0-1"
	txParam.PubKey = "03e93e7d870ce6f1c9997076c56fc24e6381c612662cd9a5a59294fac9ba7d21d7"
	txParam.AssetSymbol=string(commons.WICC)

	hash, err := SignCdpLiquidateTx(privateKey, &txParam)
	if err != nil {
		t.Error("CdpLiquidateTx err: ", err)
	}
	println(hash)
}

/*
 * Dex 限价买单交易
 * Dex limit price transaction
 * fee Minimum 0.001 wicc
 * */
func TestSignDexBuyLimitTx(t *testing.T) {
	privateKey := "Y6J4aK6Wcs4A3Ex4HXdfjJ6ZsHpNZfjaS4B9w7xqEnmFEYMqQd13"
	var txParam DexLimitTxParam
	txParam.FeeSymbol = string(commons.WICC)
	txParam.Fees = 1000000
	txParam.CoinSymbol = string(commons.WUSD)
	txParam.AssetSymbol = string(commons.WICC)
	txParam.AssetAmount = 10000
	txParam.ValidHeight = 25
	txParam.Price = 25
	txParam.SrcRegId = "0-1"
	txParam.PubKey = "03e93e7d870ce6f1c9997076c56fc24e6381c612662cd9a5a59294fac9ba7d21d7"

	hash, err := SignDexBuyLimitTx(privateKey, &txParam)
	if err != nil {
		t.Error("DexBuyLimitTx err: ", err)
	}
	println(hash)
}

/*
   * Dex 限价卖单交易
   * Dex limit sell price transaction
   * fee Minimum 0.001 wicc
  * */
func TestSignDexSellLimitTx(t *testing.T) {
	privateKey := "Y6J4aK6Wcs4A3Ex4HXdfjJ6ZsHpNZfjaS4B9w7xqEnmFEYMqQd13"
	var txParam DexLimitTxParam
	txParam.FeeSymbol = string(commons.WICC)
	txParam.Fees = 100000
	txParam.CoinSymbol = string(commons.WUSD)
	txParam.AssetSymbol = string(commons.WICC)
	txParam.AssetAmount = 1000000
	txParam.ValidHeight = 282956
	txParam.Price = 200000000
	txParam.SrcRegId = "0-1"
	txParam.PubKey = "03e93e7d870ce6f1c9997076c56fc24e6381c612662cd9a5a59294fac9ba7d21d7"

	hash, err := SignDexSellLimitTx(privateKey, &txParam)
	if err != nil {
		t.Error("DexSellLimitTx err: ", err)
	}
	println(hash)
}

/*
     *  Dex 市价卖单交易
     * Dex market sell price transaction
     * fee Minimum 0.001 wicc
    * */
func TestSignDexMarketSellTx(t *testing.T) {
	privateKey := "Y6J4aK6Wcs4A3Ex4HXdfjJ6ZsHpNZfjaS4B9w7xqEnmFEYMqQd13"
	var txParam DexMarketTxParam
	txParam.FeeSymbol = string(commons.WICC)
	txParam.Fees = 1000000
	txParam.CoinSymbol = string(commons.WUSD)
	txParam.AssetSymbol = string(commons.WICC)
	txParam.AssetAmount = 10000
	txParam.ValidHeight = 25
	txParam.SrcRegId = "0-1"
	txParam.PubKey = "03e93e7d870ce6f1c9997076c56fc24e6381c612662cd9a5a59294fac9ba7d21d7"

	hash, err := SignDexMarketSellTx(privateKey, &txParam)
	if err != nil {
		t.Error("DexMarketSellTx err: ", err)
	}
	println(hash)
}

/*
     *  Dex 市价买单交易
     * Dex market buy price transaction
     * fee Minimum 0.001 wicc
    * */
func TestSignDexMarketBuyTx(t *testing.T) {
	privateKey := "Y6J4aK6Wcs4A3Ex4HXdfjJ6ZsHpNZfjaS4B9w7xqEnmFEYMqQd13"
	var txParam DexMarketTxParam
	txParam.FeeSymbol = string(commons.WICC)
	txParam.Fees = 1000000
	txParam.CoinSymbol = string(commons.WUSD)
	txParam.AssetSymbol = string(commons.WICC)
	txParam.AssetAmount = 10000
	txParam.ValidHeight = 25
	txParam.SrcRegId = "0-1"
	txParam.PubKey = "03e93e7d870ce6f1c9997076c56fc24e6381c612662cd9a5a59294fac9ba7d21d7"

	hash, err := SignDexMarketBuyTx(privateKey, &txParam)
	if err != nil {
		t.Error("DexMarketBuyTx err: ", err)
	}
	println(hash)
}

/*
 *  Dex 取消挂单交易
 * Dex cancel order tx
 * fee Minimum 0.001 wicc
* */
func TestSignDexCancelTx(t *testing.T) {
	privateKey := "Y6J4aK6Wcs4A3Ex4HXdfjJ6ZsHpNZfjaS4B9w7xqEnmFEYMqQd13"
	var txParam DexCancelTxParam
	txParam.FeeSymbol = string(commons.WICC)
	txParam.Fees = 100000
	txParam.DexTxid = "009c0e665acdd9e8ae754f9a51337b85bb8996980a93d6175b61edccd3cdc144"
	txParam.ValidHeight = 25
	txParam.SrcRegId = "0-1"
	txParam.PubKey = "03e93e7d870ce6f1c9997076c56fc24e6381c612662cd9a5a59294fac9ba7d21d7"

	hash, err := SignDexCancelTx(privateKey, &txParam)
	if err != nil {
		t.Error("DexCancelTx err: ", err)
	}
	println(hash)
}

// Support UCOIN_TRANSFER_TX and UCOIN_CONTRACT_INVOKE_TX only
func TestDecodeRawTx(t *testing.T){
	//rawTx := "0b01df390684f0c10c82480457494343bc834002141c758724cc60db35dd387bcf619a478ec3c065f20457494343bc8340142af03ec43eb893039b5dd5bab612d73034cf1b610457555344858c1f00473045022100d68782ebf4059ac26b169ae035ca2a8c1533c4f5639c9fd64445f205d86fbf2c022008b7ed1467ec9321382284ce9d762967a604602a26295f4d569f9a15b643e1db"
	//rawTx := "0b01df3921036c5397f3227a1e209952829d249b7ad0f615e43b763ac15e3a6f52627a10df210457494343bc83400214079b9296a00a2b655787fa90e66ec3cde4bf1c8c0457494343bc834014079b9296a00a2b655787fa90e66ec3cde4bf1c8c0457494343866700473045022100a80f36f5b260bdbb76a46f0d9563bdfe0c79a8c2f7b5935b5e79a3c280040551022001ba6dee175619b2c3ba19976690d848ffdbfb4df596bdf866932a4b7befb022"
	rawTx := "0b01df3921036c5397f3227a1e209952829d249b7ad0f615e43b763ac15e3a6f52627a10df210457494343bc8340020200010457494343bc834014eddae751c608603c933778a0b2466816bc928d790457494343bc83400047304502210098611641f54926ec74885348013616c9d25a0f8e07a537e9ff59f1d1c0f0e568022021c7c3e55c437fa0f67ca5e6796a0df4577e7c175ef9b57310c6a4be3b649275"
	result ,err := commons.DecodeRawTx(rawTx,WAYKI_TESTNET)
	if err != nil {
		t.Error("Umarshal failed:", err)
	}
	jsonBytes, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		t.Error("Umarshal failed:", err)
	}

	t.Log("DecodeRawTx result=\n",string(jsonBytes))
}

