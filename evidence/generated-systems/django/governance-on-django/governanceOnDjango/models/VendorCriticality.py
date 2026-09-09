from django.db import models
 #======================================================================
# 
# Encapsulates data for model VendorCriticality
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class VendorCriticality Declaration (enumerated type)
#======================================================================
from enum import Enum 
class VendorCriticality(Enum):   # A subclass of Enum
	Low = 'Low'
	Medium = 'Medium'
	High = 'High'
	Critical = 'Critical'
