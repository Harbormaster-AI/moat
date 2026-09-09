from django.db import models
 #======================================================================
# 
# Encapsulates data for model LimitScope
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LimitScope Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LimitScope(Enum):   # A subclass of Enum
	PerTransaction = 'PerTransaction'
	Daily = 'Daily'
	Monthly = 'Monthly'
	Yearly = 'Yearly'
	Rolling24h = 'Rolling24h'
