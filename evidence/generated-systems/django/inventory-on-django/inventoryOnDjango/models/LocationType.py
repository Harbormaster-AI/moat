from django.db import models
 #======================================================================
# 
# Encapsulates data for model LocationType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LocationType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LocationType(Enum):   # A subclass of Enum
	Bin = 'Bin'
	Bulk = 'Bulk'
	Staging = 'Staging'
	Dock = 'Dock'
	Picking = 'Picking'
	Packing = 'Packing'
	Quality = 'Quality'
	Return = 'Return'
	ColdStorage = 'ColdStorage'
