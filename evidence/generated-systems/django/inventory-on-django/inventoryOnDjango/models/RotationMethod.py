from django.db import models
 #======================================================================
# 
# Encapsulates data for model RotationMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RotationMethod Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RotationMethod(Enum):   # A subclass of Enum
	FIFO = 'FIFO'
	LIFO = 'LIFO'
	FEFO = 'FEFO'
