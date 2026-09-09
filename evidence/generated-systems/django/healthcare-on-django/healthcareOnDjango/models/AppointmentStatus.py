from django.db import models
 #======================================================================
# 
# Encapsulates data for model AppointmentStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AppointmentStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AppointmentStatus(Enum):   # A subclass of Enum
	Proposed = 'Proposed'
	Booked = 'Booked'
	Arrived = 'Arrived'
	Fulfilled = 'Fulfilled'
	Cancelled = 'Cancelled'
	NoShow = 'NoShow'
	EnteredInError = 'EnteredInError'
