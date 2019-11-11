# go-hash
封装golang中hash【md5与sha1】方法
### 使用
```cassandraql
package main

import (
	"fmt"
	"go-hash/hash"
)

func main() {
	str := "hash字符串"
	h := &hash.Hash{str}
	fmt.Println(h.HashMd5()) //e778550dfe8365ac5885eae119048ca6
	fmt.Println(h.HashSha1())//211fe1e0d6892794b0741f02caf2d79854996c4b
}

```
