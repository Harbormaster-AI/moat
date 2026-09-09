from django.db import models
 #======================================================================
# 
# Encapsulates data for model ScreeningType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ScreeningType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ScreeningType(Enum):   # A subclass of Enum
	Sanctions = 'Sanctions'
	PEP = 'PEP'
	AdverseMedia = 'AdverseMedia'
