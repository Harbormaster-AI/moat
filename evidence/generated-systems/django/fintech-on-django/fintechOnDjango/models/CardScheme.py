from django.db import models
 #======================================================================
# 
# Encapsulates data for model CardScheme
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CardScheme Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CardScheme(Enum):   # A subclass of Enum
	Visa = 'Visa'
	Mastercard = 'Mastercard'
	Amex = 'Amex'
	Discover = 'Discover'
	UnionPay = 'UnionPay'
