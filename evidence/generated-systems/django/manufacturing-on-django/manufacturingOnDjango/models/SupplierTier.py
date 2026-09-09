from django.db import models
 #======================================================================
# 
# Encapsulates data for model SupplierTier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SupplierTier Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SupplierTier(Enum):   # A subclass of Enum
	Tier1 = 'Tier1'
	Tier2 = 'Tier2'
	Tier3 = 'Tier3'
