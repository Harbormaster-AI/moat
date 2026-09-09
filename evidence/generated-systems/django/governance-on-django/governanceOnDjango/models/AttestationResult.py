from django.db import models
 #======================================================================
# 
# Encapsulates data for model AttestationResult
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AttestationResult Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AttestationResult(Enum):   # A subclass of Enum
	Affirmative = 'Affirmative'
	Negative = 'Negative'
	Qualified = 'Qualified'
