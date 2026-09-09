from django.db import models
 #======================================================================
# 
# Encapsulates data for model OrderType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class OrderType(Enum):   # A subclass of Enum
	Market = 'Market'
	Limit = 'Limit'
	Stop = 'Stop'
	StopLimit = 'StopLimit'
