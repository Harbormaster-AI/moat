from django.db import models
 #======================================================================
# 
# Encapsulates data for model TokenizationStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TokenizationStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TokenizationStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Suspended = 'Suspended'
	Deactivated = 'Deactivated'
