from django.db import models
 #======================================================================
# 
# Encapsulates data for model AccrualUnit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccrualUnit Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AccrualUnit(Enum):   # A subclass of Enum
	Hours = 'Hours'
	Days = 'Days'
