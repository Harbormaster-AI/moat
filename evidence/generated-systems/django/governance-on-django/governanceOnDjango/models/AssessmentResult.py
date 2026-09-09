from django.db import models
 #======================================================================
# 
# Encapsulates data for model AssessmentResult
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AssessmentResult Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AssessmentResult(Enum):   # A subclass of Enum
	Pass = 'Pass'
	ConditionalPass = 'ConditionalPass'
	Fail = 'Fail'
