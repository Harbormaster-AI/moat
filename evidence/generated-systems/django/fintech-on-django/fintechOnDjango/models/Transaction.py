from django.db import models
from fintechOnDjango.models.TransactionType import TransactionType
from fintechOnDjango.models.TransactionStatus import TransactionStatus

#======================================================================
# 
# Encapsulates data for model Transaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Transaction Declaration
#======================================================================
class Transaction (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	amount = Money
	fee = Money
	exchangeRate = models.CharField(max_length=64, null=True)
	createdAt = DateTime
	completedAt = DateTime
	narrative = models.CharField(max_length=200, null=True)
	account = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	wallet = models.ForeignKey('Wallet', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	paymentOrder = models.ForeignKey('PaymentOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	card = models.ForeignKey('PaymentCard', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	relatedTransactions = models.ManyToManyField('Transaction',  blank=True, related_name='+')
	alerts = models.ManyToManyField('ComplianceAlert',  blank=True, related_name='+')
	transactionType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TransactionType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TransactionStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.amount
		str = str + self.fee
		str = str + self.exchangeRate
		str = str + self.createdAt
		str = str + self.completedAt
		str = str + self.narrative
		str = str + self.transactionType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Transaction";
    
	def objectType(self):
		return "Transaction";
