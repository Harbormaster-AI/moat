from django.db import models
 #======================================================================
# 
# Encapsulates data for model PaymentStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PaymentStatus(Enum):   # A subclass of Enum
	Pending = 'Pending'
	Settled = 'Settled'
	Failed = 'Failed'
	Refunded = 'Refunded'
	Reversed = 'Reversed'
