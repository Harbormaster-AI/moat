from django.db import models
 #======================================================================
# 
# Encapsulates data for model QualitySeverity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualitySeverity Declaration (enumerated type)
#======================================================================
from enum import Enum 
class QualitySeverity(Enum):   # A subclass of Enum
	Minor = 'Minor'
	Major = 'Major'
	Critical = 'Critical'
