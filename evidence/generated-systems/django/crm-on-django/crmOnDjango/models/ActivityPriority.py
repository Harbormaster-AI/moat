from django.db import models
 #======================================================================
# 
# Encapsulates data for model ActivityPriority
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ActivityPriority Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ActivityPriority(Enum):   # A subclass of Enum
	Low = 'Low'
	Normal = 'Normal'
	High = 'High'
	Urgent = 'Urgent'
