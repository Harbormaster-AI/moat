from django.db import models
 #======================================================================
# 
# Encapsulates data for model InventoryAlertType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryAlertType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InventoryAlertType(Enum):   # A subclass of Enum
	BelowMin = 'BelowMin'
	AboveMax = 'AboveMax'
	StockoutRisk = 'StockoutRisk'
	ExcessStock = 'ExcessStock'
	ExpiryRisk = 'ExpiryRisk'
