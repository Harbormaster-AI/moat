from django.db import models
 #======================================================================
# 
# Encapsulates data for model DeliveryMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DeliveryMethod Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DeliveryMethod(Enum):   # A subclass of Enum
	Classroom = 'Classroom'
	Virtual = 'Virtual'
	SelfPaced = 'SelfPaced'
	Blended = 'Blended'
