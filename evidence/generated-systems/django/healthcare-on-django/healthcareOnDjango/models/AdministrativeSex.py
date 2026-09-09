from django.db import models
 #======================================================================
# 
# Encapsulates data for model AdministrativeSex
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdministrativeSex Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AdministrativeSex(Enum):   # A subclass of Enum
	Male = 'Male'
	Female = 'Female'
	Unknown = 'Unknown'
