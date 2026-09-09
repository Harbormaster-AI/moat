from django.db import models
 #======================================================================
# 
# Encapsulates data for model PurchaseOrderStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PurchaseOrderStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PurchaseOrderStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Submitted = 'Submitted'
	Acknowledged = 'Acknowledged'
	PartiallyReceived = 'PartiallyReceived'
	Received = 'Received'
	Closed = 'Closed'
	Cancelled = 'Cancelled'
