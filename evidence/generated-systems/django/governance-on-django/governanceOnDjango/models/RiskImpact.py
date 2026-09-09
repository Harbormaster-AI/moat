from django.db import models
 #======================================================================
# 
# Encapsulates data for model RiskImpact
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RiskImpact Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RiskImpact(Enum):   # A subclass of Enum
	Insignificant = 'Insignificant'
	Minor = 'Minor'
	Moderate = 'Moderate'
	Major = 'Major'
	Severe = 'Severe'
