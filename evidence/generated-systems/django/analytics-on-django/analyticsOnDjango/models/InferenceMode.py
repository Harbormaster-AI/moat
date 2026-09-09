from django.db import models
 #======================================================================
# 
# Encapsulates data for model InferenceMode
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InferenceMode Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InferenceMode(Enum):   # A subclass of Enum
	Batch = 'Batch'
	RealTime = 'RealTime'
