from django.db import models
from hrOnDjango.models.PerformanceRating import PerformanceRating
from hrOnDjango.models.ReviewStatus import ReviewStatus

#======================================================================
# 
# Encapsulates data for model PerformanceReview
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerformanceReview Declaration
#======================================================================
class PerformanceReview (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	reviewNumber = models.CharField(max_length=200, null=True)
	reviewDate = models.DateField(null=True)
	reviewerComments = models.CharField(max_length=200, null=True)
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	reviewer = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	cycle = models.ForeignKey('PerformanceCycle', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	competencyRatings = models.ManyToManyField('CompetencyRating',  blank=True, related_name='+')
	goals = models.ManyToManyField('Goal',  blank=True, related_name='+')
	rating = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PerformanceRating])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReviewStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.reviewNumber
		str = str + self.reviewDate
		str = str + self.reviewerComments
		str = str + self.rating
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PerformanceReview";
    
	def objectType(self):
		return "PerformanceReview";
