from django.db import models
 #======================================================================
# 
# Encapsulates data for model DimensionType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DimensionType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DimensionType(Enum):   # A subclass of Enum
	Categorical = 'Categorical'
	Temporal = 'Temporal'
	Geospatial = 'Geospatial'
	Hierarchical = 'Hierarchical'
