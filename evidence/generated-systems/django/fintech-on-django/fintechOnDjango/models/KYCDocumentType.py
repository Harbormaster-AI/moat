from django.db import models
 #======================================================================
# 
# Encapsulates data for model KYCDocumentType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KYCDocumentType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class KYCDocumentType(Enum):   # A subclass of Enum
	Passport = 'Passport'
	NationalID = 'NationalID'
	DriverLicense = 'DriverLicense'
	BusinessRegistration = 'BusinessRegistration'
	ProofOfAddress = 'ProofOfAddress'
