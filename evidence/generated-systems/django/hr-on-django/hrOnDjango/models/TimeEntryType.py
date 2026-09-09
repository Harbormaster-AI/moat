from django.db import models
 #======================================================================
# 
# Encapsulates data for model TimeEntryType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimeEntryType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TimeEntryType(Enum):   # A subclass of Enum
	Regular = 'Regular'
	Overtime = 'Overtime'
	Sick = 'Sick'
	Vacation = 'Vacation'
	Unpaid = 'Unpaid'
