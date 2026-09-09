from django.db import models
 #======================================================================
# 
# Encapsulates data for model TeamType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TeamType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TeamType(Enum):   # A subclass of Enum
	Sales = 'Sales'
	Service = 'Service'
	Marketing = 'Marketing'
	AccountTeam = 'AccountTeam'
	DealDesk = 'DealDesk'
	CrossFunctional = 'CrossFunctional'
