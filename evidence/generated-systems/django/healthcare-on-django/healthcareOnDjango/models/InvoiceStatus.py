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
	Draft = 'Draft'
	Issued = 'Issued'
	PartiallyPaid = 'PartiallyPaid'
	Paid = 'Paid'
	Overdue = 'Overdue'
	Cancelled = 'Cancelled'
