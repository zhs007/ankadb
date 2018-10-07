package main

import (
	"encoding/json"
	"fmt"

	pb "github.com/zhs007/ankadb/samples/trading/proto"
)

func main() {
	mapvar := make(map[string]interface{})

	c := pb.Candle{
		Open:         1,
		Close:        2,
		High:         3,
		Low:          4,
		Volume:       5,
		OpenInterest: 6,
		Curtime:      7,
	}

	c1 := make(map[string]interface{})
	c1["open"] = 1
	c1["close"] = 2
	c1["high"] = 3
	c1["low"] = 4
	c1["volume"] = 5
	c1["openInterest"] = 6
	c1["curtime"] = 7

	lstc := make([]pb.Candle, 1)
	lstc[0] = c

	mapvar["code"] = "pta"
	mapvar["name"] = "pta1801"
	mapvar["candles"] = lstc
	mapvar["endTime"] = 0
	mapvar["startTime"] = 0
	mapvar["candle"] = c1

	r := executeQuery(`
	mutation InsertCandles($code:String!, $name:String!, $candle:CandleInput){
		insertCandles(code:$code,name:$name,candle:$candle){
			code,
			name}
		}`, mapvar, schema)
	executeQuery("query GetAllData($code:String!, $name:String!, $startTime:Timestamp, $endTime:Timestamp){candleChunks(code:$code,name:$name,startTime:$startTime,endTime:$endTime){open,close}}", mapvar, schema)
	b, _ := json.Marshal(r)
	fmt.Println(string(b))
}
