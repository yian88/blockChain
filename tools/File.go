package tools

import "os"

/**
 *
 * @author:yianzhou
 * @email:yianzhou88@gmail.com
 * @DateTime:2022/3/28 11:25
 *判断文件是否存在
 **/
//true：文件已经存在
//false:文件不存在
func FileExist(path string)bool{
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}