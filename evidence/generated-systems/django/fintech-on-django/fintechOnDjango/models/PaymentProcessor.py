from django.db import models

#======================================================================
# 
# Encapsulates data for model PaymentProcessor
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentProcessor Declaration
#======================================================================
class PaymentProcessor (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	processorCode = models.CharField(max_length=200, null=True)
	networkSupport = models.CharField(max_length=200, null=True)
	institutions = models.ManyToManyField('FinancialInstitution',  blank=True, related_name='+')
	contracts = models.ManyToManyField('PaymentContract',  blank=True, related_name='+')
	settlements = models.ManyToManyField('SettlementBatch',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.processorCode
		str = str + self.networkSupport
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PaymentProcessor";
    
	def objectType(self):
		return "PaymentProcessor";
