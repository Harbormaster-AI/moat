from django.db import models
from governanceOnDjango.models.EvidenceType import EvidenceType

#======================================================================
# 
# Encapsulates data for model Evidence
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Evidence Declaration
#======================================================================
class Evidence (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	locationUrl = URL
	receivedDate = models.DateField(null=True)
	controlTest = models.ForeignKey('ControlTest_', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	control = models.ForeignKey('Control', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	obligation = models.ForeignKey('Obligation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workpaper = models.ForeignKey('AuditWorkpaper', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	evidenceType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in EvidenceType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.locationUrl
		str = str + self.receivedDate
		str = str + self.evidenceType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Evidence";
    
	def objectType(self):
		return "Evidence";
