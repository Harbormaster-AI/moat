from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.Placement import Placement
from advertisingOnDjango.models.LineItem import LineItem
from advertisingOnDjango.models.AdSlot import AdSlot
from advertisingOnDjango.models.Deal import Deal
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Placement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlacementDelegate Declaration
#======================================================================
class PlacementDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, placementId ):
		try:	
			placement = Placement.objects.filter(id=placementId)
			return placement.first();
		except Placement.DoesNotExist:
			raise ProcessingError("Placement with id " + str(placementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, placement):
		for model in serializers.deserialize("json", placement):
			model.save()
			return model;

	def create(self, placement):
		placement.save()
		return placement;

	def saveFromJson(self, placement):
		for model in serializers.deserialize("json", placement):
			model.save()
			return placement;
	
	def save(self, placement):
		placement.save()
		return placement;
	
	def delete(self, placementId ):
		errMsg = "Failed to delete Placement from db using id " + str(placementId)
		
		try:
			placement = Placement.objects.get(id=placementId)
			placement.delete()
			return True
		except Placement.DoesNotExist:
			raise ProcessingError("Placement with id " + str(placementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Placement.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Placement from db")
		except Exception:
			return None;
		
	def assignLineItem( self, placementId, lineItemId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.LineItemDelegate import LineItemDelegate

		errMsg = "Failed to assign element " + str(lineItemId) + " for LineItem on Placement"

		try:
			# get the Placement from db
			placement = self.get( placementId ).first()	
			
			# get the LineItem from db
			lineItem = LineItemDelegate().get(lineItemId).first();
			
			# assign the LineItem		
			placement.lineItem = lineItem
			
			#save it
			placement.save()

			# reload and return the appropriate version					
			return self.get( placementId );
		except Placement.DoesNotExist:
			raise ProcessingError(errMsg + " : Placement with id " + str(placementId) + " does not exist.")
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLineItem( self, placementId ):
		errMsg = "Failed to unassign element " + str(lineItemId) + " for LineItem on Placement"

		try:
			# get the Placement from db
			placement = self.get( placementId ).first()	
			
			# assign to None for unassignment
			placement.lineItem = None			

			#save it
			placement.save()

			# reload and return the appropriate version					
			return self.get( placementId );
		except Placement.DoesNotExist:
			raise ProcessingError(errMsg + " : Placement with id " + str(placementId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAdSlot( self, placementId, adSlotId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdSlotDelegate import AdSlotDelegate

		errMsg = "Failed to assign element " + str(adSlotId) + " for AdSlot on Placement"

		try:
			# get the Placement from db
			placement = self.get( placementId ).first()	
			
			# get the AdSlot from db
			adSlot = AdSlotDelegate().get(adSlotId).first();
			
			# assign the AdSlot		
			placement.adSlot = adSlot
			
			#save it
			placement.save()

			# reload and return the appropriate version					
			return self.get( placementId );
		except Placement.DoesNotExist:
			raise ProcessingError(errMsg + " : Placement with id " + str(placementId) + " does not exist.")
		except AdSlot.DoesNotExist:
			raise ProcessingError(errMsg + " : AdSlot with id " + str(adSlotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAdSlot( self, placementId ):
		errMsg = "Failed to unassign element " + str(adSlotId) + " for AdSlot on Placement"

		try:
			# get the Placement from db
			placement = self.get( placementId ).first()	
			
			# assign to None for unassignment
			placement.adSlot = None			

			#save it
			placement.save()

			# reload and return the appropriate version					
			return self.get( placementId );
		except Placement.DoesNotExist:
			raise ProcessingError(errMsg + " : Placement with id " + str(placementId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDeal( self, placementId, dealId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.DealDelegate import DealDelegate

		errMsg = "Failed to assign element " + str(dealId) + " for Deal on Placement"

		try:
			# get the Placement from db
			placement = self.get( placementId ).first()	
			
			# get the Deal from db
			deal = DealDelegate().get(dealId).first();
			
			# assign the Deal		
			placement.deal = deal
			
			#save it
			placement.save()

			# reload and return the appropriate version					
			return self.get( placementId );
		except Placement.DoesNotExist:
			raise ProcessingError(errMsg + " : Placement with id " + str(placementId) + " does not exist.")
		except Deal.DoesNotExist:
			raise ProcessingError(errMsg + " : Deal with id " + str(dealId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDeal( self, placementId ):
		errMsg = "Failed to unassign element " + str(dealId) + " for Deal on Placement"

		try:
			# get the Placement from db
			placement = self.get( placementId ).first()	
			
			# assign to None for unassignment
			placement.deal = None			

			#save it
			placement.save()

			# reload and return the appropriate version					
			return self.get( placementId );
		except Placement.DoesNotExist:
			raise ProcessingError(errMsg + " : Placement with id " + str(placementId) + " does not exist.")
		except Exception:
			return None;
		
