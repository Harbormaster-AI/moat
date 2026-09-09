from django.db import models
 #======================================================================
# 
# Encapsulates data for model PaymentProviderType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentProviderType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PaymentProviderType(Enum):   # A subclass of Enum
	PSP = 'PSP'
	Gateway = 'Gateway'
	Aggregator = 'Aggregator'
	Manual = 'Manual'
