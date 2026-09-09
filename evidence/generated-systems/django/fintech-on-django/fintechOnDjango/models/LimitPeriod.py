from django.db import models
 #======================================================================
# 
# Encapsulates data for model LimitPeriod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LimitPeriod Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LimitPeriod(Enum):   # A subclass of Enum
	None = 'None'
	Day = 'Day'
	Week = 'Week'
	Month = 'Month'
	Year = 'Year'
