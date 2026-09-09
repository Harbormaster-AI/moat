from django.db import models
 #======================================================================
# 
# Encapsulates data for model MaintenanceStrategy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceStrategy Declaration (enumerated type)
#======================================================================
from enum import Enum 
class MaintenanceStrategy(Enum):   # A subclass of Enum
	TimeBased = 'TimeBased'
	UsageBased = 'UsageBased'
	ConditionBased = 'ConditionBased'
	Predictive = 'Predictive'
	Corrective = 'Corrective'
