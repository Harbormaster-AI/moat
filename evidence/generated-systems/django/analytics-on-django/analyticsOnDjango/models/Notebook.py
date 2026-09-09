from django.db import models
from analyticsOnDjango.models.NotebookLanguage import NotebookLanguage

#======================================================================
# 
# Encapsulates data for model Notebook
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Notebook Declaration
#======================================================================
class Notebook (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	repository = RepositoryRef
	workspace = models.ForeignKey('AnalyticsWorkspace', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	experiments = models.ManyToManyField('Experiment',  blank=True, related_name='+')
	queries = models.ManyToManyField('BIQuery',  blank=True, related_name='+')
	language = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in NotebookLanguage])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.repository
		str = str + self.language
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Notebook";
    
	def objectType(self):
		return "Notebook";
