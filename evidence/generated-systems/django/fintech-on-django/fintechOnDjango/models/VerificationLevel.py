from django.db import models
 #======================================================================
# 
# Encapsulates data for model VerificationLevel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class VerificationLevel Declaration (enumerated type)
#======================================================================
from enum import Enum 
class VerificationLevel(Enum):   # A subclass of Enum
	Basic = 'Basic'
	Standard = 'Standard'
	Enhanced = 'Enhanced'
