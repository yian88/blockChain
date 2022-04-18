package transaction

import (
	"bytes"
	"day0228/tools"
	"encoding/gob"
)

/**
 *
 * @author:yianzhou
 * @email:yianzhou88@gmail.com
 * @DateTime:2022/4/18 9:41
 *
 **/
//声明整个交易的结构体
type Transcation struct {
	//交易的唯一标识
	TxId []byte
	//多个交易输出
	Output []Output
	//多个交易输入
	Input []Input
}
/**
把Transcation结构体进行序列化
 */
func (txs *Transcation)Serialize()([]byte,error){
	var result bytes.Buffer
	en := gob.NewEncoder(&result)
	err := en.Encode(txs)
	if err != nil{
		return nil,err
	}
	return result.Bytes(),nil
}
/**
创建coinbase交易
 */
func NewCoinbase(address string)(*Transcation,error){
	cb := Transcation{
		Output:[]Output{
			{
				Value: 50,
				ScriptPubkey: []byte(address),
			},
		},
		Input: nil,
	}
	txBytes, err := cb.Serialize()
	if err != nil{
		return nil,err
	}
	hash := tools.GetHash(txBytes)
	cb.TxId = hash
	return &cb,nil
}
