from django.db import models
 #======================================================================
# 
# Encapsulates data for model PacingType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PacingType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PacingType(Enum):   # A subclass of Enum
	Even = 'Even'
	ASAP = 'ASAP'
	Smooth = 'Smooth'
