from django.db import models
 #======================================================================
# 
# Encapsulates data for model LandingGearType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LandingGearType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LandingGearType(Enum):   # A subclass of Enum
	Tricycle = 'Tricycle'
	Tandem = 'Tandem'
	Taildragger = 'Taildragger'
	Skid = 'Skid'
	Floats = 'Floats'
	Retractable = 'Retractable'
