from django.db import models
 #======================================================================
# 
# Encapsulates data for model DocumentType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DocumentType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DocumentType(Enum):   # A subclass of Enum
	Resume = 'Resume'
	CoverLetter = 'CoverLetter'
	ID = 'ID'
	Certification = 'Certification'
	Contract = 'Contract'
	Policy = 'Policy'
	Other = 'Other'
