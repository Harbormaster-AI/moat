from django.db import models
 #======================================================================
# 
# Encapsulates data for model ForecastMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ForecastMethod Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ForecastMethod(Enum):   # A subclass of Enum
	MovingAverage = 'MovingAverage'
	ExponentialSmoothing = 'ExponentialSmoothing'
	Croston = 'Croston'
	ARIMA = 'ARIMA'
	Manual = 'Manual'
