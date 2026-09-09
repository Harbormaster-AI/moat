from django.db import models
from fintechOnDjango.models.PaymentMethod import PaymentMethod
from fintechOnDjango.models.PaymentOrderStatus import PaymentOrderStatus
from fintechOnDjango.models.PaymentPriority import PaymentPriority

#======================================================================
# 
# Encapsulates data for model PaymentOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentOrder Declaration
#======================================================================
class PaymentOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	orderReference = models.CharField(max_length=200, null=True)
	requestedExecutionDate = models.DateField(null=True)
	purpose = models.CharField(max_length=200, null=True)
	sourceAccount = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	destinationAccount = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	beneficiary = models.ForeignKey('Beneficiary', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	transactions = models.ManyToManyField('Transaction',  blank=True, related_name='+')
	fxDeal = models.ForeignKey('FXDeal', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	fees = models.ManyToManyField('AppliedFee',  blank=True, related_name='+')
	paymentMethod = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentMethod])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentOrderStatus])
	priority = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentPriority])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.orderReference
		str = str + self.requestedExecutionDate
		str = str + self.purpose
		str = str + self.paymentMethod
		str = str + self.status
		str = str + self.priority
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PaymentOrder";
    
	def objectType(self):
		return "PaymentOrder";
