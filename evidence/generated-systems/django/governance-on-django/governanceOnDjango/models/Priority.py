from django.db import models
 #======================================================================
# 
# Encapsulates data for model Priority
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Priority Declaration (enumerated type)
#======================================================================
from enum import Enum 
class Priority(Enum):   # A subclass of Enum
	Low = 'Low'
	Medium = 'Medium'
	High = 'High'
	Urgent = 'Urgent'
