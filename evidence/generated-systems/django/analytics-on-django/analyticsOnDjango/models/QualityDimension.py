from django.db import models
 #======================================================================
# 
# Encapsulates data for model QualityDimension
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualityDimension Declaration (enumerated type)
#======================================================================
from enum import Enum 
class QualityDimension(Enum):   # A subclass of Enum
	Completeness = 'Completeness'
	Accuracy = 'Accuracy'
	Consistency = 'Consistency'
	Timeliness = 'Timeliness'
	Uniqueness = 'Uniqueness'
	Validity = 'Validity'
