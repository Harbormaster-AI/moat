from django.db import models
 #======================================================================
# 
# Encapsulates data for model DecisionOutcome
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DecisionOutcome Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DecisionOutcome(Enum):   # A subclass of Enum
	Approve = 'Approve'
	Decline = 'Decline'
	Refer = 'Refer'
