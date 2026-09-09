from django.db import models
 #======================================================================
# 
# Encapsulates data for model UnderwritingDecisionType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UnderwritingDecisionType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class UnderwritingDecisionType(Enum):   # A subclass of Enum
	Approve = 'Approve'
	ConditionalApprove = 'ConditionalApprove'
	Refer = 'Refer'
	Decline = 'Decline'
