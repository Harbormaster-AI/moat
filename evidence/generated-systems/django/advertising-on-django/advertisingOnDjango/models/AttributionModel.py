from django.db import models
 #======================================================================
# 
# Encapsulates data for model AttributionModel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AttributionModel Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AttributionModel(Enum):   # A subclass of Enum
	LastClick = 'LastClick'
	FirstTouch = 'FirstTouch'
	Linear = 'Linear'
	TimeDecay = 'TimeDecay'
	PositionBased = 'PositionBased'
	DataDriven = 'DataDriven'
