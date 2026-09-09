from django.db import models
from fintechOnDjango.models.ScreeningType import ScreeningType
from fintechOnDjango.models.ScreeningStatus import ScreeningStatus

#======================================================================
# 
# Encapsulates data for model Screening
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Screening Declaration
#======================================================================
class Screening (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	score = RiskScore
	screenedAt = DateTime
	kycProfile = models.ForeignKey('KYCProfile', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	alerts = models.ManyToManyField('ComplianceAlert',  blank=True, related_name='+')
	screeningType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ScreeningType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ScreeningStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.score
		str = str + self.screenedAt
		str = str + self.screeningType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Screening";
    
	def objectType(self):
		return "Screening";
