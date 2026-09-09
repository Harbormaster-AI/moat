from django.db import models
 #======================================================================
# 
# Encapsulates data for model FindingStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FindingStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class FindingStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	InRemediation = 'InRemediation'
	Validated = 'Validated'
	Closed = 'Closed'
