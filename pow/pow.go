package pow

import (
	"bytes"
	"day0228/tools"
	"day0228/transaction"
	"fmt"
	"math/big"
	"strconv"
)

/**
 *
 * @author:yianzhou
 * @email:yianzhou88@gmail.com
 * @DateTime:2022/2/28 9:40
 *
 **/
const  BITS = 5;//难度系数，前面有多少个零

type ProofOfWork struct {
	//Block *block.Block
	//TimeStamp int64
	//PervHash []byte
	//Data []byte
	Block BlockInterface
	Target  *big.Int//256位
}

type BlockInterface interface{
	GetTimeStamp() int64
	GetPrevHash() []byte
	GetTxs() []transaction.Transcation
}
/**
实例化一个pow结构体，并且返回
 */
//func NewPow(timeStamp int64,pervHash []byte,data []byte)*ProofOfWork{
//	target := big.NewInt(1)//声明一个大整数类型1
//	target = target.Lsh(target,255-BITS)//左移多少位
//	pow := ProofOfWork{
//		TimeStamp: timeStamp,
//		PervHash : pervHash,
//		Data: data,
//		Target: target,
//	}
//	return &pow
//}

/**
实例化一个pow结构体，并且返回
*/
func NewPow(block BlockInterface)*ProofOfWork{
	target := big.NewInt(1)//声明一个大整数类型1
	target = target.Lsh(target,255-BITS)//左移多少位
	pow := ProofOfWork{
		Block: block,
		Target: target,
	}
	return &pow
}
/**
工作量证明
a<target
 */
func (pow *ProofOfWork)Run()([]byte,int64){
	var nonce int64//随机数
	nonce = 0
	//block := pow.Block
	//nonce = nonce
	timeStampBytes := []byte(strconv.FormatInt(pow.Block.GetTimeStamp(),10))
	//转型 []byte转成大整数
	num := big.NewInt(0)
	for{
		nonceBytes := []byte(strconv.FormatInt(nonce,10))
		txsBytes := []byte{}
		for _,value := range pow.Block.GetTxs(){
			txBytes, _ := value.Serialize()
			txsBytes = append(txsBytes,txBytes...)
		}
		hashBytes := bytes.Join([][]byte{pow.Block.GetPrevHash(),txsBytes,timeStampBytes,nonceBytes},[]byte{})
		hash := tools.GetHash(hashBytes)//当前区块的hash值
		fmt.Println("正在寻找nonce，当前的nonce为",nonce)
		//转型 []byte转成大整数
		num = num.SetBytes(hash)
		//    -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		if(num.Cmp(pow.Target)==-1){
			return hash,nonce
		}
		nonce++
	}
	return nil,0
}