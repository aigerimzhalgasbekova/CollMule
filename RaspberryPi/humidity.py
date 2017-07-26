import sys
sys.path.append('/home/pi/Sensly/')
import smbus
##import RPi.GPPIO as GPIO
##from Sensors import Sensor, Gas
from bme_combo import *
sensor = BME280(mode=BME280_OSAMPLE_8)
# Fetch the current humidity
humidity = sensor.read_humidity()
print str(humidity)
