from django.db import models
 #======================================================================
# 
# Encapsulates data for model BloodType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BloodType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class BloodType(Enum):   # A subclass of Enum
	APositive = 'APositive'
	ANegative = 'ANegative'
	BPositive = 'BPositive'
	BNegative = 'BNegative'
	ABPositive = 'ABPositive'
	ABNegative = 'ABNegative'
	OPositive = 'OPositive'
	ONegative = 'ONegative'
