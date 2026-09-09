from django.db import models
 #======================================================================
# 
# Encapsulates data for model ActivityType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ActivityType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ActivityType(Enum):   # A subclass of Enum
	Task = 'Task'
	Call = 'Call'
	Meeting = 'Meeting'
	Demo = 'Demo'
	FollowUp = 'FollowUp'
