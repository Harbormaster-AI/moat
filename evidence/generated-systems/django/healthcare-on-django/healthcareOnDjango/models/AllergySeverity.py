from django.db import models
 #======================================================================
# 
# Encapsulates data for model AllergySeverity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AllergySeverity Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AllergySeverity(Enum):   # A subclass of Enum
	Mild = 'Mild'
	Moderate = 'Moderate'
	Severe = 'Severe'
	LifeThreatening = 'LifeThreatening'
