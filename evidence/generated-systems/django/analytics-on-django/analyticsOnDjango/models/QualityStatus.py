from django.db import models
 #======================================================================
# 
# Encapsulates data for model QualityStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualityStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class QualityStatus(Enum):   # A subclass of Enum
	Passed = 'Passed'
	Failed = 'Failed'
	Warning = 'Warning'
	Skipped = 'Skipped'
