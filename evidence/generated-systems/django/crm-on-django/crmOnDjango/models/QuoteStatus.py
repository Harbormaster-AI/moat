from django.db import models
 #======================================================================
# 
# Encapsulates data for model QuoteStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QuoteStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class QuoteStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Presented = 'Presented'
	Approved = 'Approved'
	Rejected = 'Rejected'
	Accepted = 'Accepted'
	Expired = 'Expired'
	Withdrawn = 'Withdrawn'
