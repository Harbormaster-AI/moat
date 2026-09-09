from django.db import models
 #======================================================================
# 
# Encapsulates data for model OrderSide
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderSide Declaration (enumerated type)
#======================================================================
from enum import Enum 
class OrderSide(Enum):   # A subclass of Enum
	Buy = 'Buy'
	Sell = 'Sell'
