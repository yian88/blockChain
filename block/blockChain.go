package block

import (
	"day0228/transaction"
	"errors"
	"github.com/bolt"
)

/**
 *
 * @author:yianzhou
 * @email:yianzhou88@gmail.com
 * @DateTime:2022/2/28 9:03
 *
 **/
const CHAIN_DB_PATH = "./chain.db"//要保存的文件地址
const BUCKET_BLOCK = "chain_blocks"//存区块桶的名字
const BUCKET_STATUS = "chain_status"//用来存最后一个区块的hash值
const LAST_HASH = "last_hash"//桶2key的名字
type BlockChain struct {
	//Blocks []*Block
	DB *bolt.DB
	LastHash []byte
}

func NewChain(address string)(*BlockChain,error){

	//bc.Blocks = []*Block{}
	//打开数据库
	db, err := bolt.Open(CHAIN_DB_PATH, 0600, nil)
	//defer db.Close()
	if err != nil{
		return nil,err
	}
	var lastHash []byte
	//向数据库中添加数据
	//update同一个时间内，只用一个人来进行写操作
	err = db.Update(func(tx *bolt.Tx) error {

		//gensis := GensisBlock(data)
		//bk, err := tx.CreateBucket([]byte(BUCKET_BLOCK))
		//if err != nil{
		//	return err
		//}
		//serialize, err := gensis.Serialize()
		//if err != nil{
		//	return err
		//}
		//bk.Put(gensis.Hash,serialize)
		////把最后一个区块的hash值放到另一个桶中
		//bk2, err := tx.CreateBucket([]byte(BUCKET_STATUS))
		//if err != nil{
		//	return err
		//}
		//bk2.Put([]byte(LAST_HASH),gensis.Hash)
		//lashHash=gensis.Hash

		//有桶
		bk := tx.Bucket([]byte(BUCKET_BLOCK))
		if bk == nil{
			coinbase,err := transaction.NewCoinbase(address)
			if err != nil{
				return err
			}
			genesis := GenesisBlock(*coinbase)//创建创世区块
			bk, err := tx.CreateBucket([]byte(BUCKET_BLOCK))
			if err !=nil{
				return err
			}
			serialize, err := genesis.Serialize()
			if err !=nil{
				return err
			}
			bk.Put(genesis.Hash,serialize)
			//创建桶2
			bk2, err := tx.CreateBucket([]byte(BUCKET_STATUS))
			bk2.Put([]byte(LAST_HASH),genesis.Hash)
			lastHash = genesis.Hash
		}else{
			bk2 := tx.Bucket([]byte(BUCKET_STATUS))
			lastHash = bk2.Get([]byte(LAST_HASH))
		}
		return nil
	})
	bc := BlockChain{
		DB: db,
		LastHash: lastHash,
	}
	//bc.Blocks = append(bc.Blocks,gensis)
	return &bc,err
}
func (bc *BlockChain)AddBlock(tx []transaction.Transcation)error{
	//defer bc.DB.Close()
	//创建区块
	//错误的地方，newblock的参数位置写反了
	newBlock := NewBlock(bc.LastHash,tx)//bc.Blocks[len(bc.Blocks)-1].Hash 获取链上最后一个的hash值
	//将创建好的区块添加到区块链中
	//bc.Blocks = append(bc.Blocks,newBlock)
	err :=	bc.DB.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(BUCKET_BLOCK))
		if bk == nil{
			return errors.New("桶中没有区块")
		}
		serialize, err := newBlock.Serialize()
		if err != nil{
			return err
		}
		bk.Put(newBlock.Hash,serialize)

		bk2 := tx.Bucket([]byte(BUCKET_STATUS))
		if bk2 == nil{
			return errors.New("桶中没有区块")
		}
		bk2.Put([]byte(LAST_HASH),newBlock.Hash)
		bc.LastHash = newBlock.Hash
		return nil
	})
	return err
}
//创建一个迭代器对象,迭代器只能再有区块链的情况下才能使用
func (bc *BlockChain)Iterator()*ChainIterator{
	iterator := ChainIterator{
		DB:bc.DB,
		currentHash: bc.LastHash,

	}
	return &iterator
}
