from django.db import models
from insuranceOnDjango.models.DistributionChannelType import DistributionChannelType

#======================================================================
# 
# Encapsulates data for model Distributor
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Distributor Declaration
#======================================================================
class Distributor (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	licenseNumber = models.CharField(max_length=200, null=True)
	region = models.CharField(max_length=200, null=True)
	insurers = models.ManyToManyField('Insurer',  blank=True, related_name='+')
	agents = models.ManyToManyField('Agent',  blank=True, related_name='+')
	policies = models.ManyToManyField('Policy',  blank=True, related_name='+')
	distributorType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DistributionChannelType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.licenseNumber
		str = str + self.region
		str = str + self.distributorType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Distributor";
    
	def objectType(self):
		return "Distributor";
