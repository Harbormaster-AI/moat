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
	New = 'New'
	Screening = 'Screening'
	Interview = 'Interview'
	Offer = 'Offer'
	Hired = 'Hired'
	Rejected = 'Rejected'
	Withdrawn = 'Withdrawn'
