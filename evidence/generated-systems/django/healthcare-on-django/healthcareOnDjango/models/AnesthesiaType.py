from django.db import models
 #======================================================================
# 
# Encapsulates data for model AnesthesiaType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AnesthesiaType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AnesthesiaType(Enum):   # A subclass of Enum
	None = 'None'
	Local = 'Local'
	Regional = 'Regional'
	General = 'General'
	Sedation = 'Sedation'
