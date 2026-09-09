from django.db import models
 #======================================================================
# 
# Encapsulates data for model MaintenanceOrderStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceOrderStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class MaintenanceOrderStatus(Enum):   # A subclass of Enum
	Created = 'Created'
	Approved = 'Approved'
	Scheduled = 'Scheduled'
	InProgress = 'InProgress'
	Completed = 'Completed'
	Cancelled = 'Cancelled'
