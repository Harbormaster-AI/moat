from django.db import models
 #======================================================================
# 
# Encapsulates data for model WorkAuthorizationStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkAuthorizationStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class WorkAuthorizationStatus(Enum):   # A subclass of Enum
	NotRequired = 'NotRequired'
	Pending = 'Pending'
	Authorized = 'Authorized'
	Expired = 'Expired'
