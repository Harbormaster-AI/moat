from django.db import models
from hrOnDjango.models.PerformanceRating import PerformanceRating

#======================================================================
# 
# Encapsulates data for model CompetencyRating
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompetencyRating Declaration
#======================================================================
class CompetencyRating (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	comment = models.CharField(max_length=200, null=True)
	review = models.ForeignKey('PerformanceReview', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	competency = models.ForeignKey('Competency', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	rating = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PerformanceRating])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.comment
		str = str + self.rating
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CompetencyRating";
    
	def objectType(self):
		return "CompetencyRating";
