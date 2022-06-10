package main

import (
	"strconv"

	. "github.com/ahmetb/go-linq/v3"
)

type Filter struct {
	Codes            []string
	MinPE            float32
	MaxPE            float32
	MinPB            float32
	MaxPB            float32
	MinDividendYield float32
	Stocks           []Stock
}

func NewFilter(stocks []Stock) *Filter {
	return &Filter{Stocks: stocks}
}

func (f *Filter) FilterByCodes(codes ...string) []Stock {
	f.Codes = codes

	if len(codes) == 0 {
		return f.Stocks
	}

	var stocks []Stock
	for _, c := range codes {
		found := From(f.Stocks).Distinct().FirstWithT(func(s Stock) bool {
			return s.Code == c
		})
		if found != nil {
			stocks = append(stocks, found.(Stock))
		}
	}
	f.Stocks = stocks
	return f.Stocks
}

func (f *Filter) FilterByPE(min, max float32) []Stock {
	f.MinPE = min
	f.MaxPE = max

	From(f.Stocks).WhereT(func(s Stock) bool {
		pe, err := strconv.ParseFloat(s.PEratio, 32)
		if err != nil {
			return false
		}
		if max <= 0 {
			return min <= float32(pe)
		}
		return min <= float32(pe) && float32(pe) <= max
	}).ToSlice(&f.Stocks)
	return f.Stocks
}

func (f *Filter) FilterByPB(min, max float32) {
	f.MinPB = min
	f.MaxPB = max

	From(f.Stocks).WhereT(func(s Stock) bool {
		pb, err := strconv.ParseFloat(s.PBratio, 64)
		if err != nil {
			return false
		}
		return min <= float32(pb) && float32(pb) <= max
	}).ToSlice(&f.Stocks)
}

func (f *Filter) FilterByDividendYield(minDividendYield float32) {
	f.MinDividendYield = minDividendYield

	From(f.Stocks).WhereT(func(s Stock) bool {
		yield, err := strconv.ParseFloat(s.DividendYield, 32)
		if err != nil {
			return false
		}
		return float32(yield) >= minDividendYield
	}).ToSlice(&f.Stocks)
}
