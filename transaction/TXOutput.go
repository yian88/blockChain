package transaction

/**
 *
 * @author:yianzhou
 * @email:yianzhou88@gmail.com
 * @DateTime:2022/4/18 9:29
 *
 **/
type Output struct {
	//描述输出的币的金额
	Value int
	//锁定脚本
	ScriptPubkey []byte

}
