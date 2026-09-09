from django.db import models
 #======================================================================
# 
# Encapsulates data for model SupplierApprovalStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SupplierApprovalStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SupplierApprovalStatus(Enum):   # A subclass of Enum
	Applied = 'Applied'
	Approved = 'Approved'
	OnHold = 'OnHold'
	Suspended = 'Suspended'
