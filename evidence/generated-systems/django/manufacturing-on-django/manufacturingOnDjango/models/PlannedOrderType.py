from django.db import models
 #======================================================================
# 
# Encapsulates data for model PlannedOrderType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlannedOrderType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PlannedOrderType(Enum):   # A subclass of Enum
	WorkOrder = 'WorkOrder'
	PurchaseRequisition = 'PurchaseRequisition'
	TransferOrder = 'TransferOrder'
