from django.db import models
from governanceOnDjango.models.SystemType import SystemType

#======================================================================
# 
# Encapsulates data for model System_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class System_ Declaration
#======================================================================
class System_ (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	ownerDepartment = models.CharField(max_length=200, null=True)
	processingActivities = models.ManyToManyField('DataProcessingActivity',  blank=True, related_name='+')
	recordsRepositories = models.ManyToManyField('RecordsRepository',  blank=True, related_name='+')
	systemType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SystemType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.ownerDepartment
		str = str + self.systemType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "System_";
    
	def objectType(self):
		return "System_";
