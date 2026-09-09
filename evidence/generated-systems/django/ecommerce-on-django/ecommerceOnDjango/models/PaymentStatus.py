from django.db import models
 #======================================================================
# 
# Encapsulates data for model PaymentStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PaymentStatus(Enum):   # A subclass of Enum
	Authorized = 'Authorized'
	Captured = 'Captured'
	PartiallyCaptured = 'PartiallyCaptured'
	Declined = 'Declined'
	Refunded = 'Refunded'
	PartiallyRefunded = 'PartiallyRefunded'
	Voided = 'Voided'
	Pending = 'Pending'
