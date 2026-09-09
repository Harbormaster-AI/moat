from django.db import models
 #======================================================================
# 
# Encapsulates data for model ClientType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClientType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ClientType(Enum):   # A subclass of Enum
	Confidential = 'Confidential'
	Public = 'Public'
