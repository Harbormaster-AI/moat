from django.db import models
 #======================================================================
# 
# Encapsulates data for model StockStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StockStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class StockStatus(Enum):   # A subclass of Enum
	Available = 'Available'
	Reserved = 'Reserved'
	Damaged = 'Damaged'
	Hold = 'Hold'
	Quarantined = 'Quarantined'
	InTransit = 'InTransit'
	PendingInspection = 'PendingInspection'
