package b

import "sync"
import sub2 "git.oschina.net/jscode/list-interfaces/testdata/b/sub"
import . "git.oschina.net/jscode/list-interfaces/testdata/b/suba"

type B struct {}

type IA interface  {
	Add(a sub2.SubSA, locker sync.Locker, b B, subsa1 SubSa1)
}
