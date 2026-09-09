from django.db import models
 #======================================================================
# 
# Encapsulates data for model NonconformanceStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NonconformanceStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class NonconformanceStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	Contained = 'Contained'
	UnderInvestigation = 'UnderInvestigation'
	Dispositioned = 'Dispositioned'
	Closed = 'Closed'
