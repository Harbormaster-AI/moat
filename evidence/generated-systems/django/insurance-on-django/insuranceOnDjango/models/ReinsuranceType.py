from django.db import models
 #======================================================================
# 
# Encapsulates data for model ReinsuranceType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReinsuranceType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ReinsuranceType(Enum):   # A subclass of Enum
	Treaty = 'Treaty'
	Facultative = 'Facultative'
