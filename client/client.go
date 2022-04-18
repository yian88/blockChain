package client

import (
	"day0228/block"
	"day0228/tools"
	"flag"
	"fmt"
	"os"
)

/**
 *
 * @author:yianzhou
 * @email:yianzhou88@gmail.com
 * @DateTime:2022/3/28 10:53
 *	用户程序的交互入口
 *	只负责读取用户的命令和参数，并进行解析
 *	解析后调用对应的功能
 **/
type Cli struct {


}
func(cl *Cli)Run(){
	//获取用户输入的参数
	args := os.Args
	/*确定系统有哪些功能，需要哪些参数
	a.创建带有创世区块的区块链，参数：有1个 创世区块的交易信息 string
	main.exe createChain --data "参数"
	b.添加新的区块到区块链中。参数：有1个 新区区块的交易信息 string
	main.exe addblock --data "参数"
	c.打印所有区块信息 ，参数: 无
	main.exe printblock
	d.获取当前区块链中区块的个数 参数:无
	main.exe getblockcount
	e.输出当前系统的使用说明 --help
	main.exe help
	*/
	switch args[1] {
	case "createchain":
		cl.createchain()
	case "addblock":
		cl.addblock()
	case "printblock":
		cl.printblock()
	case "getblockcount":
		cl.getblockcount();
	case "help":
		cl.help()
	default:
		fmt.Println("没有对应的功能")
		os.Exit(1)
	}
}

func (cl *Cli) createchain (){
	createchain := flag.NewFlagSet("createchain", flag.ExitOnError)
	address := createchain.String("address","","账户名称")
	createchain.Parse(os.Args[2:])
	exist := tools.FileExist("./chain.db")
	if exist {
		fmt.Println("数据库已存在")
		return
	}
	chain, err := block.NewChain(*address)
	defer chain.DB.Close()
	if err != nil{
		fmt.Println("创建失败")
		return
	}
	fmt.Println("创建成功")
}

func (cl *Cli) addblock (){
	addblock := flag.NewFlagSet("addblock", flag.ExitOnError)
	fileExist := tools.FileExist("./chain.db")
	if !fileExist{
		fmt.Println("区块链不存在")
		return
	}
	chain, err := block.NewChain("")
	defer chain.DB.Close()
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	data := addblock.String("data","","添加区块的交易信息")
	addblock.Parse(os.Args[2:])
	err = chain.AddBlock([]byte(*data))
	if err != nil{
		fmt.Println("添加失败")
		return
	}
	fmt.Println("添加成功")
}
func (cl *Cli)printblock(){
	fileExist := tools.FileExist("./chain.db")
	if !fileExist{
		fmt.Println("区块链不存在")
		return
	}
	chain, _ := block.NewChain("")
	defer chain.DB.Close()
	slice := chain.GetAllBlock()
	for _,value := range slice {
		fmt.Println(string(value.PrevHash))
	}
}
func (cl *Cli)getblockcount(){
	fileExist := tools.FileExist("./chain.db")
	if !fileExist{
		fmt.Println("区块链不存在")
		return
	}
	chain, _ := block.NewChain("")
	defer chain.DB.Close()
	silce := chain.GetAllBlock()
	fmt.Println("区块的个数有:",len(silce))
}
func (cl *Cli) help (){
	fmt.Println("createchain:用于创建一个区块链","data:表示区块中的交易信息")
	fmt.Println("addblock:用于向区块链添加一个区块","data:表示区块中的交易信息")
	fmt.Println("printblock:用于查找区块链中所用的区块")
	fmt.Println("getblockcount:用于查找区块链中的个数")
}