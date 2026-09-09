from django.db import models
 #======================================================================
# 
# Encapsulates data for model DayOfWeek
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DayOfWeek Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DayOfWeek(Enum):   # A subclass of Enum
	Monday = 'Monday'
	Tuesday = 'Tuesday'
	Wednesday = 'Wednesday'
	Thursday = 'Thursday'
	Friday = 'Friday'
	Saturday = 'Saturday'
	Sunday = 'Sunday'
