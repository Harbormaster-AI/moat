from django.db import models
 #======================================================================
# 
# Encapsulates data for model PaymentMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentMethod Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PaymentMethod(Enum):   # A subclass of Enum
	ACH = 'ACH'
	CreditCard = 'CreditCard'
	DebitCard = 'DebitCard'
	Check = 'Check'
	Cash = 'Cash'
	Wire = 'Wire'
