from django.db import models
 #======================================================================
# 
# Encapsulates data for model PayoutStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayoutStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PayoutStatus(Enum):   # A subclass of Enum
	Pending = 'Pending'
	Scheduled = 'Scheduled'
	Paid = 'Paid'
	Failed = 'Failed'
	Cancelled = 'Cancelled'
