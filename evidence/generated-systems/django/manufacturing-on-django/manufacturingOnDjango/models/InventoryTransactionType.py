from django.db import models
 #======================================================================
# 
# Encapsulates data for model InventoryTransactionType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryTransactionType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InventoryTransactionType(Enum):   # A subclass of Enum
	Receipt = 'Receipt'
	Issue = 'Issue'
	Return = 'Return'
	Adjustment = 'Adjustment'
	Transfer = 'Transfer'
	Consumption = 'Consumption'
	ProductionReceipt = 'ProductionReceipt'
	Scrap = 'Scrap'
