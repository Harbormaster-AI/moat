from django.db import models
 #======================================================================
# 
# Encapsulates data for model WorkOrderStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkOrderStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class WorkOrderStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	InProgress = 'InProgress'
	AwaitingParts = 'AwaitingParts'
	Closed = 'Closed'
	Deferred = 'Deferred'
