from django.db import models
 #======================================================================
# 
# Encapsulates data for model ShiftType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShiftType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ShiftType(Enum):   # A subclass of Enum
	Day = 'Day'
	Swing = 'Swing'
	Night = 'Night'
	Weekend = 'Weekend'
