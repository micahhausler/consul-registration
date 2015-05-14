package main

import (
	"flag"
	"fmt"
	"github.com/micahhausler/consul-registration/open"
	"github.com/micahhausler/consul-registration/post"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

type strslice []string

func (i *strslice) String() string {
	return fmt.Sprintf("%d", *i)
}

func (i *strslice) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var serviceTags strslice

func main() {
	containerNamePtr := flag.String("container", "", "The container name to watch")
	consulAddrPtr := flag.String("consul", "consul.service.consul", "The address or IP for consul")
	serviceNamePtr := flag.String("name", "", "The service name for consul")
	serviceIdPtr := flag.String("id", "", "The service ID for consul")
	flag.Var(&serviceTags, "tag", "A tag to be applied to the service. Repeat option for multiple tags")
	sleepPtr := flag.Int("sleep", 30, "How long to wait between checking in with consul.")

	checkTtlPtr := flag.Int("ttl", 45, "TTL for the service. Make this larget than -sleep")
	checkHttpPtr := flag.String("http", "", "See https://www.consul.io/docs/agent/checks.html")
	checkIntervalPtr := flag.Int("interval", 0, "Interval for consul's HTTP check")
	checkScriptPtr := flag.String("script", "", "Script on consul server to execute")

	oncePtr := flag.Bool("once", false, "Only register the service once, then exit")

	flag.Parse()

	containerAddress := open.FindAddress(*containerNamePtr)
	fmt.Printf("container address: \"%s\"\n", containerAddress)

	check := post.Check{
		Ttl:      *checkTtlPtr,
		Http:     *checkHttpPtr,
		Interval: *checkIntervalPtr,
		Script:   *checkScriptPtr,
	}

	registration := post.Registration{
		Id:      *serviceIdPtr,
		Name:    *serviceNamePtr,
		Tags:    serviceTags,
		Address: containerAddress,
		Check:   &check,
	}

	post.RegisterService(&registration, *consulAddrPtr)
	if *oncePtr {
		os.Exit(0)
	}
	for {
		sleepTime := time.Duration(*sleepPtr) * time.Second
		fmt.Printf("sleeping %v seconds\n", sleepTime)
		time.Sleep(sleepTime)
		post.RegisterService(&registration, *consulAddrPtr)
	}
}
