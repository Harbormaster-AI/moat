from django.db import models
 #======================================================================
# 
# Encapsulates data for model OrderStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class OrderStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Active = 'Active'
	OnHold = 'OnHold'
	Completed = 'Completed'
	Cancelled = 'Cancelled'
