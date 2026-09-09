from django.db import models
 #======================================================================
# 
# Encapsulates data for model DataProviderType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataProviderType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DataProviderType(Enum):   # A subclass of Enum
	FirstParty = 'FirstParty'
	SecondParty = 'SecondParty'
	ThirdParty = 'ThirdParty'
