from django.db import models
 #======================================================================
# 
# Encapsulates data for model PayerType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayerType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PayerType(Enum):   # A subclass of Enum
	Commercial = 'Commercial'
	Government = 'Government'
	SelfInsured = 'SelfInsured'
