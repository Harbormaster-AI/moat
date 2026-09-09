from django.db import models
 #======================================================================
# 
# Encapsulates data for model FXPriceType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FXPriceType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class FXPriceType(Enum):   # A subclass of Enum
	Indicative = 'Indicative'
	Firm = 'Firm'
