from django.db import models
 #======================================================================
# 
# Encapsulates data for model PaymentOrderStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentOrderStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PaymentOrderStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Submitted = 'Submitted'
	Processing = 'Processing'
	Completed = 'Completed'
	Cancelled = 'Cancelled'
	Failed = 'Failed'
