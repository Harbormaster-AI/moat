from django.db import models
from governanceOnDjango.models.AssessmentType import AssessmentType

#======================================================================
# 
# Encapsulates data for model RiskAssessment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RiskAssessment Declaration
#======================================================================
class RiskAssessment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	assessmentDate = models.DateField(null=True)
	assessor = models.CharField(max_length=200, null=True)
	summary = models.CharField(max_length=200, null=True)
	risk = models.ForeignKey('Risk', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	assessmentType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AssessmentType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.assessmentDate
		str = str + self.assessor
		str = str + self.summary
		str = str + self.assessmentType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "RiskAssessment";
    
	def objectType(self):
		return "RiskAssessment";
