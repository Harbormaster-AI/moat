from django.db import models
 #======================================================================
# 
# Encapsulates data for model TimeUnit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimeUnit Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TimeUnit(Enum):   # A subclass of Enum
	Second = 'Second'
	Minute = 'Minute'
	Hour = 'Hour'
	Day = 'Day'
