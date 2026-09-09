from django.db import models
 #======================================================================
# 
# Encapsulates data for model CaseSeverity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CaseSeverity Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CaseSeverity(Enum):   # A subclass of Enum
	Minor = 'Minor'
	Major = 'Major'
	Critical = 'Critical'
	Blocker = 'Blocker'
