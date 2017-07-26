import sys
sys.path.append('/home/pi/Sensly/')
import smbus
##import RPi.GPPIO as GPIO
##from Sensors import Sensor, Gas
from bme_combo import *
sensor = BME280(mode=BME280_OSAMPLE_8)
# Fetch the current temperatire
temperature = sensor.read_temperature()
print str(temperature)
