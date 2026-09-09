from django.db import models
from analyticsOnDjango.models.AccessLevel import AccessLevel
from analyticsOnDjango.models.SubjectType import SubjectType

#======================================================================
# 
# Encapsulates data for model AccessPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccessPolicy Declaration
#======================================================================
class AccessPolicy (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	subjectName = models.CharField(max_length=200, null=True)
	workspace = models.ForeignKey('AnalyticsWorkspace', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	dashboards = models.ManyToManyField('Dashboard',  blank=True, related_name='+')
	reports = models.ManyToManyField('Report',  blank=True, related_name='+')
	models = models.ManyToManyField('Model',  blank=True, related_name='+')
	featureSets = models.ManyToManyField('FeatureSet',  blank=True, related_name='+')
	accessLevel = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AccessLevel])
	subjectType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SubjectType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.subjectName
		str = str + self.accessLevel
		str = str + self.subjectType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AccessPolicy";
    
	def objectType(self):
		return "AccessPolicy";
