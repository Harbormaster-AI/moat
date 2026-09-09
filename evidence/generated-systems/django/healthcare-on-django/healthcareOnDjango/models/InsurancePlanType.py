from django.db import models
 #======================================================================
# 
# Encapsulates data for model InsurancePlanType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsurancePlanType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InsurancePlanType(Enum):   # A subclass of Enum
	HMO = 'HMO'
	PPO = 'PPO'
	EPO = 'EPO'
	POS = 'POS'
	Indemnity = 'Indemnity'
	MedicareAdvantage = 'MedicareAdvantage'
	MedicaidManagedCare = 'MedicaidManagedCare'
