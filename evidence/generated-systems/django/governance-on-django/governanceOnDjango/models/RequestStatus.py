from django.db import models
 #======================================================================
# 
# Encapsulates data for model RequestStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RequestStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RequestStatus(Enum):   # A subclass of Enum
	Received = 'Received'
	InValidation = 'InValidation'
	InProgress = 'InProgress'
	OnHold = 'OnHold'
	Fulfilled = 'Fulfilled'
	Rejected = 'Rejected'
