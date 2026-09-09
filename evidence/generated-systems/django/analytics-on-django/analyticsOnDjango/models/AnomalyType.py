from django.db import models
 #======================================================================
# 
# Encapsulates data for model AnomalyType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AnomalyType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AnomalyType(Enum):   # A subclass of Enum
	Spike = 'Spike'
	Drop = 'Drop'
	Drift = 'Drift'
	Seasonal = 'Seasonal'
	LevelShift = 'LevelShift'
