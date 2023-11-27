package main

import (
	"github.com/codechalls/core/common"
	pkg "github.com/codechalls/pkg/wc"
)

func main() {

	err := common.DisplayLogo()
	if err != nil {
		return
	}

	pkg.WordCount()
}
