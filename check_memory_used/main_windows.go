package main

import "github.com/StackExchange/wmi"

//
// WMIを使ってOSのメモリ使用量を取得します。
// 戻り値 : int64 現在の使用量(byte), int64 OSのメモリサイズ(byte), error エラー
//
func getMemoyUsed() (int64, int64, error) {
	// WMI Query
	type Win32_OperatingSystem struct {
		FreePhysicalMemory     int64
		TotalVisibleMemorySize int64
	}
	var ret []Win32_OperatingSystem
	query := wmi.CreateQuery(&ret, "")
	err := wmi.Query(query, &ret)
	if err != nil {
		return -1, -1, err
	}
	return (ret[0].TotalVisibleMemorySize - ret[0].FreePhysicalMemory), ret[0].TotalVisibleMemorySize, nil
}
