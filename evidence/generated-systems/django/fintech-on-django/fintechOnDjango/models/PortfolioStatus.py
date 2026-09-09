from django.db import models
 #======================================================================
# 
# Encapsulates data for model PortfolioStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PortfolioStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PortfolioStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Closed = 'Closed'
	Suspended = 'Suspended'
