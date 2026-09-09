from django.db import models
 #======================================================================
# 
# Encapsulates data for model TerminalStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerminalStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TerminalStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Inactive = 'Inactive'
	Decommissioned = 'Decommissioned'
