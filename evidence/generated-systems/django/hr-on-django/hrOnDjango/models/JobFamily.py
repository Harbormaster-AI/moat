from django.db import models

#======================================================================
# 
# Encapsulates data for model JobFamily
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobFamily Declaration
#======================================================================
class JobFamily (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	description = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	jobProfiles = models.ManyToManyField('JobProfile',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.description
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "JobFamily";
    
	def objectType(self):
		return "JobFamily";
