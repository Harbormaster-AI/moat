from django.db import models
 #======================================================================
# 
# Encapsulates data for model InsuredObjectType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsuredObjectType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InsuredObjectType(Enum):   # A subclass of Enum
	Vehicle = 'Vehicle'
	Property = 'Property'
	Person = 'Person'
	Equipment = 'Equipment'
	LiabilityExposure = 'LiabilityExposure'
