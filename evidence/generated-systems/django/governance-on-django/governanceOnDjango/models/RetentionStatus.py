from django.db import models
 #======================================================================
# 
# Encapsulates data for model RetentionStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RetentionStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RetentionStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Approved = 'Approved'
	InEffect = 'InEffect'
	Suspended = 'Suspended'
	Retired = 'Retired'
