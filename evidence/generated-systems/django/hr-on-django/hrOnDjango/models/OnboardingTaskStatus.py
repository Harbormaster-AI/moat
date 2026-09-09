from django.db import models
 #======================================================================
# 
# Encapsulates data for model OnboardingTaskStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OnboardingTaskStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class OnboardingTaskStatus(Enum):   # A subclass of Enum
	NotStarted = 'NotStarted'
	InProgress = 'InProgress'
	Blocked = 'Blocked'
	Completed = 'Completed'
