from django.db import models
 #======================================================================
# 
# Encapsulates data for model ContractStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContractStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ContractStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Active = 'Active'
	Expiring = 'Expiring'
	Terminated = 'Terminated'
	Archived = 'Archived'
