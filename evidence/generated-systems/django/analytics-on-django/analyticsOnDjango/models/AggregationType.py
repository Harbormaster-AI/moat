from django.db import models
 #======================================================================
# 
# Encapsulates data for model AggregationType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AggregationType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AggregationType(Enum):   # A subclass of Enum
	Sum = 'Sum'
	Average = 'Average'
	Min = 'Min'
	Max = 'Max'
	Median = 'Median'
	Count = 'Count'
	DistinctCount = 'DistinctCount'
