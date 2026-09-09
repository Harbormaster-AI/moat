from django.db import models
 #======================================================================
# 
# Encapsulates data for model PublisherType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PublisherType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PublisherType(Enum):   # A subclass of Enum
	Site = 'Site'
	App = 'App'
	Network = 'Network'
	CTVApp = 'CTVApp'
