from django.db import models
 #======================================================================
# 
# Encapsulates data for model PostingStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PostingStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PostingStatus(Enum):   # A subclass of Enum
	Pending = 'Pending'
	Posted = 'Posted'
	Reversed = 'Reversed'
