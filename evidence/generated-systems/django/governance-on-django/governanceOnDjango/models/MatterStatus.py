from django.db import models
 #======================================================================
# 
# Encapsulates data for model MatterStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MatterStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class MatterStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	ActiveDiscovery = 'ActiveDiscovery'
	Negotiation = 'Negotiation'
	Settled = 'Settled'
	Closed = 'Closed'
