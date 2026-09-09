from django.db import models
 #======================================================================
# 
# Encapsulates data for model TrainingStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrainingStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TrainingStatus(Enum):   # A subclass of Enum
	Enrolled = 'Enrolled'
	InProgress = 'InProgress'
	Completed = 'Completed'
	Failed = 'Failed'
	Cancelled = 'Cancelled'
