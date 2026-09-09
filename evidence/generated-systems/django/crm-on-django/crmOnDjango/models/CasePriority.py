from django.db import models
 #======================================================================
# 
# Encapsulates data for model CasePriority
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CasePriority Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CasePriority(Enum):   # A subclass of Enum
	Low = 'Low'
	Medium = 'Medium'
	High = 'High'
	Critical = 'Critical'
