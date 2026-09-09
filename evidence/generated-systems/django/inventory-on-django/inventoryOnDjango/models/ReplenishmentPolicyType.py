from django.db import models
 #======================================================================
# 
# Encapsulates data for model ReplenishmentPolicyType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReplenishmentPolicyType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ReplenishmentPolicyType(Enum):   # A subclass of Enum
	MinMax = 'MinMax'
	ReorderPoint = 'ReorderPoint'
	EOQ = 'EOQ'
	Kanban = 'Kanban'
