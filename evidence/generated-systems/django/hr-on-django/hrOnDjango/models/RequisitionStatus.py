from django.db import models
 #======================================================================
# 
# Encapsulates data for model RequisitionStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RequisitionStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RequisitionStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Open = 'Open'
	OnHold = 'OnHold'
	Closed = 'Closed'
	Cancelled = 'Cancelled'
