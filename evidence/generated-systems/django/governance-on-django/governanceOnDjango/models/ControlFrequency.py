from django.db import models
 #======================================================================
# 
# Encapsulates data for model ControlFrequency
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ControlFrequency Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ControlFrequency(Enum):   # A subclass of Enum
	Continuous = 'Continuous'
	Daily = 'Daily'
	Weekly = 'Weekly'
	Monthly = 'Monthly'
	Quarterly = 'Quarterly'
	Annually = 'Annually'
	AdHoc = 'AdHoc'
