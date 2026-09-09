from django.db import models
 #======================================================================
# 
# Encapsulates data for model OrderStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class OrderStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Submitted = 'Submitted'
	PartiallyFulfilled = 'PartiallyFulfilled'
	Fulfilled = 'Fulfilled'
	Invoiced = 'Invoiced'
	Cancelled = 'Cancelled'
