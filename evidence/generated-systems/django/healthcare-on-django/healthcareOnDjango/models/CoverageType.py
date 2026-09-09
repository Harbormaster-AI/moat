from django.db import models
 #======================================================================
# 
# Encapsulates data for model CoverageType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CoverageType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CoverageType(Enum):   # A subclass of Enum
	Medical = 'Medical'
	Pharmacy = 'Pharmacy'
	Dental = 'Dental'
	Vision = 'Vision'
	BehavioralHealth = 'BehavioralHealth'
