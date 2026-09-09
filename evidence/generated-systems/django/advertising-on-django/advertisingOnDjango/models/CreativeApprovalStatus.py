from django.db import models
 #======================================================================
# 
# Encapsulates data for model CreativeApprovalStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeApprovalStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CreativeApprovalStatus(Enum):   # A subclass of Enum
	Pending = 'Pending'
	Approved = 'Approved'
	Rejected = 'Rejected'
