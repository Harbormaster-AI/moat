from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.models.Insurer import Insurer
from insuranceOnDjango.models.Customer import Customer
from insuranceOnDjango.models.InsuranceProduct import InsuranceProduct
from insuranceOnDjango.models.Agent import Agent
from insuranceOnDjango.models.PolicyCoverage import PolicyCoverage
from insuranceOnDjango.models.InsuredObject import InsuredObject
from insuranceOnDjango.models.Endorsement import Endorsement
from insuranceOnDjango.models.BillingAccount import BillingAccount
from insuranceOnDjango.models.Beneficiary import Beneficiary
from insuranceOnDjango.models.Claim import Claim
from insuranceOnDjango.models.ReinsuranceAgreement import ReinsuranceAgreement
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Policy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyDelegate Declaration
#======================================================================
class PolicyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, policyId ):
		try:	
			policy = Policy.objects.filter(id=policyId)
			return policy.first();
		except Policy.DoesNotExist:
			raise ProcessingError("Policy with id " + str(policyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, policy):
		for model in serializers.deserialize("json", policy):
			model.save()
			return model;

	def create(self, policy):
		policy.save()
		return policy;

	def saveFromJson(self, policy):
		for model in serializers.deserialize("json", policy):
			model.save()
			return policy;
	
	def save(self, policy):
		policy.save()
		return policy;
	
	def delete(self, policyId ):
		errMsg = "Failed to delete Policy from db using id " + str(policyId)
		
		try:
			policy = Policy.objects.get(id=policyId)
			policy.delete()
			return True
		except Policy.DoesNotExist:
			raise ProcessingError("Policy with id " + str(policyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Policy.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Policy from db")
		except Exception:
			return None;
		
	def assignInsurer( self, policyId, insurerId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsurerDelegate import InsurerDelegate

		errMsg = "Failed to assign element " + str(insurerId) + " for Insurer on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# get the Insurer from db
			insurer = InsurerDelegate().get(insurerId).first();
			
			# assign the Insurer		
			policy.insurer = insurer
			
			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInsurer( self, policyId ):
		errMsg = "Failed to unassign element " + str(insurerId) + " for Insurer on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# assign to None for unassignment
			policy.insurer = None			

			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCustomer( self, policyId, customerId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			policy.customer = customer
			
			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, policyId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# assign to None for unassignment
			policy.customer = None			

			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
		
	def assignProduct( self, policyId, productId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsuranceProductDelegate import InsuranceProductDelegate

		errMsg = "Failed to assign element " + str(productId) + " for Product on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# get the InsuranceProduct from db
			insuranceProduct = InsuranceProductDelegate().get(productId).first();
			
			# assign the Product		
			policy.product = insuranceProduct
			
			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except InsuranceProduct.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuranceProduct with id " + str(productId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProduct( self, policyId ):
		errMsg = "Failed to unassign element " + str(productId) + " for Product on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# assign to None for unassignment
			policy.insuranceProduct = None			

			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAgent( self, policyId, agentId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.AgentDelegate import AgentDelegate

		errMsg = "Failed to assign element " + str(agentId) + " for Agent on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# get the Agent from db
			agent = AgentDelegate().get(agentId).first();
			
			# assign the Agent		
			policy.agent = agent
			
			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Agent.DoesNotExist:
			raise ProcessingError(errMsg + " : Agent with id " + str(agentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAgent( self, policyId ):
		errMsg = "Failed to unassign element " + str(agentId) + " for Agent on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# assign to None for unassignment
			policy.agent = None			

			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
		
	def assignBillingAccount( self, policyId, billingAccountId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.BillingAccountDelegate import BillingAccountDelegate

		errMsg = "Failed to assign element " + str(billingAccountId) + " for BillingAccount on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# get the BillingAccount from db
			billingAccount = BillingAccountDelegate().get(billingAccountId).first();
			
			# assign the BillingAccount		
			policy.billingAccount = billingAccount
			
			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except BillingAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingAccount with id " + str(billingAccountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBillingAccount( self, policyId ):
		errMsg = "Failed to unassign element " + str(billingAccountId) + " for BillingAccount on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# assign to None for unassignment
			policy.billingAccount = None			

			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
		
	def addCoverages( self, policyId, coveragesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyCoverageDelegate import PolicyCoverageDelegate

		errMsg = "Failed to add elements " + str(coveragesIds) + " for Coverages on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = coveragesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PolicyCoverage		
				policyCoverage = PolicyCoverageDelegate().get(id).first();	
				# add the PolicyCoverage
				policy.coverages.add(policyCoverage)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except PolicyCoverage.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyCoverage does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCoverages( self, policyId, coveragesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyCoverageDelegate import PolicyCoverageDelegate

		errMsg = "Failed to remove elements " + str(coveragesIds) + " for Coverages on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = coveragesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PolicyCoverage		
				policyCoverage = PolicyCoverageDelegate().get(id).first();	
				# add the PolicyCoverage
				policy.coverages.remove(policyCoverage)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except PolicyCoverage.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyCoverage does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInsuredObjects( self, policyId, insuredObjectsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsuredObjectDelegate import InsuredObjectDelegate

		errMsg = "Failed to add elements " + str(insuredObjectsIds) + " for InsuredObjects on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = insuredObjectsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InsuredObject		
				insuredObject = InsuredObjectDelegate().get(id).first();	
				# add the InsuredObject
				policy.insuredObjects.add(insuredObject)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except InsuredObject.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuredObject does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInsuredObjects( self, policyId, insuredObjectsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsuredObjectDelegate import InsuredObjectDelegate

		errMsg = "Failed to remove elements " + str(insuredObjectsIds) + " for InsuredObjects on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = insuredObjectsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InsuredObject		
				insuredObject = InsuredObjectDelegate().get(id).first();	
				# add the InsuredObject
				policy.insuredObjects.remove(insuredObject)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except InsuredObject.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuredObject does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEndorsements( self, policyId, endorsementsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.EndorsementDelegate import EndorsementDelegate

		errMsg = "Failed to add elements " + str(endorsementsIds) + " for Endorsements on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = endorsementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Endorsement		
				endorsement = EndorsementDelegate().get(id).first();	
				# add the Endorsement
				policy.endorsements.add(endorsement)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Endorsement.DoesNotExist:
			raise ProcessingError(errMsg + " : Endorsement does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEndorsements( self, policyId, endorsementsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.EndorsementDelegate import EndorsementDelegate

		errMsg = "Failed to remove elements " + str(endorsementsIds) + " for Endorsements on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = endorsementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Endorsement		
				endorsement = EndorsementDelegate().get(id).first();	
				# add the Endorsement
				policy.endorsements.remove(endorsement)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Endorsement.DoesNotExist:
			raise ProcessingError(errMsg + " : Endorsement does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addBeneficiaries( self, policyId, beneficiariesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.BeneficiaryDelegate import BeneficiaryDelegate

		errMsg = "Failed to add elements " + str(beneficiariesIds) + " for Beneficiaries on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = beneficiariesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Beneficiary		
				beneficiary = BeneficiaryDelegate().get(id).first();	
				# add the Beneficiary
				policy.beneficiaries.add(beneficiary)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Beneficiary.DoesNotExist:
			raise ProcessingError(errMsg + " : Beneficiary does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeBeneficiaries( self, policyId, beneficiariesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.BeneficiaryDelegate import BeneficiaryDelegate

		errMsg = "Failed to remove elements " + str(beneficiariesIds) + " for Beneficiaries on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = beneficiariesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Beneficiary		
				beneficiary = BeneficiaryDelegate().get(id).first();	
				# add the Beneficiary
				policy.beneficiaries.remove(beneficiary)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Beneficiary.DoesNotExist:
			raise ProcessingError(errMsg + " : Beneficiary does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addClaims( self, policyId, claimsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to add elements " + str(claimsIds) + " for Claims on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				policy.claims.add(claim)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeClaims( self, policyId, claimsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to remove elements " + str(claimsIds) + " for Claims on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				policy.claims.remove(claim)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReinsuranceAgreements( self, policyId, reinsuranceAgreementsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ReinsuranceAgreementDelegate import ReinsuranceAgreementDelegate

		errMsg = "Failed to add elements " + str(reinsuranceAgreementsIds) + " for ReinsuranceAgreements on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = reinsuranceAgreementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ReinsuranceAgreement		
				reinsuranceAgreement = ReinsuranceAgreementDelegate().get(id).first();	
				# add the ReinsuranceAgreement
				policy.reinsuranceAgreements.add(reinsuranceAgreement)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except ReinsuranceAgreement.DoesNotExist:
			raise ProcessingError(errMsg + " : ReinsuranceAgreement does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReinsuranceAgreements( self, policyId, reinsuranceAgreementsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ReinsuranceAgreementDelegate import ReinsuranceAgreementDelegate

		errMsg = "Failed to remove elements " + str(reinsuranceAgreementsIds) + " for ReinsuranceAgreements on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = reinsuranceAgreementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ReinsuranceAgreement		
				reinsuranceAgreement = ReinsuranceAgreementDelegate().get(id).first();	
				# add the ReinsuranceAgreement
				policy.reinsuranceAgreements.remove(reinsuranceAgreement)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except ReinsuranceAgreement.DoesNotExist:
			raise ProcessingError(errMsg + " : ReinsuranceAgreement does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
