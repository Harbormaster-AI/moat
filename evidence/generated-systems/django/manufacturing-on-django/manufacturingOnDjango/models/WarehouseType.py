from django.db import models
 #======================================================================
# 
# Encapsulates data for model WarehouseType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WarehouseType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class WarehouseType(Enum):   # A subclass of Enum
	RawMaterial = 'RawMaterial'
	WIP = 'WIP'
	FinishedGoods = 'FinishedGoods'
	Distribution = 'Distribution'
