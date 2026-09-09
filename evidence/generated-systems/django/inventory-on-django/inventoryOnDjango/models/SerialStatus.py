from django.db import models
 #======================================================================
# 
# Encapsulates data for model SerialStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SerialStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SerialStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Assigned = 'Assigned'
	InTransit = 'InTransit'
	Consumed = 'Consumed'
	Returned = 'Returned'
	Scrapped = 'Scrapped'
