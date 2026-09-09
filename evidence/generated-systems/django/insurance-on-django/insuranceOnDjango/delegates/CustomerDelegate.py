from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Customer import Customer
from insuranceOnDjango.models.Application import Application
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.models.Claim import Claim
from insuranceOnDjango.models.Agent import Agent
from insuranceOnDjango.models.Beneficiary import Beneficiary
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Customer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CustomerDelegate Declaration
#======================================================================
class CustomerDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, customerId ):
		try:	
			customer = Customer.objects.filter(id=customerId)
			return customer.first();
		except Customer.DoesNotExist:
			raise ProcessingError("Customer with id " + str(customerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, customer):
		for model in serializers.deserialize("json", customer):
			model.save()
			return model;

	def create(self, customer):
		customer.save()
		return customer;

	def saveFromJson(self, customer):
		for model in serializers.deserialize("json", customer):
			model.save()
			return customer;
	
	def save(self, customer):
		customer.save()
		return customer;
	
	def delete(self, customerId ):
		errMsg = "Failed to delete Customer from db using id " + str(customerId)
		
		try:
			customer = Customer.objects.get(id=customerId)
			customer.delete()
			return True
		except Customer.DoesNotExist:
			raise ProcessingError("Customer with id " + str(customerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Customer.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Customer from db")
		except Exception:
			return None;
		
	def addApplications( self, customerId, applicationsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ApplicationDelegate import ApplicationDelegate

		errMsg = "Failed to add elements " + str(applicationsIds) + " for Applications on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = applicationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Application		
				application = ApplicationDelegate().get(id).first();	
				# add the Application
				customer.applications.add(application)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeApplications( self, customerId, applicationsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ApplicationDelegate import ApplicationDelegate

		errMsg = "Failed to remove elements " + str(applicationsIds) + " for Applications on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = applicationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Application		
				application = ApplicationDelegate().get(id).first();	
				# add the Application
				customer.applications.remove(application)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPolicies( self, customerId, policiesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to add elements " + str(policiesIds) + " for Policies on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				customer.policies.add(policy)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePolicies( self, customerId, policiesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to remove elements " + str(policiesIds) + " for Policies on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				customer.policies.remove(policy)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addClaims( self, customerId, claimsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to add elements " + str(claimsIds) + " for Claims on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				customer.claims.add(claim)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeClaims( self, customerId, claimsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to remove elements " + str(claimsIds) + " for Claims on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				customer.claims.remove(claim)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAgents( self, customerId, agentsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.AgentDelegate import AgentDelegate

		errMsg = "Failed to add elements " + str(agentsIds) + " for Agents on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = agentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Agent		
				agent = AgentDelegate().get(id).first();	
				# add the Agent
				customer.agents.add(agent)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Agent.DoesNotExist:
			raise ProcessingError(errMsg + " : Agent does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAgents( self, customerId, agentsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.AgentDelegate import AgentDelegate

		errMsg = "Failed to remove elements " + str(agentsIds) + " for Agents on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = agentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Agent		
				agent = AgentDelegate().get(id).first();	
				# add the Agent
				customer.agents.remove(agent)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Agent.DoesNotExist:
			raise ProcessingError(errMsg + " : Agent does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addBeneficiaries( self, customerId, beneficiariesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.BeneficiaryDelegate import BeneficiaryDelegate

		errMsg = "Failed to add elements " + str(beneficiariesIds) + " for Beneficiaries on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = beneficiariesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Beneficiary		
				beneficiary = BeneficiaryDelegate().get(id).first();	
				# add the Beneficiary
				customer.beneficiaries.add(beneficiary)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Beneficiary.DoesNotExist:
			raise ProcessingError(errMsg + " : Beneficiary does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeBeneficiaries( self, customerId, beneficiariesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.BeneficiaryDelegate import BeneficiaryDelegate

		errMsg = "Failed to remove elements " + str(beneficiariesIds) + " for Beneficiaries on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = beneficiariesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Beneficiary		
				beneficiary = BeneficiaryDelegate().get(id).first();	
				# add the Beneficiary
				customer.beneficiaries.remove(beneficiary)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Beneficiary.DoesNotExist:
			raise ProcessingError(errMsg + " : Beneficiary does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
