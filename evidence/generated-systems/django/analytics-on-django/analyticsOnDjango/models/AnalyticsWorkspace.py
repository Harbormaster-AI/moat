from django.db import models
from analyticsOnDjango.models.GovernanceTier import GovernanceTier

#======================================================================
# 
# Encapsulates data for model AnalyticsWorkspace
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AnalyticsWorkspace Declaration
#======================================================================
class AnalyticsWorkspace (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	businessDomain = models.CharField(max_length=200, null=True)
	ownerTeam = models.CharField(max_length=200, null=True)
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	dataSources = models.ManyToManyField('DataSource',  blank=True, related_name='+')
	pipelines = models.ManyToManyField('DataPipeline',  blank=True, related_name='+')
	dashboards = models.ManyToManyField('Dashboard',  blank=True, related_name='+')
	reports = models.ManyToManyField('Report',  blank=True, related_name='+')
	notebooks = models.ManyToManyField('Notebook',  blank=True, related_name='+')
	models = models.ManyToManyField('Model',  blank=True, related_name='+')
	featureSets = models.ManyToManyField('FeatureSet',  blank=True, related_name='+')
	policies = models.ManyToManyField('AccessPolicy',  blank=True, related_name='+')
	lineageNodes = models.ManyToManyField('LineageNode',  blank=True, related_name='+')
	governanceTier = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in GovernanceTier])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.businessDomain
		str = str + self.ownerTeam
		str = str + self.governanceTier
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AnalyticsWorkspace";
    
	def objectType(self):
		return "AnalyticsWorkspace";
