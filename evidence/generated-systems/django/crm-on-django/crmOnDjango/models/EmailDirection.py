from django.db import models
 #======================================================================
# 
# Encapsulates data for model EmailDirection
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmailDirection Declaration (enumerated type)
#======================================================================
from enum import Enum 
class EmailDirection(Enum):   # A subclass of Enum
	Inbound = 'Inbound'
	Outbound = 'Outbound'
	Internal = 'Internal'
