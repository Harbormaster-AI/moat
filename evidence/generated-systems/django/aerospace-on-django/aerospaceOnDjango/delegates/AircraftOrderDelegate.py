from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.AircraftOrder import AircraftOrder
from aerospaceOnDjango.models.Operator import Operator
from aerospaceOnDjango.models.AircraftVariant import AircraftVariant
from aerospaceOnDjango.models.Quote import Quote
from aerospaceOnDjango.models.PurchaseAgreement import PurchaseAgreement
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AircraftOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftOrderDelegate Declaration
#======================================================================
class AircraftOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, aircraftOrderId ):
		try:	
			aircraftOrder = AircraftOrder.objects.filter(id=aircraftOrderId)
			return aircraftOrder.first();
		except AircraftOrder.DoesNotExist:
			raise ProcessingError("AircraftOrder with id " + str(aircraftOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, aircraftOrder):
		for model in serializers.deserialize("json", aircraftOrder):
			model.save()
			return model;

	def create(self, aircraftOrder):
		aircraftOrder.save()
		return aircraftOrder;

	def saveFromJson(self, aircraftOrder):
		for model in serializers.deserialize("json", aircraftOrder):
			model.save()
			return aircraftOrder;
	
	def save(self, aircraftOrder):
		aircraftOrder.save()
		return aircraftOrder;
	
	def delete(self, aircraftOrderId ):
		errMsg = "Failed to delete AircraftOrder from db using id " + str(aircraftOrderId)
		
		try:
			aircraftOrder = AircraftOrder.objects.get(id=aircraftOrderId)
			aircraftOrder.delete()
			return True
		except AircraftOrder.DoesNotExist:
			raise ProcessingError("AircraftOrder with id " + str(aircraftOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AircraftOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AircraftOrder from db")
		except Exception:
			return None;
		
	def assignOperator( self, aircraftOrderId, operatorId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.OperatorDelegate import OperatorDelegate

		errMsg = "Failed to assign element " + str(operatorId) + " for Operator on AircraftOrder"

		try:
			# get the AircraftOrder from db
			aircraftOrder = self.get( aircraftOrderId ).first()	
			
			# get the Operator from db
			operator = OperatorDelegate().get(operatorId).first();
			
			# assign the Operator		
			aircraftOrder.operator = operator
			
			#save it
			aircraftOrder.save()

			# reload and return the appropriate version					
			return self.get( aircraftOrderId );
		except AircraftOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOrder with id " + str(aircraftOrderId) + " does not exist.")
		except Operator.DoesNotExist:
			raise ProcessingError(errMsg + " : Operator with id " + str(operatorId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOperator( self, aircraftOrderId ):
		errMsg = "Failed to unassign element " + str(operatorId) + " for Operator on AircraftOrder"

		try:
			# get the AircraftOrder from db
			aircraftOrder = self.get( aircraftOrderId ).first()	
			
			# assign to None for unassignment
			aircraftOrder.operator = None			

			#save it
			aircraftOrder.save()

			# reload and return the appropriate version					
			return self.get( aircraftOrderId );
		except AircraftOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOrder with id " + str(aircraftOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignVariant( self, aircraftOrderId, variantId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to assign element " + str(variantId) + " for Variant on AircraftOrder"

		try:
			# get the AircraftOrder from db
			aircraftOrder = self.get( aircraftOrderId ).first()	
			
			# get the AircraftVariant from db
			aircraftVariant = AircraftVariantDelegate().get(variantId).first();
			
			# assign the Variant		
			aircraftOrder.variant = aircraftVariant
			
			#save it
			aircraftOrder.save()

			# reload and return the appropriate version					
			return self.get( aircraftOrderId );
		except AircraftOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOrder with id " + str(aircraftOrderId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(variantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignVariant( self, aircraftOrderId ):
		errMsg = "Failed to unassign element " + str(variantId) + " for Variant on AircraftOrder"

		try:
			# get the AircraftOrder from db
			aircraftOrder = self.get( aircraftOrderId ).first()	
			
			# assign to None for unassignment
			aircraftOrder.aircraftVariant = None			

			#save it
			aircraftOrder.save()

			# reload and return the appropriate version					
			return self.get( aircraftOrderId );
		except AircraftOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOrder with id " + str(aircraftOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignQuote( self, aircraftOrderId, quoteId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to assign element " + str(quoteId) + " for Quote on AircraftOrder"

		try:
			# get the AircraftOrder from db
			aircraftOrder = self.get( aircraftOrderId ).first()	
			
			# get the Quote from db
			quote = QuoteDelegate().get(quoteId).first();
			
			# assign the Quote		
			aircraftOrder.quote = quote
			
			#save it
			aircraftOrder.save()

			# reload and return the appropriate version					
			return self.get( aircraftOrderId );
		except AircraftOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOrder with id " + str(aircraftOrderId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignQuote( self, aircraftOrderId ):
		errMsg = "Failed to unassign element " + str(quoteId) + " for Quote on AircraftOrder"

		try:
			# get the AircraftOrder from db
			aircraftOrder = self.get( aircraftOrderId ).first()	
			
			# assign to None for unassignment
			aircraftOrder.quote = None			

			#save it
			aircraftOrder.save()

			# reload and return the appropriate version					
			return self.get( aircraftOrderId );
		except AircraftOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOrder with id " + str(aircraftOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPurchaseAgreement( self, aircraftOrderId, purchaseAgreementId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.PurchaseAgreementDelegate import PurchaseAgreementDelegate

		errMsg = "Failed to assign element " + str(purchaseAgreementId) + " for PurchaseAgreement on AircraftOrder"

		try:
			# get the AircraftOrder from db
			aircraftOrder = self.get( aircraftOrderId ).first()	
			
			# get the PurchaseAgreement from db
			purchaseAgreement = PurchaseAgreementDelegate().get(purchaseAgreementId).first();
			
			# assign the PurchaseAgreement		
			aircraftOrder.purchaseAgreement = purchaseAgreement
			
			#save it
			aircraftOrder.save()

			# reload and return the appropriate version					
			return self.get( aircraftOrderId );
		except AircraftOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOrder with id " + str(aircraftOrderId) + " does not exist.")
		except PurchaseAgreement.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseAgreement with id " + str(purchaseAgreementId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPurchaseAgreement( self, aircraftOrderId ):
		errMsg = "Failed to unassign element " + str(purchaseAgreementId) + " for PurchaseAgreement on AircraftOrder"

		try:
			# get the AircraftOrder from db
			aircraftOrder = self.get( aircraftOrderId ).first()	
			
			# assign to None for unassignment
			aircraftOrder.purchaseAgreement = None			

			#save it
			aircraftOrder.save()

			# reload and return the appropriate version					
			return self.get( aircraftOrderId );
		except AircraftOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOrder with id " + str(aircraftOrderId) + " does not exist.")
		except Exception:
			return None;
		
