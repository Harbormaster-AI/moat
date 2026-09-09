from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.DemandSignal import DemandSignal
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.Reservation import Reservation
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DemandSignal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DemandSignalDelegate Declaration
#======================================================================
class DemandSignalDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, demandSignalId ):
		try:	
			demandSignal = DemandSignal.objects.filter(id=demandSignalId)
			return demandSignal.first();
		except DemandSignal.DoesNotExist:
			raise ProcessingError("DemandSignal with id " + str(demandSignalId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, demandSignal):
		for model in serializers.deserialize("json", demandSignal):
			model.save()
			return model;

	def create(self, demandSignal):
		demandSignal.save()
		return demandSignal;

	def saveFromJson(self, demandSignal):
		for model in serializers.deserialize("json", demandSignal):
			model.save()
			return demandSignal;
	
	def save(self, demandSignal):
		demandSignal.save()
		return demandSignal;
	
	def delete(self, demandSignalId ):
		errMsg = "Failed to delete DemandSignal from db using id " + str(demandSignalId)
		
		try:
			demandSignal = DemandSignal.objects.get(id=demandSignalId)
			demandSignal.delete()
			return True
		except DemandSignal.DoesNotExist:
			raise ProcessingError("DemandSignal with id " + str(demandSignalId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DemandSignal.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DemandSignal from db")
		except Exception:
			return None;
		
	def assignSku( self, demandSignalId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on DemandSignal"

		try:
			# get the DemandSignal from db
			demandSignal = self.get( demandSignalId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			demandSignal.sku = stockKeepingUnit
			
			#save it
			demandSignal.save()

			# reload and return the appropriate version					
			return self.get( demandSignalId );
		except DemandSignal.DoesNotExist:
			raise ProcessingError(errMsg + " : DemandSignal with id " + str(demandSignalId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, demandSignalId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on DemandSignal"

		try:
			# get the DemandSignal from db
			demandSignal = self.get( demandSignalId ).first()	
			
			# assign to None for unassignment
			demandSignal.stockKeepingUnit = None			

			#save it
			demandSignal.save()

			# reload and return the appropriate version					
			return self.get( demandSignalId );
		except DemandSignal.DoesNotExist:
			raise ProcessingError(errMsg + " : DemandSignal with id " + str(demandSignalId) + " does not exist.")
		except Exception:
			return None;
		
	def addReservations( self, demandSignalId, reservationsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.ReservationDelegate import ReservationDelegate

		errMsg = "Failed to add elements " + str(reservationsIds) + " for Reservations on DemandSignal"

		try:
			# get the DemandSignal
			demandSignal = self.get( demandSignalId ).first()
				
			# split on a comma with no spaces
			idList = reservationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Reservation		
				reservation = ReservationDelegate().get(id).first();	
				# add the Reservation
				demandSignal.reservations.add(reservation)
				
			# save it		
			demandSignal.save()
			
			# reload and return the appropriate version
			return self.get( demandSignalId );
		except DemandSignal.DoesNotExist:
			raise ProcessingError(errMsg + " : DemandSignal with id " + str(demandSignalId) + " does not exist.")
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReservations( self, demandSignalId, reservationsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.ReservationDelegate import ReservationDelegate

		errMsg = "Failed to remove elements " + str(reservationsIds) + " for Reservations on DemandSignal"

		try:
			# get the DemandSignal
			demandSignal = self.get( demandSignalId ).first()
				
			# split on a comma with no spaces
			idList = reservationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Reservation		
				reservation = ReservationDelegate().get(id).first();	
				# add the Reservation
				demandSignal.reservations.remove(reservation)
				
			# save it		
			demandSignal.save()
			
			# reload and return the appropriate version
			return self.get( demandSignalId );
		except DemandSignal.DoesNotExist:
			raise ProcessingError(errMsg + " : DemandSignal with id " + str(demandSignalId) + " does not exist.")
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
