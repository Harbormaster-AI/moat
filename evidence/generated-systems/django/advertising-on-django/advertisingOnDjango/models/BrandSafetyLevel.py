from django.db import models
 #======================================================================
# 
# Encapsulates data for model BrandSafetyLevel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BrandSafetyLevel Declaration (enumerated type)
#======================================================================
from enum import Enum 
class BrandSafetyLevel(Enum):   # A subclass of Enum
	None = 'None'
	Moderate = 'Moderate'
	Strict = 'Strict'
