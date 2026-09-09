from django.db import models
from ecommerceOnDjango.models.ReturnStatus import ReturnStatus

#======================================================================
# 
# Encapsulates data for model ReturnRequest
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReturnRequest Declaration
#======================================================================
class ReturnRequest (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	returnNumber = models.CharField(max_length=200, null=True)
	createdAt = models.DateField(null=True)
	refundAmount = Money
	order = models.ForeignKey('Order', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	items = models.ManyToManyField('ReturnItem',  blank=True, related_name='+')
	refund = models.OneToOneField('Refund', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	shipment = models.ForeignKey('Shipment', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReturnStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.returnNumber
		str = str + self.createdAt
		str = str + self.refundAmount
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ReturnRequest";
    
	def objectType(self):
		return "ReturnRequest";
