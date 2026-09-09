from django.db import models
 #======================================================================
# 
# Encapsulates data for model CountStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CountStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CountStatus(Enum):   # A subclass of Enum
	Planned = 'Planned'
	InProgress = 'InProgress'
	Completed = 'Completed'
	Posted = 'Posted'
	Cancelled = 'Cancelled'
