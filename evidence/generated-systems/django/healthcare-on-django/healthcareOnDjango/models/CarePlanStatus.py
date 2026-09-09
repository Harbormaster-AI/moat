from django.db import models
 #======================================================================
# 
# Encapsulates data for model CarePlanStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CarePlanStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CarePlanStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Active = 'Active'
	Suspended = 'Suspended'
	Completed = 'Completed'
	Cancelled = 'Cancelled'
