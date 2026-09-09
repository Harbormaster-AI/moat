from django.db import models
 #======================================================================
# 
# Encapsulates data for model DataTaskType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataTaskType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DataTaskType(Enum):   # A subclass of Enum
	Extract = 'Extract'
	Transform = 'Transform'
	Load = 'Load'
	Validate = 'Validate'
	Enrich = 'Enrich'
