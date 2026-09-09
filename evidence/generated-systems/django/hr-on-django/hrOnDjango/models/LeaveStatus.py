from django.db import models
 #======================================================================
# 
# Encapsulates data for model LeaveStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeaveStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LeaveStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Submitted = 'Submitted'
	Approved = 'Approved'
	Rejected = 'Rejected'
	Cancelled = 'Cancelled'
	Taken = 'Taken'
