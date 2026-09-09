from django.db import models
 #======================================================================
# 
# Encapsulates data for model ClaimStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ClaimStatus(Enum):   # A subclass of Enum
	Submitted = 'Submitted'
	InProcess = 'InProcess'
	Paid = 'Paid'
	Denied = 'Denied'
	Adjusted = 'Adjusted'
	Void = 'Void'
