from django.db import models
from analyticsOnDjango.models.AnomalyType import AnomalyType
from analyticsOnDjango.models.AlertSeverity import AlertSeverity

#======================================================================
# 
# Encapsulates data for model Anomaly
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Anomaly Declaration
#======================================================================
class Anomaly (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	occurredAt = models.DateField(null=True)
	details = models.CharField(max_length=200, null=True)
	timeSeries = models.ForeignKey('TimeSeries', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	alert = models.ForeignKey('Alert', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dataset = models.ForeignKey('DataSet', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	anomalyType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AnomalyType])
	severity = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AlertSeverity])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.occurredAt
		str = str + self.details
		str = str + self.anomalyType
		str = str + self.severity
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Anomaly";
    
	def objectType(self):
		return "Anomaly";
