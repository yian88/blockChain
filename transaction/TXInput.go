package transaction

/**
 *
 * @author:yianzhou
 * @email:yianzhou88@gmail.com
 * @DateTime:2022/4/18 9:39
 *
 **/
//交易输入的结构体
type Input struct {
	//确定要消费的交易输出在那个交易中
	Txid []byte
	//交易输出索引位置
	Vout int
	//解锁脚本
	ScriptSig []byte

}