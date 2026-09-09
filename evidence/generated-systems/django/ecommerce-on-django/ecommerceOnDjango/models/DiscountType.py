from django.db import models
 #======================================================================
# 
# Encapsulates data for model DiscountType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DiscountType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DiscountType(Enum):   # A subclass of Enum
	AmountOff = 'AmountOff'
	PercentOff = 'PercentOff'
	BuyXGetY = 'BuyXGetY'
	FreeShipping = 'FreeShipping'
