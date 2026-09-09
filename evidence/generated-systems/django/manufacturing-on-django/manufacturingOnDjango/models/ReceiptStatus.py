from django.db import models
 #======================================================================
# 
# Encapsulates data for model ReceiptStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReceiptStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ReceiptStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	PartiallyProcessed = 'PartiallyProcessed'
	Completed = 'Completed'
	Rejected = 'Rejected'
