from django.db import models
 #======================================================================
# 
# Encapsulates data for model DeviceType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DeviceType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DeviceType(Enum):   # A subclass of Enum
	Pacemaker = 'Pacemaker'
	InsulinPump = 'InsulinPump'
	BloodPressureMonitor = 'BloodPressureMonitor'
	GlucoseMeter = 'GlucoseMeter'
	PulseOximeter = 'PulseOximeter'
	Ventilator = 'Ventilator'
	InfusionPump = 'InfusionPump'
	WearableTracker = 'WearableTracker'
