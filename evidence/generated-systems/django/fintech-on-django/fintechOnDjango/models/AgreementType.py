from django.db import models
 #======================================================================
# 
# Encapsulates data for model AgreementType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AgreementType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AgreementType(Enum):   # A subclass of Enum
	TermsOfService = 'TermsOfService'
	PrivacyPolicy = 'PrivacyPolicy'
	LoanAgreement = 'LoanAgreement'
	AccountAgreement = 'AccountAgreement'
