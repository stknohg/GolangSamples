package main

import "github.com/guillermo/go.procmeminfo"

//
// /proc/meminfoの内容からOSのメモリ使用量を取得します。
// 戻り値 : int64 現在の使用量(byte), int64 OSのメモリサイズ(byte), error エラー
//
func getMemoyUsed() (int64, int64, error) {
	meminfo := &procmeminfo.MemInfo{}
	err := meminfo.Update()
	if err != nil {
		return -1, -1, err
	}
	// Available = MemFree + Buffers + Cached
	// Total     = MemTotal
	total := int64(meminfo.Total())
	used := total - int64(meminfo.Available())
	return used, total, nil
}
