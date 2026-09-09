from django.db import models
 #======================================================================
# 
# Encapsulates data for model UnitOfMeasure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UnitOfMeasure Declaration (enumerated type)
#======================================================================
from enum import Enum 
class UnitOfMeasure(Enum):   # A subclass of Enum
	Each = 'Each'
	Kilogram = 'Kilogram'
	Gram = 'Gram'
	Pound = 'Pound'
	Liter = 'Liter'
	Meter = 'Meter'
	Centimeter = 'Centimeter'
	Millimeter = 'Millimeter'
	Hour = 'Hour'
	Minute = 'Minute'
	Box = 'Box'
	Pallet = 'Pallet'
