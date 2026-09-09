from django.db import models
 #======================================================================
# 
# Encapsulates data for model InterestRateType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InterestRateType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InterestRateType(Enum):   # A subclass of Enum
	Fixed = 'Fixed'
	Variable = 'Variable'
