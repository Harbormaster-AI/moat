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
	Case = 'Case'
	Pallet = 'Pallet'
	Dozen = 'Dozen'
	Gram = 'Gram'
	Kilogram = 'Kilogram'
	Pound = 'Pound'
	Ounce = 'Ounce'
	Milliliter = 'Milliliter'
	Liter = 'Liter'
	CubicMeter = 'CubicMeter'
	Meter = 'Meter'
	Foot = 'Foot'
	SquareMeter = 'SquareMeter'
