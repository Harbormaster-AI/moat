from django.db import models
 #======================================================================
# 
# Encapsulates data for model LegalHoldStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LegalHoldStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LegalHoldStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Released = 'Released'
	Superseded = 'Superseded'
