from django.db import models
 #======================================================================
# 
# Encapsulates data for model FraudSignalType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FraudSignalType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class FraudSignalType(Enum):   # A subclass of Enum
	Velocity = 'Velocity'
	GeolocationMismatch = 'GeolocationMismatch'
	AmountOutlier = 'AmountOutlier'
	DeviceFingerprint = 'DeviceFingerprint'
	BehavioralChange = 'BehavioralChange'
