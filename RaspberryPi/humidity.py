import sys
sys.path.append('/home/pi/Sensly/')
import smbus
from bme_combo import *
sensor = BME280(mode=BME280_OSAMPLE_8)
# Fetch the current humidity
humidity = sensor.read_humidity()
print str(humidity)
