from django.db import models
 #======================================================================
# 
# Encapsulates data for model GoalStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GoalStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class GoalStatus(Enum):   # A subclass of Enum
	NotStarted = 'NotStarted'
	InProgress = 'InProgress'
	Completed = 'Completed'
	Deferred = 'Deferred'
	Cancelled = 'Cancelled'
