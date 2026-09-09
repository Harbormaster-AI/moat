from django.db import models
 #======================================================================
# 
# Encapsulates data for model SubscriptionStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubscriptionStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SubscriptionStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Paused = 'Paused'
	Cancelled = 'Cancelled'
	Expired = 'Expired'
