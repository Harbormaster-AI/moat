from django.db import models
 #======================================================================
# 
# Encapsulates data for model RefundStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RefundStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RefundStatus(Enum):   # A subclass of Enum
	Requested = 'Requested'
	Approved = 'Approved'
	Declined = 'Declined'
	Processed = 'Processed'
