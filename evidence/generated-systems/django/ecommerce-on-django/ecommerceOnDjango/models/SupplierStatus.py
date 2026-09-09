from django.db import models
 #======================================================================
# 
# Encapsulates data for model SupplierStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SupplierStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SupplierStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Inactive = 'Inactive'
	Onboarding = 'Onboarding'
	Suspended = 'Suspended'
