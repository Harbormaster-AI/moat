from django.db import models
 #======================================================================
# 
# Encapsulates data for model WalletProvider
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WalletProvider Declaration (enumerated type)
#======================================================================
from enum import Enum 
class WalletProvider(Enum):   # A subclass of Enum
	ApplePay = 'ApplePay'
	GooglePay = 'GooglePay'
	SamsungPay = 'SamsungPay'
	Other = 'Other'
