from django.db import models
 #======================================================================
# 
# Encapsulates data for model PaymentMethodType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentMethodType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PaymentMethodType(Enum):   # A subclass of Enum
	DirectDeposit = 'DirectDeposit'
	Check = 'Check'
	Cash = 'Cash'
	InternationalTransfer = 'InternationalTransfer'
