package service

import(
 "fmt"
 "log"
 "os/exec"
 "strings"

 "github.com/paypal/gatt"
)

// Following function puts all fetched data from sensors into characteristis that are advertised via BLE

func NewIAQService() *gatt.Service{
        s := gatt.NewService(gatt.MustParseUUID("19fc95c0-c111-11e3-9904-0002a5d5c51b"))
        s.AddCharacteristic(gatt.MustParseUUID("21fac9e0-c111-11e3-9904-0002a5d5c51b")).HandleReadFunc(
          func(rsp gatt.ResponseWriter, req *gatt.ReadRequest){
            out, err := exec.Command("sh", "-c", "sudo python /home/pi/temperature.py").Output()
              if err != nil {
                fmt.Fprintf(rsp, "error occured %s", err)
                log.Println("Wrote: error %s", err)
            } else {
                stringout := string(out)
                stringout = strings.TrimSpace(stringout)
                fmt.Fprintf(rsp, stringout)
                log.Println("Wrote:", stringout)
            }
          })

        s.AddCharacteristic(gatt.MustParseUUID("31fac9e0-c111-11e3-9904-0002a5d5c51b")).HandleReadFunc(
          func(rsp gatt.ResponseWriter, req *gatt.ReadRequest){
            out, err := exec.Command("sh", "-c", "sudo python /home/pi/humidity.py").Output()
              if err != nil {
                fmt.Fprintf(rsp, "error occured %s", err)
                log.Println("Wrote: error %s", err)
            } else {
                stringout := string(out)
                stringout = strings.TrimSpace(stringout)
                fmt.Fprintf(rsp, stringout)
                log.Println("Wrote:", stringout)
            }
          })

          s.AddCharacteristic(gatt.MustParseUUID("41fac9e0-c111-11e3-9904-0002a5d5c51b")).HandleReadFunc(
          func(rsp gatt.ResponseWriter, req *gatt.ReadRequest){
            out, err := exec.Command("sh", "-c", "sudo python /home/pi/PMconc.py").Output()
              if err != nil {
                fmt.Fprintf(rsp, "error occured %s", err)
                log.Println("Wrote: error %s", err)
            } else {
                stringout := string(out)
                stringout = strings.TrimSpace(stringout)
                fmt.Fprintf(rsp, stringout)
                log.Println("Wrote:", stringout)
            }
          })

          return s
}
