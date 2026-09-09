from django.db import models
 #======================================================================
# 
# Encapsulates data for model TargetingOperator
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TargetingOperator Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TargetingOperator(Enum):   # A subclass of Enum
	Include = 'Include'
	Exclude = 'Exclude'
