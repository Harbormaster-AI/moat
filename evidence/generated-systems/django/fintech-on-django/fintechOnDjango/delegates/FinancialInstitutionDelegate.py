from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.FinancialInstitution import FinancialInstitution
from fintechOnDjango.models.Branch import Branch
from fintechOnDjango.models.Customer import Customer
from fintechOnDjango.models.ProductOffering import ProductOffering
from fintechOnDjango.models.PaymentProcessor import PaymentProcessor
from fintechOnDjango.models.CompliancePolicy import CompliancePolicy
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model FinancialInstitution
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FinancialInstitutionDelegate Declaration
#======================================================================
class FinancialInstitutionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, financialInstitutionId ):
		try:	
			financialInstitution = FinancialInstitution.objects.filter(id=financialInstitutionId)
			return financialInstitution.first();
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError("FinancialInstitution with id " + str(financialInstitutionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, financialInstitution):
		for model in serializers.deserialize("json", financialInstitution):
			model.save()
			return model;

	def create(self, financialInstitution):
		financialInstitution.save()
		return financialInstitution;

	def saveFromJson(self, financialInstitution):
		for model in serializers.deserialize("json", financialInstitution):
			model.save()
			return financialInstitution;
	
	def save(self, financialInstitution):
		financialInstitution.save()
		return financialInstitution;
	
	def delete(self, financialInstitutionId ):
		errMsg = "Failed to delete FinancialInstitution from db using id " + str(financialInstitutionId)
		
		try:
			financialInstitution = FinancialInstitution.objects.get(id=financialInstitutionId)
			financialInstitution.delete()
			return True
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError("FinancialInstitution with id " + str(financialInstitutionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = FinancialInstitution.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all FinancialInstitution from db")
		except Exception:
			return None;
		
	def addBranches( self, financialInstitutionId, branchesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.BranchDelegate import BranchDelegate

		errMsg = "Failed to add elements " + str(branchesIds) + " for Branches on FinancialInstitution"

		try:
			# get the FinancialInstitution
			financialInstitution = self.get( financialInstitutionId ).first()
				
			# split on a comma with no spaces
			idList = branchesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Branch		
				branch = BranchDelegate().get(id).first();	
				# add the Branch
				financialInstitution.branches.add(branch)
				
			# save it		
			financialInstitution.save()
			
			# reload and return the appropriate version
			return self.get( financialInstitutionId );
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(financialInstitutionId) + " does not exist.")
		except Branch.DoesNotExist:
			raise ProcessingError(errMsg + " : Branch does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeBranches( self, financialInstitutionId, branchesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.BranchDelegate import BranchDelegate

		errMsg = "Failed to remove elements " + str(branchesIds) + " for Branches on FinancialInstitution"

		try:
			# get the FinancialInstitution
			financialInstitution = self.get( financialInstitutionId ).first()
				
			# split on a comma with no spaces
			idList = branchesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Branch		
				branch = BranchDelegate().get(id).first();	
				# add the Branch
				financialInstitution.branches.remove(branch)
				
			# save it		
			financialInstitution.save()
			
			# reload and return the appropriate version
			return self.get( financialInstitutionId );
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(financialInstitutionId) + " does not exist.")
		except Branch.DoesNotExist:
			raise ProcessingError(errMsg + " : Branch does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCustomers( self, financialInstitutionId, customersIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to add elements " + str(customersIds) + " for Customers on FinancialInstitution"

		try:
			# get the FinancialInstitution
			financialInstitution = self.get( financialInstitutionId ).first()
				
			# split on a comma with no spaces
			idList = customersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Customer		
				customer = CustomerDelegate().get(id).first();	
				# add the Customer
				financialInstitution.customers.add(customer)
				
			# save it		
			financialInstitution.save()
			
			# reload and return the appropriate version
			return self.get( financialInstitutionId );
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(financialInstitutionId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCustomers( self, financialInstitutionId, customersIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to remove elements " + str(customersIds) + " for Customers on FinancialInstitution"

		try:
			# get the FinancialInstitution
			financialInstitution = self.get( financialInstitutionId ).first()
				
			# split on a comma with no spaces
			idList = customersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Customer		
				customer = CustomerDelegate().get(id).first();	
				# add the Customer
				financialInstitution.customers.remove(customer)
				
			# save it		
			financialInstitution.save()
			
			# reload and return the appropriate version
			return self.get( financialInstitutionId );
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(financialInstitutionId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addProductOfferings( self, financialInstitutionId, productOfferingsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ProductOfferingDelegate import ProductOfferingDelegate

		errMsg = "Failed to add elements " + str(productOfferingsIds) + " for ProductOfferings on FinancialInstitution"

		try:
			# get the FinancialInstitution
			financialInstitution = self.get( financialInstitutionId ).first()
				
			# split on a comma with no spaces
			idList = productOfferingsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ProductOffering		
				productOffering = ProductOfferingDelegate().get(id).first();	
				# add the ProductOffering
				financialInstitution.productOfferings.add(productOffering)
				
			# save it		
			financialInstitution.save()
			
			# reload and return the appropriate version
			return self.get( financialInstitutionId );
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(financialInstitutionId) + " does not exist.")
		except ProductOffering.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductOffering does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProductOfferings( self, financialInstitutionId, productOfferingsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ProductOfferingDelegate import ProductOfferingDelegate

		errMsg = "Failed to remove elements " + str(productOfferingsIds) + " for ProductOfferings on FinancialInstitution"

		try:
			# get the FinancialInstitution
			financialInstitution = self.get( financialInstitutionId ).first()
				
			# split on a comma with no spaces
			idList = productOfferingsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ProductOffering		
				productOffering = ProductOfferingDelegate().get(id).first();	
				# add the ProductOffering
				financialInstitution.productOfferings.remove(productOffering)
				
			# save it		
			financialInstitution.save()
			
			# reload and return the appropriate version
			return self.get( financialInstitutionId );
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(financialInstitutionId) + " does not exist.")
		except ProductOffering.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductOffering does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPaymentProcessors( self, financialInstitutionId, paymentProcessorsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentProcessorDelegate import PaymentProcessorDelegate

		errMsg = "Failed to add elements " + str(paymentProcessorsIds) + " for PaymentProcessors on FinancialInstitution"

		try:
			# get the FinancialInstitution
			financialInstitution = self.get( financialInstitutionId ).first()
				
			# split on a comma with no spaces
			idList = paymentProcessorsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PaymentProcessor		
				paymentProcessor = PaymentProcessorDelegate().get(id).first();	
				# add the PaymentProcessor
				financialInstitution.paymentProcessors.add(paymentProcessor)
				
			# save it		
			financialInstitution.save()
			
			# reload and return the appropriate version
			return self.get( financialInstitutionId );
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(financialInstitutionId) + " does not exist.")
		except PaymentProcessor.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProcessor does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePaymentProcessors( self, financialInstitutionId, paymentProcessorsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentProcessorDelegate import PaymentProcessorDelegate

		errMsg = "Failed to remove elements " + str(paymentProcessorsIds) + " for PaymentProcessors on FinancialInstitution"

		try:
			# get the FinancialInstitution
			financialInstitution = self.get( financialInstitutionId ).first()
				
			# split on a comma with no spaces
			idList = paymentProcessorsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PaymentProcessor		
				paymentProcessor = PaymentProcessorDelegate().get(id).first();	
				# add the PaymentProcessor
				financialInstitution.paymentProcessors.remove(paymentProcessor)
				
			# save it		
			financialInstitution.save()
			
			# reload and return the appropriate version
			return self.get( financialInstitutionId );
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(financialInstitutionId) + " does not exist.")
		except PaymentProcessor.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProcessor does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCompliancePolicies( self, financialInstitutionId, compliancePoliciesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CompliancePolicyDelegate import CompliancePolicyDelegate

		errMsg = "Failed to add elements " + str(compliancePoliciesIds) + " for CompliancePolicies on FinancialInstitution"

		try:
			# get the FinancialInstitution
			financialInstitution = self.get( financialInstitutionId ).first()
				
			# split on a comma with no spaces
			idList = compliancePoliciesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CompliancePolicy		
				compliancePolicy = CompliancePolicyDelegate().get(id).first();	
				# add the CompliancePolicy
				financialInstitution.compliancePolicies.add(compliancePolicy)
				
			# save it		
			financialInstitution.save()
			
			# reload and return the appropriate version
			return self.get( financialInstitutionId );
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(financialInstitutionId) + " does not exist.")
		except CompliancePolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : CompliancePolicy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCompliancePolicies( self, financialInstitutionId, compliancePoliciesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CompliancePolicyDelegate import CompliancePolicyDelegate

		errMsg = "Failed to remove elements " + str(compliancePoliciesIds) + " for CompliancePolicies on FinancialInstitution"

		try:
			# get the FinancialInstitution
			financialInstitution = self.get( financialInstitutionId ).first()
				
			# split on a comma with no spaces
			idList = compliancePoliciesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CompliancePolicy		
				compliancePolicy = CompliancePolicyDelegate().get(id).first();	
				# add the CompliancePolicy
				financialInstitution.compliancePolicies.remove(compliancePolicy)
				
			# save it		
			financialInstitution.save()
			
			# reload and return the appropriate version
			return self.get( financialInstitutionId );
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(financialInstitutionId) + " does not exist.")
		except CompliancePolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : CompliancePolicy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
