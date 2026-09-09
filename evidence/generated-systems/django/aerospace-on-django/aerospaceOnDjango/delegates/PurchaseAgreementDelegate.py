from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.PurchaseAgreement import PurchaseAgreement
from aerospaceOnDjango.models.AircraftOrder import AircraftOrder
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PurchaseAgreement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PurchaseAgreementDelegate Declaration
#======================================================================
class PurchaseAgreementDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, purchaseAgreementId ):
		try:	
			purchaseAgreement = PurchaseAgreement.objects.filter(id=purchaseAgreementId)
			return purchaseAgreement.first();
		except PurchaseAgreement.DoesNotExist:
			raise ProcessingError("PurchaseAgreement with id " + str(purchaseAgreementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, purchaseAgreement):
		for model in serializers.deserialize("json", purchaseAgreement):
			model.save()
			return model;

	def create(self, purchaseAgreement):
		purchaseAgreement.save()
		return purchaseAgreement;

	def saveFromJson(self, purchaseAgreement):
		for model in serializers.deserialize("json", purchaseAgreement):
			model.save()
			return purchaseAgreement;
	
	def save(self, purchaseAgreement):
		purchaseAgreement.save()
		return purchaseAgreement;
	
	def delete(self, purchaseAgreementId ):
		errMsg = "Failed to delete PurchaseAgreement from db using id " + str(purchaseAgreementId)
		
		try:
			purchaseAgreement = PurchaseAgreement.objects.get(id=purchaseAgreementId)
			purchaseAgreement.delete()
			return True
		except PurchaseAgreement.DoesNotExist:
			raise ProcessingError("PurchaseAgreement with id " + str(purchaseAgreementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PurchaseAgreement.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PurchaseAgreement from db")
		except Exception:
			return None;
		
	def assignAircraftOrder( self, purchaseAgreementId, aircraftOrderId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftOrderDelegate import AircraftOrderDelegate

		errMsg = "Failed to assign element " + str(aircraftOrderId) + " for AircraftOrder on PurchaseAgreement"

		try:
			# get the PurchaseAgreement from db
			purchaseAgreement = self.get( purchaseAgreementId ).first()	
			
			# get the AircraftOrder from db
			aircraftOrder = AircraftOrderDelegate().get(aircraftOrderId).first();
			
			# assign the AircraftOrder		
			purchaseAgreement.aircraftOrder = aircraftOrder
			
			#save it
			purchaseAgreement.save()

			# reload and return the appropriate version					
			return self.get( purchaseAgreementId );
		except PurchaseAgreement.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseAgreement with id " + str(purchaseAgreementId) + " does not exist.")
		except AircraftOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOrder with id " + str(aircraftOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAircraftOrder( self, purchaseAgreementId ):
		errMsg = "Failed to unassign element " + str(aircraftOrderId) + " for AircraftOrder on PurchaseAgreement"

		try:
			# get the PurchaseAgreement from db
			purchaseAgreement = self.get( purchaseAgreementId ).first()	
			
			# assign to None for unassignment
			purchaseAgreement.aircraftOrder = None			

			#save it
			purchaseAgreement.save()

			# reload and return the appropriate version					
			return self.get( purchaseAgreementId );
		except PurchaseAgreement.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseAgreement with id " + str(purchaseAgreementId) + " does not exist.")
		except Exception:
			return None;
		
