from django.db import models
from healthcareOnDjango.models.CareSettingType import CareSettingType

#======================================================================
# 
# Encapsulates data for model CareTeam
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CareTeam Declaration
#======================================================================
class CareTeam (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	department = models.ForeignKey('Department', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	clinicians = models.ManyToManyField('Clinician',  blank=True, related_name='+')
	patients = models.ManyToManyField('Patient',  blank=True, related_name='+')
	careSetting = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CareSettingType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.careSetting
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CareTeam";
    
	def objectType(self):
		return "CareTeam";
