from django.db import models
 #======================================================================
# 
# Encapsulates data for model AssessmentType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AssessmentType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AssessmentType(Enum):   # A subclass of Enum
	SelfAssessment = 'SelfAssessment'
	InternalAssessment = 'InternalAssessment'
	ExternalAssessment = 'ExternalAssessment'
	ReadinessReview = 'ReadinessReview'
