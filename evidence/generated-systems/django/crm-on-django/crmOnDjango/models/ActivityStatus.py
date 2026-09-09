from django.db import models
 #======================================================================
# 
# Encapsulates data for model ActivityStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ActivityStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ActivityStatus(Enum):   # A subclass of Enum
	NotStarted = 'NotStarted'
	InProgress = 'InProgress'
	Completed = 'Completed'
	Deferred = 'Deferred'
	Cancelled = 'Cancelled'
