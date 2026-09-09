from django.db import models
 #======================================================================
# 
# Encapsulates data for model ReservationStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReservationStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ReservationStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Confirmed = 'Confirmed'
	Released = 'Released'
	Fulfilled = 'Fulfilled'
	Cancelled = 'Cancelled'
	Expired = 'Expired'
