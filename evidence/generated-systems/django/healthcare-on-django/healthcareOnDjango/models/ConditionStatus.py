from django.db import models
 #======================================================================
# 
# Encapsulates data for model ConditionStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConditionStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ConditionStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Recurrence = 'Recurrence'
	Relapse = 'Relapse'
	Remission = 'Remission'
	Resolved = 'Resolved'
