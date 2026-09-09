from django.db import models
 #======================================================================
# 
# Encapsulates data for model Disposition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Disposition Declaration (enumerated type)
#======================================================================
from enum import Enum 
class Disposition(Enum):   # A subclass of Enum
	Release = 'Release'
	Scrap = 'Scrap'
	ReturnToVendor = 'ReturnToVendor'
	Rework = 'Rework'
