from django.db import models
 #======================================================================
# 
# Encapsulates data for model BOMStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BOMStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class BOMStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Released = 'Released'
	Obsolete = 'Obsolete'
