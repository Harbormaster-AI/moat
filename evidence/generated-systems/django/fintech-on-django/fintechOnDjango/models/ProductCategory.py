from django.db import models
 #======================================================================
# 
# Encapsulates data for model ProductCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductCategory Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ProductCategory(Enum):   # A subclass of Enum
	Checking = 'Checking'
	Savings = 'Savings'
	CreditCard = 'CreditCard'
	Loan = 'Loan'
	Investment = 'Investment'
	Insurance = 'Insurance'
	Payments = 'Payments'
	FX = 'FX'
	Wallet = 'Wallet'
