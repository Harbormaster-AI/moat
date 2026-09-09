from django.db import models
 #======================================================================
# 
# Encapsulates data for model PaymentMethodType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentMethodType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PaymentMethodType(Enum):   # A subclass of Enum
	CreditCard = 'CreditCard'
	DebitCard = 'DebitCard'
	PayPal = 'PayPal'
	BankTransfer = 'BankTransfer'
	CashOnDelivery = 'CashOnDelivery'
	GiftCard = 'GiftCard'
	ApplePay = 'ApplePay'
	GooglePay = 'GooglePay'
	BuyNowPayLater = 'BuyNowPayLater'
