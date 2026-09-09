from django.db import models
 #======================================================================
# 
# Encapsulates data for model ApplicationStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ApplicationStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ApplicationStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Submitted = 'Submitted'
	Underwriting = 'Underwriting'
	Approved = 'Approved'
	Declined = 'Declined'
	Withdrawn = 'Withdrawn'
