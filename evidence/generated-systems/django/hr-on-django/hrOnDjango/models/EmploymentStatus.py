from django.db import models
 #======================================================================
# 
# Encapsulates data for model EmploymentStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmploymentStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class EmploymentStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	OnLeave = 'OnLeave'
	Suspended = 'Suspended'
	Terminated = 'Terminated'
