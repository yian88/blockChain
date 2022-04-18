package main

import (
	"day0228/block"
	"fmt"
)

/**
 *
 * @author:yianzhou
 * @email:yianzhou88@gmail.com
 * @DateTime:2022/2/28 8:44
 *
 **/
func main() {
	bc,err := block.NewChain([]byte("1111"))
	defer bc.DB.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bc.LastHash)
	bc.AddBlock([]byte("12334"))
	bc.AddBlock([]byte("33333"))
	//bc.AddBlock([]byte("123456hgj"))
	iterator := bc.Iterator()
	for{
		if iterator.HasNext(){
			bk, err := iterator.Next()
			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Println(string(bk.Data))
		}else{
			break
		}
	}
	//cli := client.Cli{}
	//cli.Run()
}