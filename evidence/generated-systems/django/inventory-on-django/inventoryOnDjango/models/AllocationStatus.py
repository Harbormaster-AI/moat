from django.db import models
 #======================================================================
# 
# Encapsulates data for model AllocationStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AllocationStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AllocationStatus(Enum):   # A subclass of Enum
	Proposed = 'Proposed'
	Confirmed = 'Confirmed'
	Picked = 'Picked'
	Short = 'Short'
	Cancelled = 'Cancelled'
