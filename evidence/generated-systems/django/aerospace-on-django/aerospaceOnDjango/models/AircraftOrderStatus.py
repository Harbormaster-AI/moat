from django.db import models
 #======================================================================
# 
# Encapsulates data for model AircraftOrderStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftOrderStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AircraftOrderStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Committed = 'Committed'
	InProduction = 'InProduction'
	Delivered = 'Delivered'
	Cancelled = 'Cancelled'
