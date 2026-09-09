from django.db import models
 #======================================================================
# 
# Encapsulates data for model CouponStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CouponStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CouponStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Expired = 'Expired'
	Disabled = 'Disabled'
	Exhausted = 'Exhausted'
