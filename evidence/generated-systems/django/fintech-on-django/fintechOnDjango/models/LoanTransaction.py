from django.db import models
from fintechOnDjango.models.LoanTransactionType import LoanTransactionType
from fintechOnDjango.models.PostingStatus import PostingStatus

#======================================================================
# 
# Encapsulates data for model LoanTransaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanTransaction Declaration
#======================================================================
class LoanTransaction (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	transactionId = TransactionId
	amount = Money
	postingDate = models.DateField(null=True)
	loan = models.ForeignKey('Loan', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	type = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LoanTransactionType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PostingStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.transactionId
		str = str + self.amount
		str = str + self.postingDate
		str = str + self.type
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "LoanTransaction";
    
	def objectType(self):
		return "LoanTransaction";
