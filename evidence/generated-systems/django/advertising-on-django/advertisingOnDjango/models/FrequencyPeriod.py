from django.db import models
 #======================================================================
# 
# Encapsulates data for model FrequencyPeriod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FrequencyPeriod Declaration (enumerated type)
#======================================================================
from enum import Enum 
class FrequencyPeriod(Enum):   # A subclass of Enum
	Hour = 'Hour'
	Day = 'Day'
	Week = 'Week'
	Month = 'Month'
	Lifetime = 'Lifetime'
