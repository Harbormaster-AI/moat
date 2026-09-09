from django.db import models
 #======================================================================
# 
# Encapsulates data for model PayrollStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PayrollStatus(Enum):   # A subclass of Enum
	Scheduled = 'Scheduled'
	InProgress = 'InProgress'
	Completed = 'Completed'
	Reversed = 'Reversed'
