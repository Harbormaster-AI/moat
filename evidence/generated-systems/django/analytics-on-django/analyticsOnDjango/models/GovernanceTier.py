from django.db import models
 #======================================================================
# 
# Encapsulates data for model GovernanceTier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GovernanceTier Declaration (enumerated type)
#======================================================================
from enum import Enum 
class GovernanceTier(Enum):   # A subclass of Enum
	Open = 'Open'
	Internal = 'Internal'
	Restricted = 'Restricted'
	Confidential = 'Confidential'
