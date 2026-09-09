from django.db import models
 #======================================================================
# 
# Encapsulates data for model ComparisonOperator
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComparisonOperator Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ComparisonOperator(Enum):   # A subclass of Enum
	GreaterThan = 'GreaterThan'
	GreaterThanOrEqual = 'GreaterThanOrEqual'
	LessThan = 'LessThan'
	LessThanOrEqual = 'LessThanOrEqual'
	Equal = 'Equal'
	NotEqual = 'NotEqual'
