from django.db import models
from analyticsOnDjango.models.PipelineTriggerType import PipelineTriggerType
from analyticsOnDjango.models.PipelineStatus import PipelineStatus

#======================================================================
# 
# Encapsulates data for model DataPipeline
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataPipeline Declaration
#======================================================================
class DataPipeline (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	schedule = CronSchedule
	workspace = models.ForeignKey('AnalyticsWorkspace', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	tasks = models.ManyToManyField('DataTask',  blank=True, related_name='+')
	sources = models.ManyToManyField('DataSource',  blank=True, related_name='+')
	outputs = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	lineageNode = models.ForeignKey('LineageNode', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	triggerType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PipelineTriggerType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PipelineStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.schedule
		str = str + self.triggerType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DataPipeline";
    
	def objectType(self):
		return "DataPipeline";
