from django.db import models
 #======================================================================
# 
# Encapsulates data for model DataClassificationLevel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataClassificationLevel Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DataClassificationLevel(Enum):   # A subclass of Enum
	Public = 'Public'
	Internal = 'Internal'
	Confidential = 'Confidential'
	Restricted = 'Restricted'
	HighlyRestricted = 'HighlyRestricted'
