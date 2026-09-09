from django.db import models
 #======================================================================
# 
# Encapsulates data for model FeeType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeeType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class FeeType(Enum):   # A subclass of Enum
	Fixed = 'Fixed'
	Percentage = 'Percentage'
	Tiered = 'Tiered'
	Interchange = 'Interchange'
	Network = 'Network'
	Chargeback = 'Chargeback'
	ATM = 'ATM'
	FX = 'FX'
