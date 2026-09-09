from django.db import models
 #======================================================================
# 
# Encapsulates data for model ApprovalStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ApprovalStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ApprovalStatus(Enum):   # A subclass of Enum
	Pending = 'Pending'
	Approved = 'Approved'
	Rejected = 'Rejected'
	Cancelled = 'Cancelled'
