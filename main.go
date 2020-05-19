package main

import (
	"fmt"
	"time"

	"github.com/vishvananda/netlink"
)

func main() {
	oldData := make(map[string]netlink.LinkStatistics)
	for {
		allLinks, err := netlink.LinkList()
		if err != nil {
			fmt.Println("get link list error")
			fmt.Println(err)
			return
		}
		for _, oneLink := range allLinks {
			oneAttr := oneLink.Attrs()
			if oneLink.Type() == "device" && oneAttr.OperState.String() == "up" {
				oldLinkStatistics, ok := oldData[oneAttr.Name]
				newLinkSatistics := oneAttr.Statistics
				if ok == true {
					fmt.Printf("%s RXBytes %5d KB/s      TXBytes %5d KB/s\n", oneAttr.Name, (newLinkSatistics.RxBytes-oldLinkStatistics.RxBytes)/1000, (newLinkSatistics.TxBytes-oldLinkStatistics.TxBytes)/1000)
				}
				oldData[oneAttr.Name] = *newLinkSatistics
			}
		}
		time.Sleep(1 * 1000 * 1000 * 1000)
	}
}
