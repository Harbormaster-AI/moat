from django.db import models
 #======================================================================
# 
# Encapsulates data for model PaymentTerms
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentTerms Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PaymentTerms(Enum):   # A subclass of Enum
	Prepaid = 'Prepaid'
	NetFifteen = 'NetFifteen'
	NetThirty = 'NetThirty'
	NetSixty = 'NetSixty'
