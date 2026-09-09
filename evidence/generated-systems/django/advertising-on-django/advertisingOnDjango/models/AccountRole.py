from django.db import models
 #======================================================================
# 
# Encapsulates data for model AccountRole
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccountRole Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AccountRole(Enum):   # A subclass of Enum
	Admin = 'Admin'
	Trader = 'Trader'
	Analyst = 'Analyst'
	Viewer = 'Viewer'
