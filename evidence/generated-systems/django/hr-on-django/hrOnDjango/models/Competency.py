from django.db import models

#======================================================================
# 
# Encapsulates data for model Competency
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Competency Declaration
#======================================================================
class Competency (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	category = models.CharField(max_length=200, null=True)
	jobProfiles = models.ManyToManyField('JobProfile',  blank=True, related_name='+')
	competencyRatings = models.ManyToManyField('CompetencyRating',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.category
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Competency";
    
	def objectType(self):
		return "Competency";
