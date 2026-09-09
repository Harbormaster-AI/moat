from django.db import models
 #======================================================================
# 
# Encapsulates data for model SubscriptionInterval
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubscriptionInterval Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SubscriptionInterval(Enum):   # A subclass of Enum
	Weekly = 'Weekly'
	BiWeekly = 'BiWeekly'
	Monthly = 'Monthly'
	Quarterly = 'Quarterly'
	SemiAnnual = 'SemiAnnual'
	Annual = 'Annual'
