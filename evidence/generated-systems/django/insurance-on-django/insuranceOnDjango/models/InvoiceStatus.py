from django.db import models
 #======================================================================
# 
# Encapsulates data for model InvoiceStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvoiceStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InvoiceStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	Paid = 'Paid'
	PartiallyPaid = 'PartiallyPaid'
	Void = 'Void'
