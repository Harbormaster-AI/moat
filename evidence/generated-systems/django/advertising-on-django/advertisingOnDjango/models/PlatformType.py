from django.db import models
 #======================================================================
# 
# Encapsulates data for model PlatformType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlatformType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PlatformType(Enum):   # A subclass of Enum
	Web = 'Web'
	MobileApp = 'MobileApp'
	CTV = 'CTV'
