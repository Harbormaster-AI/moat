from django.db import models
 #======================================================================
# 
# Encapsulates data for model ScheduleType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ScheduleType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ScheduleType(Enum):   # A subclass of Enum
	Fixed = 'Fixed'
	Flexible = 'Flexible'
	Rotating = 'Rotating'
