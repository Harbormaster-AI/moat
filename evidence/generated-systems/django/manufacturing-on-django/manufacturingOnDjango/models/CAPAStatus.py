from django.db import models
 #======================================================================
# 
# Encapsulates data for model CAPAStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CAPAStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CAPAStatus(Enum):   # A subclass of Enum
	Proposed = 'Proposed'
	Approved = 'Approved'
	Implemented = 'Implemented'
	Verified = 'Verified'
	Closed = 'Closed'
