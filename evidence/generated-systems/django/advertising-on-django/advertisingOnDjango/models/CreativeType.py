from django.db import models
 #======================================================================
# 
# Encapsulates data for model CreativeType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CreativeType(Enum):   # A subclass of Enum
	Image = 'Image'
	Video = 'Video'
	HTML5 = 'HTML5'
	Audio = 'Audio'
