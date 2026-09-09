from django.db import models
from ecommerceOnDjango.models.PayoutStatus import PayoutStatus

#======================================================================
# 
# Encapsulates data for model Payout
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Payout Declaration
#======================================================================
class Payout (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	payoutNumber = models.CharField(max_length=200, null=True)
	amount = Money
	scheduledDate = models.DateField(null=True)
	paidDate = models.DateField(null=True)
	seller = models.ForeignKey('Seller', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	orders = models.ManyToManyField('Order',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PayoutStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.payoutNumber
		str = str + self.amount
		str = str + self.scheduledDate
		str = str + self.paidDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Payout";
    
	def objectType(self):
		return "Payout";
