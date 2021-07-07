package main

import (
	"testing"
)

func TestFigureReturn01(t *testing.T) {
	main() //標準入力から入れていると模す
	//出力確認しました
}

func TestPlus01(t *testing.T) {
	data.push(3)
	data.push(8)
	data.push("+")
	if data.pop() != 11 {
		t.Errorf("....")
	}
}
