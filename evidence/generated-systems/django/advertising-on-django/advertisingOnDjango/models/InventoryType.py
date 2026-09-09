from django.db import models
 #======================================================================
# 
# Encapsulates data for model InventoryType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InventoryType(Enum):   # A subclass of Enum
	Display = 'Display'
	Video = 'Video'
	Native = 'Native'
	Audio = 'Audio'
	Search = 'Search'
	Social = 'Social'
	DOOH = 'DOOH'
