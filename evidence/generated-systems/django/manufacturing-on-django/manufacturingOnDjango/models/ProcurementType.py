from django.db import models
 #======================================================================
# 
# Encapsulates data for model ProcurementType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProcurementType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ProcurementType(Enum):   # A subclass of Enum
	MakeToStock = 'MakeToStock'
	MakeToOrder = 'MakeToOrder'
	Purchase = 'Purchase'
	Kanban = 'Kanban'
	Outsourced = 'Outsourced'
