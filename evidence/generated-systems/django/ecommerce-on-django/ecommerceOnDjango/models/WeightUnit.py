from django.db import models
 #======================================================================
# 
# Encapsulates data for model WeightUnit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WeightUnit Declaration (enumerated type)
#======================================================================
from enum import Enum 
class WeightUnit(Enum):   # A subclass of Enum
	Gram = 'Gram'
	Kilogram = 'Kilogram'
	Ounce = 'Ounce'
	Pound = 'Pound'
