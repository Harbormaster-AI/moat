from django.db import models
 #======================================================================
# 
# Encapsulates data for model FindingSeverity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FindingSeverity Declaration (enumerated type)
#======================================================================
from enum import Enum 
class FindingSeverity(Enum):   # A subclass of Enum
	Low = 'Low'
	Medium = 'Medium'
	High = 'High'
	Critical = 'Critical'
