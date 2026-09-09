from django.db import models
 #======================================================================
# 
# Encapsulates data for model ScheduleStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ScheduleStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ScheduleStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Published = 'Published'
	Revised = 'Revised'
	Closed = 'Closed'
