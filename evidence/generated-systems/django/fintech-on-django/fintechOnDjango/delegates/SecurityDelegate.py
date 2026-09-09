from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Security import Security
from fintechOnDjango.models.Position import Position
from fintechOnDjango.models.Trade import Trade
from fintechOnDjango.models.TradeOrder import TradeOrder
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Security
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SecurityDelegate Declaration
#======================================================================
class SecurityDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, securityId ):
		try:	
			security = Security.objects.filter(id=securityId)
			return security.first();
		except Security.DoesNotExist:
			raise ProcessingError("Security with id " + str(securityId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, security):
		for model in serializers.deserialize("json", security):
			model.save()
			return model;

	def create(self, security):
		security.save()
		return security;

	def saveFromJson(self, security):
		for model in serializers.deserialize("json", security):
			model.save()
			return security;
	
	def save(self, security):
		security.save()
		return security;
	
	def delete(self, securityId ):
		errMsg = "Failed to delete Security from db using id " + str(securityId)
		
		try:
			security = Security.objects.get(id=securityId)
			security.delete()
			return True
		except Security.DoesNotExist:
			raise ProcessingError("Security with id " + str(securityId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Security.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Security from db")
		except Exception:
			return None;
		
	def addPositions( self, securityId, positionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to add elements " + str(positionsIds) + " for Positions on Security"

		try:
			# get the Security
			security = self.get( securityId ).first()
				
			# split on a comma with no spaces
			idList = positionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				security.positions.add(position)
				
			# save it		
			security.save()
			
			# reload and return the appropriate version
			return self.get( securityId );
		except Security.DoesNotExist:
			raise ProcessingError(errMsg + " : Security with id " + str(securityId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePositions( self, securityId, positionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to remove elements " + str(positionsIds) + " for Positions on Security"

		try:
			# get the Security
			security = self.get( securityId ).first()
				
			# split on a comma with no spaces
			idList = positionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				security.positions.remove(position)
				
			# save it		
			security.save()
			
			# reload and return the appropriate version
			return self.get( securityId );
		except Security.DoesNotExist:
			raise ProcessingError(errMsg + " : Security with id " + str(securityId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTrades( self, securityId, tradesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TradeDelegate import TradeDelegate

		errMsg = "Failed to add elements " + str(tradesIds) + " for Trades on Security"

		try:
			# get the Security
			security = self.get( securityId ).first()
				
			# split on a comma with no spaces
			idList = tradesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Trade		
				trade = TradeDelegate().get(id).first();	
				# add the Trade
				security.trades.add(trade)
				
			# save it		
			security.save()
			
			# reload and return the appropriate version
			return self.get( securityId );
		except Security.DoesNotExist:
			raise ProcessingError(errMsg + " : Security with id " + str(securityId) + " does not exist.")
		except Trade.DoesNotExist:
			raise ProcessingError(errMsg + " : Trade does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTrades( self, securityId, tradesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TradeDelegate import TradeDelegate

		errMsg = "Failed to remove elements " + str(tradesIds) + " for Trades on Security"

		try:
			# get the Security
			security = self.get( securityId ).first()
				
			# split on a comma with no spaces
			idList = tradesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Trade		
				trade = TradeDelegate().get(id).first();	
				# add the Trade
				security.trades.remove(trade)
				
			# save it		
			security.save()
			
			# reload and return the appropriate version
			return self.get( securityId );
		except Security.DoesNotExist:
			raise ProcessingError(errMsg + " : Security with id " + str(securityId) + " does not exist.")
		except Trade.DoesNotExist:
			raise ProcessingError(errMsg + " : Trade does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOrders( self, securityId, ordersIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TradeOrderDelegate import TradeOrderDelegate

		errMsg = "Failed to add elements " + str(ordersIds) + " for Orders on Security"

		try:
			# get the Security
			security = self.get( securityId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TradeOrder		
				tradeOrder = TradeOrderDelegate().get(id).first();	
				# add the TradeOrder
				security.orders.add(tradeOrder)
				
			# save it		
			security.save()
			
			# reload and return the appropriate version
			return self.get( securityId );
		except Security.DoesNotExist:
			raise ProcessingError(errMsg + " : Security with id " + str(securityId) + " does not exist.")
		except TradeOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TradeOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrders( self, securityId, ordersIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TradeOrderDelegate import TradeOrderDelegate

		errMsg = "Failed to remove elements " + str(ordersIds) + " for Orders on Security"

		try:
			# get the Security
			security = self.get( securityId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TradeOrder		
				tradeOrder = TradeOrderDelegate().get(id).first();	
				# add the TradeOrder
				security.orders.remove(tradeOrder)
				
			# save it		
			security.save()
			
			# reload and return the appropriate version
			return self.get( securityId );
		except Security.DoesNotExist:
			raise ProcessingError(errMsg + " : Security with id " + str(securityId) + " does not exist.")
		except TradeOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TradeOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
