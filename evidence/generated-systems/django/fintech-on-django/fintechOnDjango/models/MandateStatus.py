from django.db import models
 #======================================================================
# 
# Encapsulates data for model MandateStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MandateStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class MandateStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Suspended = 'Suspended'
	Cancelled = 'Cancelled'
	Expired = 'Expired'
