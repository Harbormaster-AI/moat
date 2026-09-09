from django.db import models
 #======================================================================
# 
# Encapsulates data for model FXDealStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FXDealStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class FXDealStatus(Enum):   # A subclass of Enum
	Booked = 'Booked'
	Cancelled = 'Cancelled'
	Settled = 'Settled'
