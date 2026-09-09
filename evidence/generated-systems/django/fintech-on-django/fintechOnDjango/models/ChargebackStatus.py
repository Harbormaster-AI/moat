from django.db import models
 #======================================================================
# 
# Encapsulates data for model ChargebackStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ChargebackStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ChargebackStatus(Enum):   # A subclass of Enum
	Pending = 'Pending'
	Accepted = 'Accepted'
	Reversed = 'Reversed'
	Lost = 'Lost'
