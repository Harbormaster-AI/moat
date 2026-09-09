from django.db import models
 #======================================================================
# 
# Encapsulates data for model ProducerStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProducerStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ProducerStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Suspended = 'Suspended'
	Terminated = 'Terminated'
