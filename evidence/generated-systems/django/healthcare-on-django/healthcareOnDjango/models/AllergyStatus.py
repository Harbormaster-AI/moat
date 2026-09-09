from django.db import models
 #======================================================================
# 
# Encapsulates data for model AllergyStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AllergyStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AllergyStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Inactive = 'Inactive'
	Resolved = 'Resolved'
	EnteredInError = 'EnteredInError'
