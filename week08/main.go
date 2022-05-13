package main

import (
	"context"
	"fmt"
	. "github.com/hhxsv5/go-redis-memory-analysis"
	"github.com/mediocregopher/radix/v4"
)

func main() {
	ctx := context.Background()
	client, err := (radix.PoolConfig{}).New(ctx, "tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	//每种大小的value写入5w次
	write(ctx, client, 50000, "10_50000", genValue(10))     //平均每个占用11.01字节
	write(ctx, client, 50000, "20_50000", genValue(20))     //平均每个占用20.99字节
	write(ctx, client, 50000, "50_50000", genValue(50))     //平均每个占用12字节
	write(ctx, client, 50000, "100_50000", genValue(100))   //平均每个占用13字节
	write(ctx, client, 50000, "200_50000", genValue(200))   //平均每个占用13字节
	write(ctx, client, 50000, "1000_50000", genValue(1000)) //平均每个占用21.99字节
	write(ctx, client, 50000, "5000_50000", genValue(5000)) //平均每个占用67字节
	//通过info memory观察实际消耗内存350M左右

	//分析内存并输出
	analysis, err := NewAnalysisConnection("127.0.0.1", 6379, "")
	if err != nil {
		fmt.Println("something wrong:", err)
		return
	}
	defer analysis.Close()
	analysis.Start([]string{"#"})
	err = analysis.SaveReports("./reports")
	if err == nil {
		fmt.Println("done")
	} else {
		fmt.Println("error:", err)
	}
}

func write(ctx context.Context, client radix.Client, num int, key, value string) {
	for i := 0; i < num; i++ {
		newKey := fmt.Sprintf("%s#%d", key, i)
		err := client.Do(ctx, radix.Cmd(nil, "SET", newKey, value))
		if err != nil {
			fmt.Println(err)
		}
	}

}

func genValue(size int) string {
	valueArray := make([]byte, size)
	for i := 0; i < size; i++ {
		valueArray[i] = 'r'
	}
	return string(valueArray)
}
