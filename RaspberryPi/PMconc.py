import sys
sys.path.append('/home/pi/Sensly/')
import smbus
##import RPi.GPPIO as GPIO
from Sensors import Sensor, Gas
from bme_combo import *
PM = Sensor('PM',0,0)
PMcmd = 0x04
# Fetch the current PM concentration
PMconc = PM.Get_PMDensity(PMcmd)
print str(PMconc)
