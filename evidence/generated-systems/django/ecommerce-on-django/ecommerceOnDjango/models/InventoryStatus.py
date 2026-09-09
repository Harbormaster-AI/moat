from django.db import models
 #======================================================================
# 
# Encapsulates data for model InventoryStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InventoryStatus(Enum):   # A subclass of Enum
	InStock = 'InStock'
	LowStock = 'LowStock'
	OutOfStock = 'OutOfStock'
	Backorder = 'Backorder'
	Preorder = 'Preorder'
