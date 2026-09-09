from django.db import models
 #======================================================================
# 
# Encapsulates data for model InboundShipmentStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InboundShipmentStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InboundShipmentStatus(Enum):   # A subclass of Enum
	Planned = 'Planned'
	Arrived = 'Arrived'
	Received = 'Received'
	Closed = 'Closed'
	Cancelled = 'Cancelled'
