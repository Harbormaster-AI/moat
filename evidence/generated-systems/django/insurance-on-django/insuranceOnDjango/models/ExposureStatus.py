from django.db import models
 #======================================================================
# 
# Encapsulates data for model ExposureStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExposureStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ExposureStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	Closed = 'Closed'
	Pending = 'Pending'
	Reserved = 'Reserved'
