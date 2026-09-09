from django.db import models
 #======================================================================
# 
# Encapsulates data for model CycleStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CycleStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CycleStatus(Enum):   # A subclass of Enum
	Planned = 'Planned'
	Open = 'Open'
	Closed = 'Closed'
