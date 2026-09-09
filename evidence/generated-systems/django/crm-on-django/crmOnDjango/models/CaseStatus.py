from django.db import models
 #======================================================================
# 
# Encapsulates data for model CaseStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CaseStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CaseStatus(Enum):   # A subclass of Enum
	New = 'New'
	Open = 'Open'
	PendingCustomer = 'PendingCustomer'
	PendingExternal = 'PendingExternal'
	OnHold = 'OnHold'
	Resolved = 'Resolved'
	Closed = 'Closed'
	Reopened = 'Reopened'
