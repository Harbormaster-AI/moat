from django.db import models
 #======================================================================
# 
# Encapsulates data for model AuthorizationStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuthorizationStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AuthorizationStatus(Enum):   # A subclass of Enum
	Requested = 'Requested'
	PendingReview = 'PendingReview'
	Approved = 'Approved'
	Denied = 'Denied'
	Expired = 'Expired'
