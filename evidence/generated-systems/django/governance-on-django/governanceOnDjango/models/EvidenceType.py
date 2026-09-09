from django.db import models
 #======================================================================
# 
# Encapsulates data for model EvidenceType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EvidenceType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class EvidenceType(Enum):   # A subclass of Enum
	Document = 'Document'
	Screenshot = 'Screenshot'
	LogExport = 'LogExport'
	SystemReport = 'SystemReport'
	Ticket = 'Ticket'
	Attestation = 'Attestation'
	Configuration = 'Configuration'
	Dataset = 'Dataset'
