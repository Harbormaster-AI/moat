from django.db import models
from fintechOnDjango.models.AlertSeverity import AlertSeverity
from fintechOnDjango.models.AlertStatus import AlertStatus

#======================================================================
# 
# Encapsulates data for model ComplianceAlert
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComplianceAlert Declaration
#======================================================================
class ComplianceAlert (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	alertCode = models.CharField(max_length=200, null=True)
	raisedAt = DateTime
	notes = models.CharField(max_length=200, null=True)
	screening = models.ForeignKey('Screening', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	transaction = models.ForeignKey('Transaction', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	severity = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AlertSeverity])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AlertStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.alertCode
		str = str + self.raisedAt
		str = str + self.notes
		str = str + self.severity
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ComplianceAlert";
    
	def objectType(self):
		return "ComplianceAlert";
