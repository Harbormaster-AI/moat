from django.db import models
from governanceOnDjango.models.AssessmentType import AssessmentType
from governanceOnDjango.models.AssessmentResult import AssessmentResult

#======================================================================
# 
# Encapsulates data for model ThirdPartyAssessment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ThirdPartyAssessment Declaration
#======================================================================
class ThirdPartyAssessment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	assessmentDate = models.DateField(null=True)
	assessor = models.CharField(max_length=200, null=True)
	thirdParty = models.ForeignKey('ThirdParty', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	issues = models.ManyToManyField('Issue',  blank=True, related_name='+')
	assessmentType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AssessmentType])
	result = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AssessmentResult])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.assessmentDate
		str = str + self.assessor
		str = str + self.assessmentType
		str = str + self.result
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ThirdPartyAssessment";
    
	def objectType(self):
		return "ThirdPartyAssessment";
