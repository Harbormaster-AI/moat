from django.db import models
 #======================================================================
# 
# Encapsulates data for model MediaType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MediaType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class MediaType(Enum):   # A subclass of Enum
	Image = 'Image'
	Video = 'Video'
	Document = 'Document'
	Audio = 'Audio'
	Other = 'Other'
