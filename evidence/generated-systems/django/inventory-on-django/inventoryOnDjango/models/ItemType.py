from django.db import models
 #======================================================================
# 
# Encapsulates data for model ItemType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ItemType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ItemType(Enum):   # A subclass of Enum
	FinishedGood = 'FinishedGood'
	Component = 'Component'
	RawMaterial = 'RawMaterial'
	Packaging = 'Packaging'
	SparePart = 'SparePart'
	Consumable = 'Consumable'
