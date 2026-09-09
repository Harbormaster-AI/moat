from django.db import models
 #======================================================================
# 
# Encapsulates data for model InspectionStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InspectionStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	InProgress = 'InProgress'
	Completed = 'Completed'
	Accepted = 'Accepted'
	Rejected = 'Rejected'
