from django.db import models
 #======================================================================
# 
# Encapsulates data for model RiskStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RiskStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RiskStatus(Enum):   # A subclass of Enum
	Identified = 'Identified'
	Assessed = 'Assessed'
	Mitigated = 'Mitigated'
	Accepted = 'Accepted'
	Transferred = 'Transferred'
	Closed = 'Closed'
