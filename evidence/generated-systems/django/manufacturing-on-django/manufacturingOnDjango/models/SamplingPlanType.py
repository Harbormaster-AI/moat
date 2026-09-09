from django.db import models
 #======================================================================
# 
# Encapsulates data for model SamplingPlanType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SamplingPlanType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SamplingPlanType(Enum):   # A subclass of Enum
	Fixed = 'Fixed'
	Percentage = 'Percentage'
	C0 = 'C0'
