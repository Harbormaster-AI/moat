from django.db import models
 #======================================================================
# 
# Encapsulates data for model ConversionEventType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConversionEventType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ConversionEventType(Enum):   # A subclass of Enum
	Lead = 'Lead'
	Purchase = 'Purchase'
	Signup = 'Signup'
	AddToCart = 'AddToCart'
	ViewContent = 'ViewContent'
	AppInstall = 'AppInstall'
