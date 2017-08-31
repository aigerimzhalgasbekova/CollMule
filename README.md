# CollMule
IN THE DOCUMENTING PROCESS

CollMule is my Master Thesis project which title is "**CollMule: An opportunistic data collection system for IoT-based Indoor Air Quality Monitoring**"

**Supervised by** Prof. Arkady Zaslavsky, Dr Saguna, Dr Karan Mitra and Dr. Prem Prakash Jayaraman

**Abstract:** Opportunistic sensing advance methods of IoT data collection using the mobility of data mules, the proximity of transmitting sensor devices and cost efficiency to decide when, where, how and at what cost collect IoT data and deliver it to a sink. This paper proposes, develops, implements and evaluates the algorithm called CollMule which builds on and extends the 3D kNN approach to discover, negotiate, collect and deliver the sensed data in an energy- and cost-efficient manner. The developed CollMule software prototype uses Android platform to handle indoor air quality data from heterogeneous IoT devices. The CollMule evaluation is based on performing rate, power consumption and CPU usage of single algorithm cycle. The outcomes of these experiments prove the feasibility of CollMule use on mobile smart devices.

**Hardware:** smart device (with Android OS Version 5 and higher), Raspberry Pis and SenslyHats

In this repository you can find Android app (the AppPlatform folder) that: 
1. discovers sensors (Raspberry Pis) that are available in BLE range; 
2. ranks them regarding their cost-efficiency towards the smart device; 
3. connects to and collects data from them.

The app can be downloaded, built and ran on a smart device.

The RaspberryPi folder consists code sources for different purposes and have to be placed in different directories. However, the first thing you need to do is to setup [Raspbian](https://www.raspberrypi.org/downloads/raspbian/) and install software for [SenslyHat](http://www.instructables.com/id/Sensly-Hat-for-the-Raspberry-Pi-Air-Quality-Gas-De/).
After you have to install Bluez (in this [link](https://stackoverflow.com/questions/41351514/leadvertisingmanager1-missing-from-dbus-objectmanager-getmanagedobjects) you will find a nice tutorial)
Then, install [Golang](https://golang.org/doc/install/source)
At the end of the output, you will see that the Go installer asks you to set a binary directory to your path for easy access. The following commands can accomplish this:
```
PATH=$PATH:/home/pi/go/bin

export PATH

export GOROOT=/home pi/go

export GOPATH=/home/pi/gopath
```
Put this commands in the `/etc/profile` file in order to execute them for each session that you start in the future.

Next step is to download the Gatt package source files:

`go get github.com/paypal/gatt`

Now, you have to go to `/home/pi/gopath/src/github.com/paypal/gatt/examples` and change server.go file to the server.go file from RaspberryPi folder of this repository.
After, go to directory `/home/pi/gopath/src/github.com/paypal/gatt/examples/service` and add there iaq.go file.
To your home directory it is ussually `/home/pi` add all Python files from the RaspberryPi folder.

To start a BLE server use the following commands:
```
cd /home/pi/gopath/src/github.com/paypal/gatt

go build examples/server.go

sudo ./server
```
