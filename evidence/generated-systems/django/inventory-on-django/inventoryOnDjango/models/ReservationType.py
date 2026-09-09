from django.db import models
 #======================================================================
# 
# Encapsulates data for model ReservationType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReservationType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ReservationType(Enum):   # A subclass of Enum
	SalesOrder = 'SalesOrder'
	WorkOrder = 'WorkOrder'
	TransferOrder = 'TransferOrder'
	ServiceOrder = 'ServiceOrder'
	Other = 'Other'
