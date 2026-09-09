from django.db import models
 #======================================================================
# 
# Encapsulates data for model ProductLifecycleStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductLifecycleStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ProductLifecycleStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	PendingApproval = 'PendingApproval'
	Discontinued = 'Discontinued'
	Obsolete = 'Obsolete'
