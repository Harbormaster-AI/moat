from django.db import models
from healthcareOnDjango.models.SoftwareUpdateType import SoftwareUpdateType

#======================================================================
# 
# Encapsulates data for model SoftwareUpdate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SoftwareUpdate Declaration
#======================================================================
class SoftwareUpdate (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	version = models.CharField(max_length=200, null=True)
	appliedDate = models.CharField(max_length=64, null=True)
	device = models.ForeignKey('MedicalDevice', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	updateType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SoftwareUpdateType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.version
		str = str + self.appliedDate
		str = str + self.updateType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "SoftwareUpdate";
    
	def objectType(self):
		return "SoftwareUpdate";
