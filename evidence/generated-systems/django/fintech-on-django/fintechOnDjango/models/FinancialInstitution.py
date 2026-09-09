from django.db import models

#======================================================================
# 
# Encapsulates data for model FinancialInstitution
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FinancialInstitution Declaration
#======================================================================
class FinancialInstitution (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	legalName = models.CharField(max_length=200, null=True)
	countryOfIncorporation = models.CharField(max_length=200, null=True)
	bic = BIC
	website = models.CharField(max_length=200, null=True)
	branches = models.ManyToManyField('Branch',  blank=True, related_name='+')
	customers = models.ManyToManyField('Customer',  blank=True, related_name='+')
	productOfferings = models.ManyToManyField('ProductOffering',  blank=True, related_name='+')
	paymentProcessors = models.ManyToManyField('PaymentProcessor',  blank=True, related_name='+')
	compliancePolicies = models.ManyToManyField('CompliancePolicy',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.legalName
		str = str + self.countryOfIncorporation
		str = str + self.bic
		str = str + self.website
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "FinancialInstitution";
    
	def objectType(self):
		return "FinancialInstitution";
