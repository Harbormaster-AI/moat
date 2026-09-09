from django.db import models
 #======================================================================
# 
# Encapsulates data for model ExemptStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExemptStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ExemptStatus(Enum):   # A subclass of Enum
	Exempt = 'Exempt'
	NonExempt = 'NonExempt'
