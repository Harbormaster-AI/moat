from django.db import models
 #======================================================================
# 
# Encapsulates data for model AdjustmentType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdjustmentType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AdjustmentType(Enum):   # A subclass of Enum
	Increase = 'Increase'
	Decrease = 'Decrease'
	Reclassification = 'Reclassification'
