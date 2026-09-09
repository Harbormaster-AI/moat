from django.db import models
 #======================================================================
# 
# Encapsulates data for model DataSubjectRequestType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSubjectRequestType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DataSubjectRequestType(Enum):   # A subclass of Enum
	Access = 'Access'
	Rectification = 'Rectification'
	Erasure = 'Erasure'
	Restriction = 'Restriction'
	Portability = 'Portability'
	Objection = 'Objection'
	AutomatedDecisioningReview = 'AutomatedDecisioningReview'
