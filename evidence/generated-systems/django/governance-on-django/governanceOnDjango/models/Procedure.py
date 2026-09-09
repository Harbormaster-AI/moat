from django.db import models
from governanceOnDjango.models.DocumentStatus import DocumentStatus

#======================================================================
# 
# Encapsulates data for model Procedure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Procedure Declaration
#======================================================================
class Procedure (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	versionLabel = models.CharField(max_length=200, null=True)
	policy = models.ForeignKey('Policy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	controls = models.ManyToManyField('Control',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DocumentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.versionLabel
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Procedure";
    
	def objectType(self):
		return "Procedure";
