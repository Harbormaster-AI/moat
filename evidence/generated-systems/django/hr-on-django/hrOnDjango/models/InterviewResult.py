from django.db import models
 #======================================================================
# 
# Encapsulates data for model InterviewResult
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InterviewResult Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InterviewResult(Enum):   # A subclass of Enum
	Pending = 'Pending'
	Proceed = 'Proceed'
	Reject = 'Reject'
	OfferRecommended = 'OfferRecommended'
