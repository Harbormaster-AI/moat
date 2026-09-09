from django.db import models
from analyticsOnDjango.models.AlertSeverity import AlertSeverity
from analyticsOnDjango.models.AlertStatus import AlertStatus

#======================================================================
# 
# Encapsulates data for model Alert
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Alert Declaration
#======================================================================
class Alert (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	createdAt = models.DateField(null=True)
	metric = models.ForeignKey('Metric', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dashboard = models.ForeignKey('Dashboard', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dataset = models.ForeignKey('DataSet', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	rule = models.ForeignKey('QualityRule', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	anomalies = models.ManyToManyField('Anomaly',  blank=True, related_name='+')
	subscribers = models.ManyToManyField('Subscriber',  blank=True, related_name='+')
	severity = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AlertSeverity])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AlertStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.createdAt
		str = str + self.severity
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Alert";
    
	def objectType(self):
		return "Alert";
