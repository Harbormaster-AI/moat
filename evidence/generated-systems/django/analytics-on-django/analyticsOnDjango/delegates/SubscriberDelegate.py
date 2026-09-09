from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Subscriber import Subscriber
from analyticsOnDjango.models.Alert import Alert
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Subscriber
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubscriberDelegate Declaration
#======================================================================
class SubscriberDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, subscriberId ):
		try:	
			subscriber = Subscriber.objects.filter(id=subscriberId)
			return subscriber.first();
		except Subscriber.DoesNotExist:
			raise ProcessingError("Subscriber with id " + str(subscriberId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, subscriber):
		for model in serializers.deserialize("json", subscriber):
			model.save()
			return model;

	def create(self, subscriber):
		subscriber.save()
		return subscriber;

	def saveFromJson(self, subscriber):
		for model in serializers.deserialize("json", subscriber):
			model.save()
			return subscriber;
	
	def save(self, subscriber):
		subscriber.save()
		return subscriber;
	
	def delete(self, subscriberId ):
		errMsg = "Failed to delete Subscriber from db using id " + str(subscriberId)
		
		try:
			subscriber = Subscriber.objects.get(id=subscriberId)
			subscriber.delete()
			return True
		except Subscriber.DoesNotExist:
			raise ProcessingError("Subscriber with id " + str(subscriberId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Subscriber.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Subscriber from db")
		except Exception:
			return None;
		
	def addAlerts( self, subscriberId, alertsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AlertDelegate import AlertDelegate

		errMsg = "Failed to add elements " + str(alertsIds) + " for Alerts on Subscriber"

		try:
			# get the Subscriber
			subscriber = self.get( subscriberId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Alert		
				alert = AlertDelegate().get(id).first();	
				# add the Alert
				subscriber.alerts.add(alert)
				
			# save it		
			subscriber.save()
			
			# reload and return the appropriate version
			return self.get( subscriberId );
		except Subscriber.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscriber with id " + str(subscriberId) + " does not exist.")
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAlerts( self, subscriberId, alertsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AlertDelegate import AlertDelegate

		errMsg = "Failed to remove elements " + str(alertsIds) + " for Alerts on Subscriber"

		try:
			# get the Subscriber
			subscriber = self.get( subscriberId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Alert		
				alert = AlertDelegate().get(id).first();	
				# add the Alert
				subscriber.alerts.remove(alert)
				
			# save it		
			subscriber.save()
			
			# reload and return the appropriate version
			return self.get( subscriberId );
		except Subscriber.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscriber with id " + str(subscriberId) + " does not exist.")
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
