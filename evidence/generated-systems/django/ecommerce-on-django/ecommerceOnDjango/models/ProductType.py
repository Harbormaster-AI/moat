from django.db import models
 #======================================================================
# 
# Encapsulates data for model ProductType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ProductType(Enum):   # A subclass of Enum
	Physical = 'Physical'
	Digital = 'Digital'
	Service = 'Service'
	Bundle = 'Bundle'
	Subscription = 'Subscription'
