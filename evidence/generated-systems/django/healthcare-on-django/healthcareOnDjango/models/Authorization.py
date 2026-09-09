from django.db import models
from healthcareOnDjango.models.AuthorizationStatus import AuthorizationStatus

#======================================================================
# 
# Encapsulates data for model Authorization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Authorization Declaration
#======================================================================
class Authorization (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	authNumber = models.CharField(max_length=200, null=True)
	requestedService = models.CharField(max_length=200, null=True)
	coverage = models.ForeignKey('Coverage', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	order = models.ForeignKey('ClinicalOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AuthorizationStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.authNumber
		str = str + self.requestedService
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Authorization";
    
	def objectType(self):
		return "Authorization";
