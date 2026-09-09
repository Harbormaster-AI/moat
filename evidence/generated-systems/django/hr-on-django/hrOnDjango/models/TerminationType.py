from django.db import models
 #======================================================================
# 
# Encapsulates data for model TerminationType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerminationType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TerminationType(Enum):   # A subclass of Enum
	Resignation = 'Resignation'
	Dismissal = 'Dismissal'
	Layoff = 'Layoff'
	Retirement = 'Retirement'
	EndOfAssignment = 'EndOfAssignment'
