from django.db import models
 #======================================================================
# 
# Encapsulates data for model ExceptionStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExceptionStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ExceptionStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Submitted = 'Submitted'
	Approved = 'Approved'
	Rejected = 'Rejected'
	Expired = 'Expired'
