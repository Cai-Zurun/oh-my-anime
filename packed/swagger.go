package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GBQ8BYLZEACIgycDMXlienpqUX6UFovqzg/LzSElYEx94lD/KpnQQEBAQFaev4+/v66QZuCAowuGQdcauxIe7b0qXakpsOVBtdJnz97CU32dmjaasXxcM6ZJ1f2PArZYBAhorUyZELEpEAewXWsLAwM//8HeLNzZF/cl+fDwMCQzcDAAHMWA4MymrPYEc4Cu6ToiUM8SDeymgBvRiYRZoS3kE0GeQsG/jeCSLyeRBiF3SkQIMDw33E6wigkh7GygeSZGJgYOhkYGK6CVQMCAAD///5jvhZ0AQAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
