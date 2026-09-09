from django.db import models
 #======================================================================
# 
# Encapsulates data for model BusinessUnitCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BusinessUnitCategory Declaration (enumerated type)
#======================================================================
from enum import Enum 
class BusinessUnitCategory(Enum):   # A subclass of Enum
	ConsumerGoods = 'ConsumerGoods'
	IndustrialEquipment = 'IndustrialEquipment'
	Electronics = 'Electronics'
	Pharmaceuticals = 'Pharmaceuticals'
	FoodBeverage = 'FoodBeverage'
