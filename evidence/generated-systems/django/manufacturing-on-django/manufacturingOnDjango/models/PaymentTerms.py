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
	Net30 = 'Net30'
	Net45 = 'Net45'
	Net60 = 'Net60'
	Prepaid = 'Prepaid'
	COD = 'COD'
