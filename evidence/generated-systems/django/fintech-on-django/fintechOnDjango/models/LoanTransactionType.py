from django.db import models
 #======================================================================
# 
# Encapsulates data for model LoanTransactionType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanTransactionType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LoanTransactionType(Enum):   # A subclass of Enum
	Disbursement = 'Disbursement'
	Repayment = 'Repayment'
	Interest = 'Interest'
	Fee = 'Fee'
	Reversal = 'Reversal'
