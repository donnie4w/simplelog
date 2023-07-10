package logging

import (
	"testing"
)

func BenchmarkSerialLog(b *testing.B) {
	SetRollingFile(`D:\cfoldTest`, "log1.txt", 100, MB)
	SetConsole(false)
	// SetFormat(FORMAT_NANO)
	for i := 0; i < b.N; i++ {
		Debug(i, ">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		Info(i, ">>>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
		Warn(i, ">>>cccccccccccccccccccccccccccccccccccc")
		Error(i, ">>>dddddddddddddddddddddddddddddddddddd")
	}
}

func BenchmarkParallelLog(b *testing.B) {
	log := NewLogger()
	// log.SetRollingFile(`D:\cfoldTest`, "log.txt", 1000, KB)
	// log.SetRollingFileLoop(`D:\cfoldTest`, "logLoop.txt", 2000, KB, 5)
	log.SetRollingDaily("", "log.txt")
	log.SetConsole(false)
	// b.SetParallelism(8)
	i := 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i++
			log.Debug(i, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			// log.Info(i, "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
			// log.Warn(i, "cccccccccccccccccccccccccccccccccccc")
			// log.Error(i, "dddddddddddddddddddddddddddddddddddd")
			// log.Fatal(i, "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
		}
	})
}
