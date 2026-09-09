from django.db import models
 #======================================================================
# 
# Encapsulates data for model DocumentStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DocumentStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DocumentStatus(Enum):   # A subclass of Enum
	Submitted = 'Submitted'
	Approved = 'Approved'
	Rejected = 'Rejected'
	Expired = 'Expired'
