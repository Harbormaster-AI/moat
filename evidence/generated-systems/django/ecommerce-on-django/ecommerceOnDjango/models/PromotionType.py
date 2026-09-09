from django.db import models
 #======================================================================
# 
# Encapsulates data for model PromotionType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PromotionType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PromotionType(Enum):   # A subclass of Enum
	Catalog = 'Catalog'
	Cart = 'Cart'
	Shipping = 'Shipping'
