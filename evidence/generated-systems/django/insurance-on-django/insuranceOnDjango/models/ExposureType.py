from django.db import models
 #======================================================================
# 
# Encapsulates data for model ExposureType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExposureType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ExposureType(Enum):   # A subclass of Enum
	BodilyInjury = 'BodilyInjury'
	PropertyDamage = 'PropertyDamage'
	Medical = 'Medical'
	UninsuredMotorist = 'UninsuredMotorist'
	PersonalInjuryProtection = 'PersonalInjuryProtection'
	DwellingDamage = 'DwellingDamage'
	ContentsDamage = 'ContentsDamage'
	BusinessIncome = 'BusinessIncome'
