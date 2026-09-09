from django.db import models
 #======================================================================
# 
# Encapsulates data for model TimeGranularity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimeGranularity Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TimeGranularity(Enum):   # A subclass of Enum
	Minute = 'Minute'
	Hour = 'Hour'
	Day = 'Day'
	Week = 'Week'
	Month = 'Month'
	Quarter = 'Quarter'
	Year = 'Year'
