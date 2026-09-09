from django.db import models
from insuranceOnDjango.models.PolicyStatus import PolicyStatus
from insuranceOnDjango.models.PaymentPlanType import PaymentPlanType

#======================================================================
# 
# Encapsulates data for model Policy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Policy Declaration
#======================================================================
class Policy (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	policyNumber = PolicyNumber
	effectivePeriod = DateRange
	totalPremium = Money
	insurer = models.ForeignKey('Insurer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	product = models.ForeignKey('InsuranceProduct', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	agent = models.ForeignKey('Agent', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	coverages = models.ManyToManyField('PolicyCoverage',  blank=True, related_name='+')
	insuredObjects = models.ManyToManyField('InsuredObject',  blank=True, related_name='+')
	endorsements = models.ManyToManyField('Endorsement',  blank=True, related_name='+')
	billingAccount = models.ForeignKey('BillingAccount', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	beneficiaries = models.ManyToManyField('Beneficiary',  blank=True, related_name='+')
	claims = models.ManyToManyField('Claim',  blank=True, related_name='+')
	reinsuranceAgreements = models.ManyToManyField('ReinsuranceAgreement',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PolicyStatus])
	paymentPlan = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentPlanType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.policyNumber
		str = str + self.effectivePeriod
		str = str + self.totalPremium
		str = str + self.status
		str = str + self.paymentPlan
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Policy";
    
	def objectType(self):
		return "Policy";
