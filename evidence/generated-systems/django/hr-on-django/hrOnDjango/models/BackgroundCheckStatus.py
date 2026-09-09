from django.db import models
 #======================================================================
# 
# Encapsulates data for model BackgroundCheckStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BackgroundCheckStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class BackgroundCheckStatus(Enum):   # A subclass of Enum
	Ordered = 'Ordered'
	InProgress = 'InProgress'
	Clear = 'Clear'
	Adverse = 'Adverse'
	Cancelled = 'Cancelled'
