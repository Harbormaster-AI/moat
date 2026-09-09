from django.db import models
from analyticsOnDjango.models.ReportStatus import ReportStatus

#======================================================================
# 
# Encapsulates data for model Report
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Report Declaration
#======================================================================
class Report (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	audience = models.CharField(max_length=200, null=True)
	workspace = models.ForeignKey('AnalyticsWorkspace', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	visualizations = models.ManyToManyField('Visualization',  blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	semanticModels = models.ManyToManyField('SemanticModel',  blank=True, related_name='+')
	queries = models.ManyToManyField('BIQuery',  blank=True, related_name='+')
	tags = models.ManyToManyField('Tag',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReportStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.audience
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Report";
    
	def objectType(self):
		return "Report";
