package main

import (
	"fmt"
)

type foo struct {
	ValFirst   int    `json:"val_first"` //v
	FirstimStr string `json:"firstim_str"`
}

type deli_yurek struct {
	val_second   string `json: "val_second" xml: "val_second" `
	HebeleMEBELE string `json: "HebeleMEBELE" xml: "HebeleMEBELE" `
}
type kv_cakir struct {
	Polat                int    `json: "Polat" xml: "Polat" `
	Memati               int    `json: "Memati" xml: "Memati" `
	CennetMahallesiYunus string `json: "CennetMahallesiYunus" xml: "CennetMahallesiYunus" `
	Arka_sokaklar_husnu  bool   `json: "Arka_sokaklar_husnu" xml: "Arka_sokaklar_husnu" `
}

type marvel struct {
	foo
	deli_yurek
	IronMan    struct{} `json: "IronMan" xml: "IronMan" `
	JackSparow string   `json: "JackSparow" xml: "JackSparow" `
}

func forShow() {
	fmt.Println("forShow")
}
