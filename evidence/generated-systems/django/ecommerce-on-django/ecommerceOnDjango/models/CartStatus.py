from django.db import models
 #======================================================================
# 
# Encapsulates data for model CartStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CartStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CartStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Merged = 'Merged'
	Ordered = 'Ordered'
	Abandoned = 'Abandoned'
