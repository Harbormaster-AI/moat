from django.db import models
from insuranceOnDjango.models.ServiceProviderType import ServiceProviderType
from insuranceOnDjango.models.NetworkStatus import NetworkStatus

#======================================================================
# 
# Encapsulates data for model ServiceProvider
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ServiceProvider Declaration
#======================================================================
class ServiceProvider (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	taxId = models.CharField(max_length=200, null=True)
	claims = models.ManyToManyField('Claim',  blank=True, related_name='+')
	providerType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ServiceProviderType])
	networkStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in NetworkStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.taxId
		str = str + self.providerType
		str = str + self.networkStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ServiceProvider";
    
	def objectType(self):
		return "ServiceProvider";
