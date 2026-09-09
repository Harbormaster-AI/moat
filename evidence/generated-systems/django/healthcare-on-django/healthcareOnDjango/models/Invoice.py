from django.db import models
from healthcareOnDjango.models.InvoiceStatus import InvoiceStatus

#======================================================================
# 
# Encapsulates data for model Invoice
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Invoice Declaration
#======================================================================
class Invoice (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	invoiceNumber = models.CharField(max_length=200, null=True)
	totalAmount = Money
	dueDate = models.DateField(null=True)
	patient = models.ForeignKey('Patient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	claim = models.ForeignKey('Claim', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	payments = models.ManyToManyField('Payment',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InvoiceStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.invoiceNumber
		str = str + self.totalAmount
		str = str + self.dueDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Invoice";
    
	def objectType(self):
		return "Invoice";
