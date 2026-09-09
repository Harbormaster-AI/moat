from django.db import models
from aerospaceOnDjango.models.SoftwareLoadType import SoftwareLoadType

#======================================================================
# 
# Encapsulates data for model SoftwareLoad
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SoftwareLoad Declaration
#======================================================================
class SoftwareLoad (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	version = models.CharField(max_length=200, null=True)
	connectedAircraft = models.ForeignKey('ConnectedAircraft', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	avionicsSuite = models.ForeignKey('AvionicsSuite', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	loadType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SoftwareLoadType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.version
		str = str + self.loadType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "SoftwareLoad";
    
	def objectType(self):
		return "SoftwareLoad";
