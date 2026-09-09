from django.db import models
from hrOnDjango.models.PaymentMethodType import PaymentMethodType

#======================================================================
# 
# Encapsulates data for model PaymentMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentMethod Declaration
#======================================================================
class PaymentMethod (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	preferred = models.BooleanField(null=True)
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	bankAccount = models.OneToOneField('BankAccount', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	methodType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentMethodType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.preferred
		str = str + self.methodType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PaymentMethod";
    
	def objectType(self):
		return "PaymentMethod";
