from django.db import models
 #======================================================================
# 
# Encapsulates data for model ForecastCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ForecastCategory Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ForecastCategory(Enum):   # A subclass of Enum
	Pipeline = 'Pipeline'
	BestCase = 'BestCase'
	Commit = 'Commit'
	Omitted = 'Omitted'
	Closed = 'Closed'
