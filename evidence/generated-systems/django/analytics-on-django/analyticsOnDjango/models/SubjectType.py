from django.db import models
 #======================================================================
# 
# Encapsulates data for model SubjectType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubjectType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SubjectType(Enum):   # A subclass of Enum
	User = 'User'
	Group = 'Group'
	Service = 'Service'
