from django.db import models
 #======================================================================
# 
# Encapsulates data for model DeviceConnectivityStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DeviceConnectivityStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DeviceConnectivityStatus(Enum):   # A subclass of Enum
	Connected = 'Connected'
	Disconnected = 'Disconnected'
	Standby = 'Standby'
	Fault = 'Fault'
