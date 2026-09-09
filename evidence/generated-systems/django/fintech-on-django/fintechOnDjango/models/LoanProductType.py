from django.db import models
 #======================================================================
# 
# Encapsulates data for model LoanProductType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanProductType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LoanProductType(Enum):   # A subclass of Enum
	PersonalLoan = 'PersonalLoan'
	Mortgage = 'Mortgage'
	InstallmentLoan = 'InstallmentLoan'
	CreditLine = 'CreditLine'
	SME = 'SME'
