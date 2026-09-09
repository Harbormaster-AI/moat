from django.db import models
 #======================================================================
# 
# Encapsulates data for model UserStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UserStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class UserStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Inactive = 'Inactive'
	Locked = 'Locked'
	PendingInvite = 'PendingInvite'
