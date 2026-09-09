from django.db import models
 #======================================================================
# 
# Encapsulates data for model LineItemStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LineItemStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LineItemStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Scheduled = 'Scheduled'
	Running = 'Running'
	Paused = 'Paused'
	Completed = 'Completed'
	Cancelled = 'Cancelled'
