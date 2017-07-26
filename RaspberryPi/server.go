// +build

package main

import (
	"fmt"
	"log"
	"strconv"
	"os/exec"
	"strings"
	"time"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
	"github.com/paypal/gatt/examples/service"
)

var (
	acc = false
	ticker = time.NewTicker(time.Millisecond * 1500)
	ticker1 = time.NewTicker(time.Millisecond * 1500)
	con float64 = 0.0
	dis float64 = 0.0
)

func main() {
	d, err := gatt.NewDevice(option.DefaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s", err)
	}


	tick1()
	// Register optional handlers.
	d.Handle(
		gatt.CentralConnected(func(c gatt.Central) { 
			fmt.Println("Connect: ", c.ID())
			ticker1.Stop()
			//fmt.Println(dis)
			ticker1 = time.NewTicker(time.Millisecond * 1500)
			tick()
		}),

		gatt.CentralDisconnected(func(c gatt.Central) { 
			fmt.Println("Disconnect: ", c.ID())
			ticker.Stop()
			//fmt.Println(con) 
			ticker = time.NewTicker(time.Millisecond * 1500)
			tick1()
		}),
	)

	// A mandatory handler for monitoring device state.
	onStateChanged := func(d gatt.Device, s gatt.State) {
		fmt.Printf("State: %s\n", s)
		switch s {
		case gatt.StatePoweredOn:
			// Setup GAP and GATT services for Linux implementation.
			// OS X doesn't export the access of these services.
			d.AddService(service.NewGapService("Gopher")) // no effect on OS X
			d.AddService(service.NewGattService())        // no effect on OS X

			// A simple count service for demo.
			s1 := service.NewCountService()
			d.AddService(s1)

                        // An IAQ service
                        sIAQ := service.NewIAQService()
                        d.AddService(sIAQ)

			// A fake battery service for demo.
			s2 := service.NewBatteryService()
			d.AddService(s2)

			// Advertise device name and service's UUIDs.
			f:=EstimPowerLevel(0)
			for i, eof:=f(); !eof; i, eof = f() {
				check()
				if acc==true {
					d.AdvertiseNameAndServices("IAQMonitoring"+"T"+strconv.Itoa(i), []gatt.UUID{s1.UUID(), sIAQ.UUID(), s2.UUID()})
				} else {
					d.AdvertiseNameAndServices("IAQMonitoring"+"F"+strconv.Itoa(i), []gatt.UUID{s1.UUID(), sIAQ.UUID(), s2.UUID()})
				}
				fmt.Println(i)
				time.Sleep(time.Minute*5)
			}
			

			// Advertise as an OpenBeacon iBeacon
			d.AdvertiseIBeacon(gatt.MustParseUUID("AA6062F098CA42118EC4193EB73CCEB6"), 1, 2, -59)

		default:
		}
	}

	d.Init(onStateChanged)
	select {}
}

func check() bool{
	out, err := exec.Command("sh", "-c", "sudo python /home/pi/PMconc.py").Output()
              if err != nil {
		acc = true
                fmt.Println("error occured %s", err)
            } else {
                stringout := string(out)
                stringout = strings.TrimSpace(stringout)
		if stringout == "0" {
			acc = true
		}
                //fmt.Println(acc)
            }
	    return acc
}

func EstimPowerLevel(n int) func() (int, bool) {
        pwr := 0.0
        res := 0
        i := 90
	return func() (int, bool){
		if i<=n {
			return 0, true
		}
		pwr = (0.38*con + 0.34*dis)/3600
		res = int((pwr/5)*100)
		i-=res
		fmt.Println(pwr)
		return i, false
	}
}

func tick() {
	
	go func() {
		for range ticker.C {
			con+=1.5
			//fmt.Println(con)
		}
	}()
}
func tick1() {
	go func() {
		for range ticker1.C {
			dis+=1.5
			//fmt.Println(dis)
		}
	}()
}