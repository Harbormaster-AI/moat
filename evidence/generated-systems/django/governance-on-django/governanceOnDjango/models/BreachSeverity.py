from django.db import models
 #======================================================================
# 
# Encapsulates data for model BreachSeverity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BreachSeverity Declaration (enumerated type)
#======================================================================
from enum import Enum 
class BreachSeverity(Enum):   # A subclass of Enum
	Low = 'Low'
	Medium = 'Medium'
	High = 'High'
	Critical = 'Critical'
