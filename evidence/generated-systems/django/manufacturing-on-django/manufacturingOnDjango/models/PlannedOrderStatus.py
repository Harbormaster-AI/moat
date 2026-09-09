from django.db import models
 #======================================================================
# 
# Encapsulates data for model PlannedOrderStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlannedOrderStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PlannedOrderStatus(Enum):   # A subclass of Enum
	Planned = 'Planned'
	Firmed = 'Firmed'
	Released = 'Released'
	Cancelled = 'Cancelled'
