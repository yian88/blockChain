package tools

import "crypto/sha256"

/**
 *
 * @author:yianzhou
 * @email:yianzhou88@gmail.com
 * @DateTime:2022/2/28 8:48
 *
 **/
/*
使用hashsha256计算哈希值
 */
func GetHash(data []byte)[]byte{
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}