from django.db import models
 #======================================================================
# 
# Encapsulates data for model WorkLocationType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkLocationType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class WorkLocationType(Enum):   # A subclass of Enum
	Onsite = 'Onsite'
	Hybrid = 'Hybrid'
	Remote = 'Remote'
