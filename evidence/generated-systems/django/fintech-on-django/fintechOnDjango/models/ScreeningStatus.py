from django.db import models
 #======================================================================
# 
# Encapsulates data for model ScreeningStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ScreeningStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ScreeningStatus(Enum):   # A subclass of Enum
	Clear = 'Clear'
	Review = 'Review'
	Match = 'Match'
