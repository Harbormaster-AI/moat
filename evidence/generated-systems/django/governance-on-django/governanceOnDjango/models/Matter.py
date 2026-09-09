from django.db import models
from governanceOnDjango.models.MatterType import MatterType
from governanceOnDjango.models.MatterStatus import MatterStatus

#======================================================================
# 
# Encapsulates data for model Matter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Matter Declaration
#======================================================================
class Matter (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	matterName = models.CharField(max_length=200, null=True)
	leadCounsel = models.CharField(max_length=200, null=True)
	legalHolds = models.ManyToManyField('LegalHold',  blank=True, related_name='+')
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dataBreaches = models.ManyToManyField('DataBreach',  blank=True, related_name='+')
	contracts = models.ManyToManyField('Contract',  blank=True, related_name='+')
	matterType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in MatterType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in MatterStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.matterName
		str = str + self.leadCounsel
		str = str + self.matterType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Matter";
    
	def objectType(self):
		return "Matter";
