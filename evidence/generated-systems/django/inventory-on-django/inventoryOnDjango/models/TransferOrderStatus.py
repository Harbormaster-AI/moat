from django.db import models
 #======================================================================
# 
# Encapsulates data for model TransferOrderStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TransferOrderStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TransferOrderStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Released = 'Released'
	InTransit = 'InTransit'
	Received = 'Received'
	Closed = 'Closed'
	Cancelled = 'Cancelled'
