from django.db import models
 #======================================================================
# 
# Encapsulates data for model PerformanceRating
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerformanceRating Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PerformanceRating(Enum):   # A subclass of Enum
	Unsatisfactory = 'Unsatisfactory'
	NeedsImprovement = 'NeedsImprovement'
	MeetsExpectations = 'MeetsExpectations'
	ExceedsExpectations = 'ExceedsExpectations'
	Outstanding = 'Outstanding'
