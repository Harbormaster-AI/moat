from django.db import models
 #======================================================================
# 
# Encapsulates data for model TaskStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TaskStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TaskStatus(Enum):   # A subclass of Enum
	Requested = 'Requested'
	Accepted = 'Accepted'
	InProgress = 'InProgress'
	Completed = 'Completed'
	Cancelled = 'Cancelled'
	Failed = 'Failed'
