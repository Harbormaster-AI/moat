from django.db import models
 #======================================================================
# 
# Encapsulates data for model AdjustmentStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdjustmentStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AdjustmentStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Approved = 'Approved'
	Posted = 'Posted'
	Cancelled = 'Cancelled'
