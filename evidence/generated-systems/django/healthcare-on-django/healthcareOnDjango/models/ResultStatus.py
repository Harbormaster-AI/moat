from django.db import models
 #======================================================================
# 
# Encapsulates data for model ResultStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ResultStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ResultStatus(Enum):   # A subclass of Enum
	Registered = 'Registered'
	Partial = 'Partial'
	Final = 'Final'
	Corrected = 'Corrected'
	Cancelled = 'Cancelled'
