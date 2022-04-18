package block

import (
	"bytes"
	pow2 "day0228/pow"
	"day0228/tools"
	"day0228/transaction"
	"encoding/gob"
	"fmt"
	"strconv"
	"time"
)

/**
 *
 * @author:yianzhou
 * @email:yianzhou88@gmail.com
 * @DateTime:2022/2/28 8:44
 *
 **/
type Block struct {
	PrevHash []byte //上一个区块的hash值
	TimeStamp int64 //时间戳
	Nonce int64 //随机数
	//Data []byte //交易信息
	Txs []transaction.Transcation
	Hash []byte //hash值
}

func(block *Block)GetTimeStamp() int64 {
	return block.TimeStamp
}

func(block *Block)GetPrevHash() []byte {
	return block.PrevHash
}

func(block *Block)GetTxs() []transaction.Transcation {
	return block.Txs
}
/**
创建一个区块
 */
func NewBlock(prevHash []byte,txs []transaction.Transcation)*Block{
	//给block结构体赋值
	block := Block{
		PrevHash: prevHash,//上一个区块值
		TimeStamp: time.Now().Unix(),
		Txs: txs,
	}
	//实现了blockInterface接口
	pow := pow2.NewPow(&block)//工作量证明
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return &block

}
/**
序列化 把目标结构体转成一个有序的排列
json xml
 */
func (block *Block)Serialize()([]byte,error){
	var result bytes.Buffer
	en := gob.NewEncoder(&result)
	err := en.Encode(block)
	if err != nil{
		return nil,err
	}
	return result.Bytes(),nil
}
/**
反序列化 把字节切片转成结构体
 */
func DeSerialize(data []byte)(*Block,error){
	reader := bytes.NewReader(data)
	de := gob.NewDecoder(reader)
	var block *Block
	err := de.Decode(&block)
	if err != nil{
		return nil,err
	}
	return block,nil
}
/**
获取当前的区块值
 */
func (b *Block)SetHash()([]byte,error){
	//区块的hash:时间戳+上一个区块的hash值+交易信息+随机数
	timeStamp := []byte(strconv.FormatInt(b.TimeStamp,10))
	nonce := []byte(strconv.FormatInt(b.Nonce,10))
	txsBytes := []byte{}
	for _,value := range b.Txs{
		txBytes, err := value.Serialize()
		if err != nil{
			return nil,err
		}
		txsBytes = append(txsBytes,txBytes...)

	}
	hash := bytes.Join([][]byte{b.PrevHash,txsBytes,timeStamp,nonce},[]byte{})
	return tools.GetHash(hash),nil
}
/**
创建创世区块
 */
func GenesisBlock(tx transaction.Transcation)*Block{
	return NewBlock(nil,[]transaction.Transcation{tx})
}
func (bc *BlockChain)GetAllBlock()[]*Block{
	iterator := bc.Iterator()
	bk := []*Block{}
	for {
		if iterator.HasNext() {
			bl, err := iterator.Next()
			if err != nil {
				fmt.Println("失败了")
				return nil;
			}
			bk = append(bk, bl)
		}else {
			break
		}

	}
	return bk
}