from django.db import models
 #======================================================================
# 
# Encapsulates data for model ReviewStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReviewStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ReviewStatus(Enum):   # A subclass of Enum
	NotStarted = 'NotStarted'
	InProgress = 'InProgress'
	Finalized = 'Finalized'
	Acknowledged = 'Acknowledged'
