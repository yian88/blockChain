package block

import (
	"bytes"
	"errors"
	"github.com/bolt"
)

/**
 *
 * @author:yianzhou
 * @email:yianzhou88@gmail.com
 * @DateTime:2022/3/14 9:38
 *
 **/
type ChainIterator struct {
	DB *bolt.DB
	//标识位，表示当前迭代器所迭代到的位置
	currentHash []byte
}

//使用迭代器找上一个区块，因为是从后向前找区块
func(iterator *ChainIterator)Next()(*Block,error){
	//同一个时间允许多个人进行查看数据
	var block *Block
	var err error
	err = iterator.DB.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(BUCKET_BLOCK))
		if bk == nil{
			return errors.New("桶中没有值")
		}
		//最后一个区块的信息 []byte类型
		blockBytes := bk.Get(iterator.currentHash)
		//想获取最后一个区块的pervhash
		block, err = DeSerialize(blockBytes)
		iterator.currentHash = block.PrevHash
		return nil
	})
	return block,err

}

//判断是否还有下一个区块
func (iterator *ChainIterator)HasNext()bool{
	result := bytes.Compare(iterator.currentHash, nil)
	//创世区块
	return result != 0

}
