from django.db import models
 #======================================================================
# 
# Encapsulates data for model IssueStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class IssueStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class IssueStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	Investigating = 'Investigating'
	RemediationPlanned = 'RemediationPlanned'
	RemediationInProgress = 'RemediationInProgress'
	Verified = 'Verified'
	Closed = 'Closed'
