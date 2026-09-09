from django.db import models
 #======================================================================
# 
# Encapsulates data for model Applicability
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Applicability Declaration (enumerated type)
#======================================================================
from enum import Enum 
class Applicability(Enum):   # A subclass of Enum
	Mandatory = 'Mandatory'
	Recommended = 'Recommended'
	NotApplicable = 'NotApplicable'
