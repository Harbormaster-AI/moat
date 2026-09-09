from django.db import models
 #======================================================================
# 
# Encapsulates data for model SerializationMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SerializationMethod Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SerializationMethod(Enum):   # A subclass of Enum
	Serialized = 'Serialized'
	LotTracked = 'LotTracked'
	None = 'None'
