package main

import (
	"fmt"
	"vesgo/cmd"
)

/*
@Author : Vespeng
@Time   : 2025/1/20
@Desc   : 主函数
*/

const art = `___    __                         
__ |  / /__________________ ______ 
__ | / /_  _ \_  ___/_  __ \'/  __ \
__ |/ / /  __/(__  )_  /_/ // /_/ /
_____/  \___//____/ _\__, / \____/ 
                    /____/         `

func main() {
	fmt.Println(art)
	cmd.Run()
}
