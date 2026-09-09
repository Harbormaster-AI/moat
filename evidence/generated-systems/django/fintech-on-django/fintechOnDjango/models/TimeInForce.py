from django.db import models
 #======================================================================
# 
# Encapsulates data for model TimeInForce
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimeInForce Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TimeInForce(Enum):   # A subclass of Enum
	Day = 'Day'
	GTC = 'GTC'
	IOC = 'IOC'
	FOK = 'FOK'
