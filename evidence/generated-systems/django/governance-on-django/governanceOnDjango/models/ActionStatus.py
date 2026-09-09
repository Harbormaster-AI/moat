from django.db import models
 #======================================================================
# 
# Encapsulates data for model ActionStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ActionStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ActionStatus(Enum):   # A subclass of Enum
	NotStarted = 'NotStarted'
	InProgress = 'InProgress'
	Deferred = 'Deferred'
	Completed = 'Completed'
	Cancelled = 'Cancelled'
