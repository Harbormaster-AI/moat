from django.db import models
 #======================================================================
# 
# Encapsulates data for model AdjusterType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdjusterType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AdjusterType(Enum):   # A subclass of Enum
	Staff = 'Staff'
	Independent = 'Independent'
	Public = 'Public'
