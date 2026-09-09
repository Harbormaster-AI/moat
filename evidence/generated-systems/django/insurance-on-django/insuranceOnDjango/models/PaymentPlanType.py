from django.db import models
 #======================================================================
# 
# Encapsulates data for model PaymentPlanType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentPlanType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PaymentPlanType(Enum):   # A subclass of Enum
	Annual = 'Annual'
	SemiAnnual = 'SemiAnnual'
	Quarterly = 'Quarterly'
	Monthly = 'Monthly'
	PayInFull = 'PayInFull'
