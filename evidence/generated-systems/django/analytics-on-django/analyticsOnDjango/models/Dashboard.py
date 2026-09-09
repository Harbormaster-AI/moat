from django.db import models
from analyticsOnDjango.models.DashboardStatus import DashboardStatus

#======================================================================
# 
# Encapsulates data for model Dashboard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Dashboard Declaration
#======================================================================
class Dashboard (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	theme = models.CharField(max_length=200, null=True)
	workspace = models.ForeignKey('AnalyticsWorkspace', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	visualizations = models.ManyToManyField('Visualization',  blank=True, related_name='+')
	reports = models.ManyToManyField('Report',  blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	alerts = models.ManyToManyField('Alert',  blank=True, related_name='+')
	queries = models.ManyToManyField('BIQuery',  blank=True, related_name='+')
	tags = models.ManyToManyField('Tag',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DashboardStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.theme
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Dashboard";
    
	def objectType(self):
		return "Dashboard";
