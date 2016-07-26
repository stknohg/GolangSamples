//
// Go言語でNagios Pluginを書くサンプル
//   OSのメモリ使用量(%)をチェックするプラグイン
//
package main

import (
	"math"

	flags "github.com/jessevdk/go-flags"
	"github.com/olorin/nagiosplugin"
)

// 引数定義
type options struct {
	// 閾値設定
	Warning  int `short:"w" long:"warning" default:"80" description:"Warning threshold."`
	Critical int `short:"c" long:"critical" default:"90" description:"Critical threshold."`
	// 以下パフォーマンスデータ用
	Label string `long:"label" default:"Memory used" description:"(For PerformanceData)Label name."`
	Unit  string `long:"unit" default:"%" description:"(For PerformanceData)Unit."`
	Min   int    `long:"min" default:"0" description:"(For PerformanceData)Min value."`
	Max   int    `long:"max" default:"100" description:"(For PerformanceData)Max value."`
}

//
// main function
//
func main() {
	check := nagiosplugin.NewCheck()
	defer check.Finish()

	// 引数解析（ざっくり)
	var opts options //デバッグ用にローカル変数にしている
	_, err := flags.Parse(&opts)
	if err != nil {
		check.AddResultf(nagiosplugin.CRITICAL, "Argument error! : %s", err)
		return
	}

	// メモリ使用量(%)の取得
	current, max, err := getMemoyUsed()
	if err != nil {
		check.AddResultf(nagiosplugin.CRITICAL, "getMemoyUsed error! : %s", err)
		return
	}
	used := int(math.Ceil((float64(current) / float64(max)) * 100))

	// パフォーマンスデータの設定
	// グラフ名, 単位(ルールあり), 取得値, Warning, Critical, min, max
	err = check.AddPerfDatum(opts.Label, opts.Unit, float64(used), float64(opts.Min), float64(opts.Max), float64(opts.Warning), float64(opts.Critical))
	if err != nil {
		check.AddResultf(nagiosplugin.CRITICAL, "AddPerfDatum error : %s", err)
		return
	}

	// 閾値判定と結果の設定
	if used > opts.Critical {
		check.AddResultf(nagiosplugin.CRITICAL, "Memory used %d%%", used)
		return
	}
	if used > opts.Warning {
		check.AddResultf(nagiosplugin.WARNING, "Memory used %d%%", used)
		return
	}
	check.AddResultf(nagiosplugin.OK, "Memory used %d%%", used)
}
