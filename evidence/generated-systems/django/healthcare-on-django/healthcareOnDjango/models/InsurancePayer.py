from django.db import models
from healthcareOnDjango.models.PayerType import PayerType

#======================================================================
# 
# Encapsulates data for model InsurancePayer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsurancePayer Declaration
#======================================================================
class InsurancePayer (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	plans = models.ManyToManyField('InsurancePlan',  blank=True, related_name='+')
	claims = models.ManyToManyField('Claim',  blank=True, related_name='+')
	payerType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PayerType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.website
		str = str + self.payerType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InsurancePayer";
    
	def objectType(self):
		return "InsurancePayer";
