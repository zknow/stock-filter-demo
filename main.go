package main

import (
	"fmt"
	"strconv"
)

func main() {
	var stocks = GetStocks()
	StocksFilter(stocks)
}

func StocksFilter(_stocks []Stock) {
	stocks := make([]Stock, len(_stocks))
	copy(stocks, _stocks)

	f := NewFilter(stocks)
	f.FilterByPE(0, 15)
	f.FilterByPB(0, 1)
	f.FilterByDividendYield(8)

	for _, s := range f.Stocks {
		fmt.Printf("Code:%v\tName:%v\tPE:%v\tPB:%v\tDY:%v\n", s.Code, s.Name, s.PEratio, s.PBratio, s.DividendYield)
	}
}

func CalcAllEarningPE(_stocks []Stock) {
	stocks := make([]Stock, len(_stocks))
	copy(stocks, _stocks)

	filter := NewFilter(stocks)
	filter.FilterByPE(0, -1)
	var cnt = len(filter.Stocks)
	var totalPE float32
	for _, s := range filter.Stocks {
		pe, err := strconv.ParseFloat(s.PEratio, 32)
		if err != nil {
			continue
		}
		totalPE = totalPE + float32(pe)
	}
	fmt.Println("有獲利公司家數 :", cnt)
	fmt.Println("總PE :", totalPE)
	fmt.Println("平均PE :", totalPE/float32(cnt))
}
