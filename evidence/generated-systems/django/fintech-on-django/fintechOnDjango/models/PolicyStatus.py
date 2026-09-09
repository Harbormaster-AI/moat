from django.db import models
 #======================================================================
# 
# Encapsulates data for model PolicyStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PolicyStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Active = 'Active'
	Retired = 'Retired'
