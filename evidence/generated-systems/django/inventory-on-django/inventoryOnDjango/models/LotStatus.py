from django.db import models
 #======================================================================
# 
# Encapsulates data for model LotStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LotStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LotStatus(Enum):   # A subclass of Enum
	Released = 'Released'
	Quarantined = 'Quarantined'
	Expired = 'Expired'
	Blocked = 'Blocked'
	PendingTest = 'PendingTest'
