from django.db import models
 #======================================================================
# 
# Encapsulates data for model ModelLifecycle
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ModelLifecycle Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ModelLifecycle(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Staging = 'Staging'
	Production = 'Production'
	Archived = 'Archived'
