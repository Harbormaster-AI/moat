from django.db import models
from insuranceOnDjango.models.ProducerStatus import ProducerStatus

#======================================================================
# 
# Encapsulates data for model Agent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Agent Declaration
#======================================================================
class Agent (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	licenseId = models.CharField(max_length=200, null=True)
	distributor = models.ForeignKey('Distributor', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	policies = models.ManyToManyField('Policy',  blank=True, related_name='+')
	customers = models.ManyToManyField('Customer',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ProducerStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.licenseId
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Agent";
    
	def objectType(self):
		return "Agent";
