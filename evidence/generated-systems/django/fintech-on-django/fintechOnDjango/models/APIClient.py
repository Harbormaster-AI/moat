from django.db import models
from fintechOnDjango.models.ClientType import ClientType

#======================================================================
# 
# Encapsulates data for model APIClient
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class APIClient Declaration
#======================================================================
class APIClient (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	clientId = models.CharField(max_length=200, null=True)
	redirectUri = models.CharField(max_length=200, null=True)
	consents = models.ManyToManyField('Consent',  blank=True, related_name='+')
	clientType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ClientType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.clientId
		str = str + self.redirectUri
		str = str + self.clientType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "APIClient";
    
	def objectType(self):
		return "APIClient";
