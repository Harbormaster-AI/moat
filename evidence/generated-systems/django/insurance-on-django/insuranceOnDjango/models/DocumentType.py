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
	ApplicationForm = 'ApplicationForm'
	PolicyDocument = 'PolicyDocument'
	Endorsement = 'Endorsement'
	Invoice = 'Invoice'
	ClaimForm = 'ClaimForm'
	PoliceReport = 'PoliceReport'
	Estimate = 'Estimate'
	Photo = 'Photo'
	MedicalRecord = 'MedicalRecord'
	Correspondence = 'Correspondence'
