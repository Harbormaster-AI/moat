from django.db import models
 #======================================================================
# 
# Encapsulates data for model PixelType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PixelType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PixelType(Enum):   # A subclass of Enum
	Image = 'Image'
	JavaScript = 'JavaScript'
	ServerSide = 'ServerSide'
