from django.db import models
 #======================================================================
# 
# Encapsulates data for model ReturnStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReturnStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ReturnStatus(Enum):   # A subclass of Enum
	Requested = 'Requested'
	Approved = 'Approved'
	Rejected = 'Rejected'
	InTransit = 'InTransit'
	Received = 'Received'
	Refunded = 'Refunded'
	Closed = 'Closed'
