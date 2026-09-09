from django.db import models
 #======================================================================
# 
# Encapsulates data for model PerilType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerilType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PerilType(Enum):   # A subclass of Enum
	AutoAccident = 'AutoAccident'
	Fire = 'Fire'
	Theft = 'Theft'
	Windstorm = 'Windstorm'
	Flood = 'Flood'
	Hail = 'Hail'
	Earthquake = 'Earthquake'
	Vandalism = 'Vandalism'
	Injury = 'Injury'
	Death = 'Death'
