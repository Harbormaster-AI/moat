from django.db import models
 #======================================================================
# 
# Encapsulates data for model SellerStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SellerStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SellerStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Inactive = 'Inactive'
	Suspended = 'Suspended'
