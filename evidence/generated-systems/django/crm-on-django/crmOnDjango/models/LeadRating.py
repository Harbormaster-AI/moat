from django.db import models
 #======================================================================
# 
# Encapsulates data for model LeadRating
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeadRating Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LeadRating(Enum):   # A subclass of Enum
	Hot = 'Hot'
	Warm = 'Warm'
	Cold = 'Cold'
