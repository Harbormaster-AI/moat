from django.db import models
 #======================================================================
# 
# Encapsulates data for model PaymentPriority
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentPriority Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PaymentPriority(Enum):   # A subclass of Enum
	Normal = 'Normal'
	Urgent = 'Urgent'
