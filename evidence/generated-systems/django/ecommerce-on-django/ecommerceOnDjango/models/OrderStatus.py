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
	Pending = 'Pending'
	Confirmed = 'Confirmed'
	Paid = 'Paid'
	PartiallyShipped = 'PartiallyShipped'
	Shipped = 'Shipped'
	Delivered = 'Delivered'
	Cancelled = 'Cancelled'
	Refunded = 'Refunded'
	PartiallyRefunded = 'PartiallyRefunded'
