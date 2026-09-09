from django.db import models
 #======================================================================
# 
# Encapsulates data for model WarrantyType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WarrantyType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class WarrantyType(Enum):   # A subclass of Enum
	Basic = 'Basic'
	Powerplant = 'Powerplant'
	Avionics = 'Avionics'
	Corrosion = 'Corrosion'
