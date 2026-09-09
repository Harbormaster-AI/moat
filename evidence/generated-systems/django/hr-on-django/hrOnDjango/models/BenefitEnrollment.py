from django.db import models
from hrOnDjango.models.BenefitEnrollmentStatus import BenefitEnrollmentStatus
from hrOnDjango.models.CoverageLevel import CoverageLevel

#======================================================================
# 
# Encapsulates data for model BenefitEnrollment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BenefitEnrollment Declaration
#======================================================================
class BenefitEnrollment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	enrollmentId = models.CharField(max_length=200, null=True)
	effectiveFrom = models.DateField(null=True)
	effectiveTo = models.DateField(null=True)
	benefitPlan = models.ForeignKey('BenefitPlan', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dependents = models.ManyToManyField('Dependent',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in BenefitEnrollmentStatus])
	coverageLevel = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CoverageLevel])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.enrollmentId
		str = str + self.effectiveFrom
		str = str + self.effectiveTo
		str = str + self.status
		str = str + self.coverageLevel
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BenefitEnrollment";
    
	def objectType(self):
		return "BenefitEnrollment";
