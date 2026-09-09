from django.db import models
 #======================================================================
# 
# Encapsulates data for model BenefitType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BenefitType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class BenefitType(Enum):   # A subclass of Enum
	Medical = 'Medical'
	Dental = 'Dental'
	Vision = 'Vision'
	LifeInsurance = 'LifeInsurance'
	Disability = 'Disability'
	Retirement = 'Retirement'
	Wellness = 'Wellness'
