from django.db import models
 #======================================================================
# 
# Encapsulates data for model RiskLikelihood
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RiskLikelihood Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RiskLikelihood(Enum):   # A subclass of Enum
	Rare = 'Rare'
	Unlikely = 'Unlikely'
	Possible = 'Possible'
	Likely = 'Likely'
	AlmostCertain = 'AlmostCertain'
