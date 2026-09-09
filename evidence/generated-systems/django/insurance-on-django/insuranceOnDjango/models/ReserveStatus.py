from django.db import models
 #======================================================================
# 
# Encapsulates data for model ReserveStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReserveStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ReserveStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	Released = 'Released'
	Increased = 'Increased'
	Decreased = 'Decreased'
	Closed = 'Closed'
