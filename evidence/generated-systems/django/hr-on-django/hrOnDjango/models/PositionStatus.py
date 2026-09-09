from django.db import models
 #======================================================================
# 
# Encapsulates data for model PositionStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PositionStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PositionStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	Filled = 'Filled'
	Frozen = 'Frozen'
	Closed = 'Closed'
