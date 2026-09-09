from django.db import models
 #======================================================================
# 
# Encapsulates data for model DispenseStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DispenseStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DispenseStatus(Enum):   # A subclass of Enum
	Preparation = 'Preparation'
	InProgress = 'InProgress'
	Completed = 'Completed'
	Cancelled = 'Cancelled'
