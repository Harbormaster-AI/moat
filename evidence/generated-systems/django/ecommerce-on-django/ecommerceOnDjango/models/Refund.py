from django.db import models
from ecommerceOnDjango.models.RefundStatus import RefundStatus

#======================================================================
# 
# Encapsulates data for model Refund
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Refund Declaration
#======================================================================
class Refund (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	refundNumber = models.CharField(max_length=200, null=True)
	amount = Money
	reason = models.CharField(max_length=200, null=True)
	createdAt = models.DateField(null=True)
	payment = models.ForeignKey('Payment', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	order = models.ForeignKey('Order', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RefundStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.refundNumber
		str = str + self.amount
		str = str + self.reason
		str = str + self.createdAt
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Refund";
    
	def objectType(self):
		return "Refund";
