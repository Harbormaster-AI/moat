from django.db import models
 #======================================================================
# 
# Encapsulates data for model ShippingMethodType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShippingMethodType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ShippingMethodType(Enum):   # A subclass of Enum
	Standard = 'Standard'
	Expedited = 'Expedited'
	Overnight = 'Overnight'
	SameDay = 'SameDay'
	Pickup = 'Pickup'
