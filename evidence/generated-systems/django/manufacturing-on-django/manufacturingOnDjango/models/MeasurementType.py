from django.db import models
 #======================================================================
# 
# Encapsulates data for model MeasurementType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MeasurementType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class MeasurementType(Enum):   # A subclass of Enum
	Attribute = 'Attribute'
	Variable = 'Variable'
