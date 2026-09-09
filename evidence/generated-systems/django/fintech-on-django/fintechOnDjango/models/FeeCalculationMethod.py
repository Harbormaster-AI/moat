from django.db import models
 #======================================================================
# 
# Encapsulates data for model FeeCalculationMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeeCalculationMethod Declaration (enumerated type)
#======================================================================
from enum import Enum 
class FeeCalculationMethod(Enum):   # A subclass of Enum
	PerTransaction = 'PerTransaction'
	PerMonth = 'PerMonth'
	PerAnnum = 'PerAnnum'
	Slab = 'Slab'
	Tiered = 'Tiered'
