from django.db import models
 #======================================================================
# 
# Encapsulates data for model ShipmentStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShipmentStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ShipmentStatus(Enum):   # A subclass of Enum
	Pending = 'Pending'
	Packed = 'Packed'
	Shipped = 'Shipped'
	InTransit = 'InTransit'
	Delivered = 'Delivered'
	Delayed = 'Delayed'
	Returned = 'Returned'
	Cancelled = 'Cancelled'
