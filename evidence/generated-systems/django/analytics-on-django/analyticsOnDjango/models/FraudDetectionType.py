from django.db import models
 #======================================================================
# 
# Encapsulates data for model FraudDetectionType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FraudDetectionType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class FraudDetectionType(Enum):   # A subclass of Enum
	RuleBased = 'RuleBased'
	SupervisedML = 'SupervisedML'
	UnsupervisedML = 'UnsupervisedML'
	Hybrid = 'Hybrid'
