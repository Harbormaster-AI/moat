from django.db import models
 #======================================================================
# 
# Encapsulates data for model OrderLineStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderLineStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class OrderLineStatus(Enum):   # A subclass of Enum
	Pending = 'Pending'
	Fulfilled = 'Fulfilled'
	Cancelled = 'Cancelled'
	Backordered = 'Backordered'
	Returned = 'Returned'
