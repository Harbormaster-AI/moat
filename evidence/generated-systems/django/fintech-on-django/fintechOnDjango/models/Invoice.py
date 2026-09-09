from django.db import models
from fintechOnDjango.models.InvoiceStatus import InvoiceStatus

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
	issueDate = models.DateField(null=True)
	dueDate = models.DateField(null=True)
	total = Money
	currency = models.CharField(max_length=200, null=True)
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	payments = models.ManyToManyField('PaymentOrder',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InvoiceStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.invoiceNumber
		str = str + self.issueDate
		str = str + self.dueDate
		str = str + self.total
		str = str + self.currency
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Invoice";
    
	def objectType(self):
		return "Invoice";
