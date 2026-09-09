from django.db import models
 #======================================================================
# 
# Encapsulates data for model DisputeStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DisputeStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DisputeStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	Represented = 'Represented'
	Won = 'Won'
	Lost = 'Lost'
	Closed = 'Closed'
