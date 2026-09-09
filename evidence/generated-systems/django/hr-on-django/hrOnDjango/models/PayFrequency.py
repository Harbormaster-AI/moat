from django.db import models
 #======================================================================
# 
# Encapsulates data for model PayFrequency
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayFrequency Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PayFrequency(Enum):   # A subclass of Enum
	Weekly = 'Weekly'
	Biweekly = 'Biweekly'
	Semimonthly = 'Semimonthly'
	Monthly = 'Monthly'
	Quarterly = 'Quarterly'
