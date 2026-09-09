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
	Planned = 'Planned'
	Released = 'Released'
	InProcess = 'InProcess'
	Hold = 'Hold'
	Completed = 'Completed'
	Closed = 'Closed'
	Cancelled = 'Cancelled'
