from django.db import models
from insuranceOnDjango.models.CustomerType import CustomerType

#======================================================================
# 
# Encapsulates data for model Customer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Customer Declaration
#======================================================================
class Customer (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	organizationName = models.CharField(max_length=200, null=True)
	taxId = models.CharField(max_length=200, null=True)
	dateOfBirth = models.DateField(null=True)
	primaryAddress = Address
	applications = models.ManyToManyField('Application',  blank=True, related_name='+')
	policies = models.ManyToManyField('Policy',  blank=True, related_name='+')
	claims = models.ManyToManyField('Claim',  blank=True, related_name='+')
	agents = models.ManyToManyField('Agent',  blank=True, related_name='+')
	beneficiaries = models.ManyToManyField('Beneficiary',  blank=True, related_name='+')
	customerType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CustomerType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.organizationName
		str = str + self.taxId
		str = str + self.dateOfBirth
		str = str + self.primaryAddress
		str = str + self.customerType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Customer";
    
	def objectType(self):
		return "Customer";
