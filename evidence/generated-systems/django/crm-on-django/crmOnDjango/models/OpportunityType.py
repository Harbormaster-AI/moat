from django.db import models
 #======================================================================
# 
# Encapsulates data for model OpportunityType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OpportunityType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class OpportunityType(Enum):   # A subclass of Enum
	NewBusiness = 'NewBusiness'
	ExistingBusiness = 'ExistingBusiness'
	Renewal = 'Renewal'
	Upsell = 'Upsell'
	CrossSell = 'CrossSell'
