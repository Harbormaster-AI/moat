from django.db import models
 #======================================================================
# 
# Encapsulates data for model MetricType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MetricType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class MetricType(Enum):   # A subclass of Enum
	Ratio = 'Ratio'
	Rate = 'Rate'
	Count = 'Count'
	Percentage = 'Percentage'
	Index = 'Index'
	Score = 'Score'
