from django.db import models
 #======================================================================
# 
# Encapsulates data for model QualityPlanStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualityPlanStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class QualityPlanStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Released = 'Released'
	Retired = 'Retired'
