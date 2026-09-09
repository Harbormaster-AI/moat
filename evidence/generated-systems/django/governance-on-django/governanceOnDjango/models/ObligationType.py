from django.db import models
 #======================================================================
# 
# Encapsulates data for model ObligationType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ObligationType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ObligationType(Enum):   # A subclass of Enum
	Regulatory = 'Regulatory'
	Contractual = 'Contractual'
	PolicyDerived = 'PolicyDerived'
	IndustryStandard = 'IndustryStandard'
