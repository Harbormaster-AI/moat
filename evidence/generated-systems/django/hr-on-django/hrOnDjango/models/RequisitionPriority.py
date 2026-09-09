from django.db import models
 #======================================================================
# 
# Encapsulates data for model RequisitionPriority
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RequisitionPriority Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RequisitionPriority(Enum):   # A subclass of Enum
	Low = 'Low'
	Medium = 'Medium'
	High = 'High'
	Critical = 'Critical'
