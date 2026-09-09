from django.db import models
 #======================================================================
# 
# Encapsulates data for model ExceptionType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExceptionType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ExceptionType(Enum):   # A subclass of Enum
	PolicyException = 'PolicyException'
	ControlException = 'ControlException'
	RetentionException = 'RetentionException'
	RiskAcceptance = 'RiskAcceptance'
	ComplianceWaiver = 'ComplianceWaiver'
