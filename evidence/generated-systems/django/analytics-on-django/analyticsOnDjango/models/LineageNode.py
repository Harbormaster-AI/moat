from django.db import models
from analyticsOnDjango.models.LineageNodeType import LineageNodeType

#======================================================================
# 
# Encapsulates data for model LineageNode
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LineageNode Declaration
#======================================================================
class LineageNode (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	qualifiedName = models.CharField(max_length=200, null=True)
	workspace = models.ForeignKey('AnalyticsWorkspace', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	inputs = models.ManyToManyField('LineageNode',  blank=True, related_name='+')
	outputs = models.ManyToManyField('LineageNode',  blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	models = models.ManyToManyField('Model',  blank=True, related_name='+')
	pipelines = models.ManyToManyField('DataPipeline',  blank=True, related_name='+')
	dashboards = models.ManyToManyField('Dashboard',  blank=True, related_name='+')
	reports = models.ManyToManyField('Report',  blank=True, related_name='+')
	nodeType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LineageNodeType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.qualifiedName
		str = str + self.nodeType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "LineageNode";
    
	def objectType(self):
		return "LineageNode";
