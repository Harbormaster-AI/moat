from django.db import models
 #======================================================================
# 
# Encapsulates data for model FrequencyScope
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FrequencyScope Declaration (enumerated type)
#======================================================================
from enum import Enum 
class FrequencyScope(Enum):   # A subclass of Enum
	Campaign = 'Campaign'
	LineItem = 'LineItem'
	Creative = 'Creative'
