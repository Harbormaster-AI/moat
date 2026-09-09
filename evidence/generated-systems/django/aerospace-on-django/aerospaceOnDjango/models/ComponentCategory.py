from django.db import models
 #======================================================================
# 
# Encapsulates data for model ComponentCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComponentCategory Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ComponentCategory(Enum):   # A subclass of Enum
	Structure = 'Structure'
	System = 'System'
	Avionics = 'Avionics'
	Interior = 'Interior'
	LandingGear = 'LandingGear'
	Powerplant = 'Powerplant'
	Consumable = 'Consumable'
