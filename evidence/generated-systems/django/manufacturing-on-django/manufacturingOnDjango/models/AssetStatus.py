from django.db import models
 #======================================================================
# 
# Encapsulates data for model AssetStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AssetStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AssetStatus(Enum):   # A subclass of Enum
	Commissioned = 'Commissioned'
	Available = 'Available'
	InMaintenance = 'InMaintenance'
	Down = 'Down'
	Retired = 'Retired'
