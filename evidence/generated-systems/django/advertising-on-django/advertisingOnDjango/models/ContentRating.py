from django.db import models
 #======================================================================
# 
# Encapsulates data for model ContentRating
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContentRating Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ContentRating(Enum):   # A subclass of Enum
	G = 'G'
	PG = 'PG'
	PGThirteen = 'PGThirteen'
	R = 'R'
	Mature = 'Mature'
	Unrated = 'Unrated'
