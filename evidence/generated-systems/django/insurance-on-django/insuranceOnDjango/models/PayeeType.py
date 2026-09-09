from django.db import models
 #======================================================================
# 
# Encapsulates data for model PayeeType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayeeType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PayeeType(Enum):   # A subclass of Enum
	Claimant = 'Claimant'
	Beneficiary = 'Beneficiary'
	ServiceProvider = 'ServiceProvider'
	Lienholder = 'Lienholder'
	Attorney = 'Attorney'
