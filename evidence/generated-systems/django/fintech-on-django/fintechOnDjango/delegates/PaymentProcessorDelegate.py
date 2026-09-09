from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.PaymentProcessor import PaymentProcessor
from fintechOnDjango.models.FinancialInstitution import FinancialInstitution
from fintechOnDjango.models.PaymentContract import PaymentContract
from fintechOnDjango.models.SettlementBatch import SettlementBatch
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PaymentProcessor
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentProcessorDelegate Declaration
#======================================================================
class PaymentProcessorDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, paymentProcessorId ):
		try:	
			paymentProcessor = PaymentProcessor.objects.filter(id=paymentProcessorId)
			return paymentProcessor.first();
		except PaymentProcessor.DoesNotExist:
			raise ProcessingError("PaymentProcessor with id " + str(paymentProcessorId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, paymentProcessor):
		for model in serializers.deserialize("json", paymentProcessor):
			model.save()
			return model;

	def create(self, paymentProcessor):
		paymentProcessor.save()
		return paymentProcessor;

	def saveFromJson(self, paymentProcessor):
		for model in serializers.deserialize("json", paymentProcessor):
			model.save()
			return paymentProcessor;
	
	def save(self, paymentProcessor):
		paymentProcessor.save()
		return paymentProcessor;
	
	def delete(self, paymentProcessorId ):
		errMsg = "Failed to delete PaymentProcessor from db using id " + str(paymentProcessorId)
		
		try:
			paymentProcessor = PaymentProcessor.objects.get(id=paymentProcessorId)
			paymentProcessor.delete()
			return True
		except PaymentProcessor.DoesNotExist:
			raise ProcessingError("PaymentProcessor with id " + str(paymentProcessorId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PaymentProcessor.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PaymentProcessor from db")
		except Exception:
			return None;
		
	def addInstitutions( self, paymentProcessorId, institutionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.FinancialInstitutionDelegate import FinancialInstitutionDelegate

		errMsg = "Failed to add elements " + str(institutionsIds) + " for Institutions on PaymentProcessor"

		try:
			# get the PaymentProcessor
			paymentProcessor = self.get( paymentProcessorId ).first()
				
			# split on a comma with no spaces
			idList = institutionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the FinancialInstitution		
				financialInstitution = FinancialInstitutionDelegate().get(id).first();	
				# add the FinancialInstitution
				paymentProcessor.institutions.add(financialInstitution)
				
			# save it		
			paymentProcessor.save()
			
			# reload and return the appropriate version
			return self.get( paymentProcessorId );
		except PaymentProcessor.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProcessor with id " + str(paymentProcessorId) + " does not exist.")
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInstitutions( self, paymentProcessorId, institutionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.FinancialInstitutionDelegate import FinancialInstitutionDelegate

		errMsg = "Failed to remove elements " + str(institutionsIds) + " for Institutions on PaymentProcessor"

		try:
			# get the PaymentProcessor
			paymentProcessor = self.get( paymentProcessorId ).first()
				
			# split on a comma with no spaces
			idList = institutionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the FinancialInstitution		
				financialInstitution = FinancialInstitutionDelegate().get(id).first();	
				# add the FinancialInstitution
				paymentProcessor.institutions.remove(financialInstitution)
				
			# save it		
			paymentProcessor.save()
			
			# reload and return the appropriate version
			return self.get( paymentProcessorId );
		except PaymentProcessor.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProcessor with id " + str(paymentProcessorId) + " does not exist.")
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addContracts( self, paymentProcessorId, contractsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentContractDelegate import PaymentContractDelegate

		errMsg = "Failed to add elements " + str(contractsIds) + " for Contracts on PaymentProcessor"

		try:
			# get the PaymentProcessor
			paymentProcessor = self.get( paymentProcessorId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PaymentContract		
				paymentContract = PaymentContractDelegate().get(id).first();	
				# add the PaymentContract
				paymentProcessor.contracts.add(paymentContract)
				
			# save it		
			paymentProcessor.save()
			
			# reload and return the appropriate version
			return self.get( paymentProcessorId );
		except PaymentProcessor.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProcessor with id " + str(paymentProcessorId) + " does not exist.")
		except PaymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentContract does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeContracts( self, paymentProcessorId, contractsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentContractDelegate import PaymentContractDelegate

		errMsg = "Failed to remove elements " + str(contractsIds) + " for Contracts on PaymentProcessor"

		try:
			# get the PaymentProcessor
			paymentProcessor = self.get( paymentProcessorId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PaymentContract		
				paymentContract = PaymentContractDelegate().get(id).first();	
				# add the PaymentContract
				paymentProcessor.contracts.remove(paymentContract)
				
			# save it		
			paymentProcessor.save()
			
			# reload and return the appropriate version
			return self.get( paymentProcessorId );
		except PaymentProcessor.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProcessor with id " + str(paymentProcessorId) + " does not exist.")
		except PaymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentContract does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSettlements( self, paymentProcessorId, settlementsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.SettlementBatchDelegate import SettlementBatchDelegate

		errMsg = "Failed to add elements " + str(settlementsIds) + " for Settlements on PaymentProcessor"

		try:
			# get the PaymentProcessor
			paymentProcessor = self.get( paymentProcessorId ).first()
				
			# split on a comma with no spaces
			idList = settlementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SettlementBatch		
				settlementBatch = SettlementBatchDelegate().get(id).first();	
				# add the SettlementBatch
				paymentProcessor.settlements.add(settlementBatch)
				
			# save it		
			paymentProcessor.save()
			
			# reload and return the appropriate version
			return self.get( paymentProcessorId );
		except PaymentProcessor.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProcessor with id " + str(paymentProcessorId) + " does not exist.")
		except SettlementBatch.DoesNotExist:
			raise ProcessingError(errMsg + " : SettlementBatch does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSettlements( self, paymentProcessorId, settlementsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.SettlementBatchDelegate import SettlementBatchDelegate

		errMsg = "Failed to remove elements " + str(settlementsIds) + " for Settlements on PaymentProcessor"

		try:
			# get the PaymentProcessor
			paymentProcessor = self.get( paymentProcessorId ).first()
				
			# split on a comma with no spaces
			idList = settlementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SettlementBatch		
				settlementBatch = SettlementBatchDelegate().get(id).first();	
				# add the SettlementBatch
				paymentProcessor.settlements.remove(settlementBatch)
				
			# save it		
			paymentProcessor.save()
			
			# reload and return the appropriate version
			return self.get( paymentProcessorId );
		except PaymentProcessor.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProcessor with id " + str(paymentProcessorId) + " does not exist.")
		except SettlementBatch.DoesNotExist:
			raise ProcessingError(errMsg + " : SettlementBatch does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
