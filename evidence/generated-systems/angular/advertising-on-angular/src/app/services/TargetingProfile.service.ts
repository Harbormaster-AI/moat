import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {TargetingProfile} from '../models/TargetingProfile';
import {AudienceSegmentService} from '../services/AudienceSegment.service';
import {GeoRegionService} from '../services/GeoRegion.service';
import {ContentCategoryService} from '../services/ContentCategory.service';
import {BrandSafetyPolicyService} from '../services/BrandSafetyPolicy.service';
import {DeviceCriterionService} from '../services/DeviceCriterion.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TargetingProfileService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	targetingProfile : TargetingProfile;

	//********************************************************************
	// Catch all for the return value of a service call
	//********************************************************************
	result: any;

	//********************************************************************
	// sole constructor, injected with the HttpClient
	//********************************************************************
	constructor(private http: HttpClient) {
		super();
	}

		//********************************************************************
	// add a TargetingProfile
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTargetingProfile(name, AudienceSegments, GeoRegions, ContentCategories, BrandSafetyPolicy, DeviceCriteria) : Observable<any> {
		const uri_ = this.apiUrl + '/TargetingProfile/create';
		const obj = {
			      		name: name,
      		AudienceSegments: AudienceSegments != null && AudienceSegments.length > 0 ? AudienceSegments : null,
      		GeoRegions: GeoRegions != null && GeoRegions.length > 0 ? GeoRegions : null,
      		ContentCategories: ContentCategories != null && ContentCategories.length > 0 ? ContentCategories : null,
      		BrandSafetyPolicy: BrandSafetyPolicy != null && BrandSafetyPolicy.length > 0 ? BrandSafetyPolicy : null,
			DeviceCriteria: DeviceCriteria != null && DeviceCriteria.length > 0 ? DeviceCriteria : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a TargetingProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTargetingProfile(name, AudienceSegments, GeoRegions, ContentCategories, BrandSafetyPolicy, DeviceCriteria, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/TargetingProfile/update/' + id;
		const obj = {
				      		name: name,
      		AudienceSegments: AudienceSegments != null && AudienceSegments.length > 0 ? AudienceSegments : null,
      		GeoRegions: GeoRegions != null && GeoRegions.length > 0 ? GeoRegions : null,
      		ContentCategories: ContentCategories != null && ContentCategories.length > 0 ? ContentCategories : null,
      		BrandSafetyPolicy: BrandSafetyPolicy != null && BrandSafetyPolicy.length > 0 ? BrandSafetyPolicy : null,
			DeviceCriteria: DeviceCriteria != null && DeviceCriteria.length > 0 ? DeviceCriteria : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a TargetingProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTargetingProfile(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/TargetingProfile/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a TargetingProfile
	// returns the results untouched as an Observable TargetingProfile
	// TargetingProfile model
	// delegates via URI
	//********************************************************************
	getTargetingProfile(id) : Observable<TargetingProfile> {
		const uri_ = this.apiUrl + '/TargetingProfile/load/' + id;

		return this.http.get<TargetingProfile>(uri_);
	}
	
	//********************************************************************
	// gets all TargetingProfile
	// returns the results untouched as JSON representation of an
	// Observable array of TargetingProfile models
	// delegates via URI
	//********************************************************************
	getTargetingProfiles() : Observable<TargetingProfile[]> {
		const uri_ = this.apiUrl + '/TargetingProfile/';

		return this
			.http.get<TargetingProfile[]>(uri_);
	}
	
			//********************************************************************
	// assigns a BrandSafetyPolicy on a TargetingProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBrandSafetyPolicy( targetingProfileId, _brandSafetyPolicyId ): Observable<any> {

		// get the TargetingProfile from storage
		this.loadHelper( targetingProfileId );

	// get the BrandSafetyPolicy from storage
	var tmp 	= new BrandSafetyPolicyService(this.http).getBrandSafetyPolicy(_brandSafetyPolicyId);

	// assign the BrandSafetyPolicy
	this.targetingProfile.brandSafetyPolicy = tmp;

	// save the TargetingProfile
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a BrandSafetyPolicy on a TargetingProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBrandSafetyPolicy( targetingProfileId ): Observable<any> {

		// get the TargetingProfile from storage
		this.loadHelper( targetingProfileId );

	// assign BrandSafetyPolicy to null
	this.targetingProfile.brandSafetyPolicy = null;

	// save the TargetingProfile
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more audienceSegmentsIds as a AudienceSegments
	// to a TargetingProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAudienceSegments( targetingProfileId, audienceSegmentsIds ): Observable<any> {

		// get the TargetingProfile
		this.loadHelper( targetingProfileId );

	// split on a comma with no spaces
	var idList = audienceSegmentsIds.split(',')

	// iterate over array of audienceSegments ids
	idList.forEach(function (id) {
		// read the AudienceSegment
		var audienceSegment = new AudienceSegmentService(this.http).getAudienceSegment(id);
		// add the AudienceSegment if not already assigned
		if ( this.targetingProfile.audienceSegments.indexOf(audienceSegment) == -1 )
		this.targetingProfile.audienceSegments.push(audienceSegment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more audienceSegmentsIds as a AudienceSegments
	// from a TargetingProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAudienceSegments( targetingProfileId, audienceSegmentsIds ): Observable<any> {

		// get the TargetingProfile
		this.loadHelper( targetingProfileId );


	// split on a comma with no spaces
	var idList 					= audienceSegmentsIds.split(',');
	var audienceSegments 	= this.targetingProfile.audienceSegments;

	if ( audienceSegments != null && audienceSegmentsIds != null ) {

		// iterate over array of audienceSegments ids
		audienceSegments.forEach(function (obj) {
			if ( audienceSegmentsIds.indexOf(obj._id) > -1 ) {
				// remove the AudienceSegment
				this.targetingProfile.audienceSegments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more geoRegionsIds as a GeoRegions
	// to a TargetingProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addGeoRegions( targetingProfileId, geoRegionsIds ): Observable<any> {

		// get the TargetingProfile
		this.loadHelper( targetingProfileId );

	// split on a comma with no spaces
	var idList = geoRegionsIds.split(',')

	// iterate over array of geoRegions ids
	idList.forEach(function (id) {
		// read the GeoRegion
		var geoRegion = new GeoRegionService(this.http).getGeoRegion(id);
		// add the GeoRegion if not already assigned
		if ( this.targetingProfile.geoRegions.indexOf(geoRegion) == -1 )
		this.targetingProfile.geoRegions.push(geoRegion);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more geoRegionsIds as a GeoRegions
	// from a TargetingProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeGeoRegions( targetingProfileId, geoRegionsIds ): Observable<any> {

		// get the TargetingProfile
		this.loadHelper( targetingProfileId );


	// split on a comma with no spaces
	var idList 					= geoRegionsIds.split(',');
	var geoRegions 	= this.targetingProfile.geoRegions;

	if ( geoRegions != null && geoRegionsIds != null ) {

		// iterate over array of geoRegions ids
		geoRegions.forEach(function (obj) {
			if ( geoRegionsIds.indexOf(obj._id) > -1 ) {
				// remove the GeoRegion
				this.targetingProfile.geoRegions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more contentCategoriesIds as a ContentCategories
	// to a TargetingProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addContentCategories( targetingProfileId, contentCategoriesIds ): Observable<any> {

		// get the TargetingProfile
		this.loadHelper( targetingProfileId );

	// split on a comma with no spaces
	var idList = contentCategoriesIds.split(',')

	// iterate over array of contentCategories ids
	idList.forEach(function (id) {
		// read the ContentCategory
		var contentCategory = new ContentCategoryService(this.http).getContentCategory(id);
		// add the ContentCategory if not already assigned
		if ( this.targetingProfile.contentCategories.indexOf(contentCategory) == -1 )
		this.targetingProfile.contentCategories.push(contentCategory);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more contentCategoriesIds as a ContentCategories
	// from a TargetingProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeContentCategories( targetingProfileId, contentCategoriesIds ): Observable<any> {

		// get the TargetingProfile
		this.loadHelper( targetingProfileId );


	// split on a comma with no spaces
	var idList 					= contentCategoriesIds.split(',');
	var contentCategories 	= this.targetingProfile.contentCategories;

	if ( contentCategories != null && contentCategoriesIds != null ) {

		// iterate over array of contentCategories ids
		contentCategories.forEach(function (obj) {
			if ( contentCategoriesIds.indexOf(obj._id) > -1 ) {
				// remove the ContentCategory
				this.targetingProfile.contentCategories.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more deviceCriteriaIds as a DeviceCriteria
	// to a TargetingProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDeviceCriteria( targetingProfileId, deviceCriteriaIds ): Observable<any> {

		// get the TargetingProfile
		this.loadHelper( targetingProfileId );

	// split on a comma with no spaces
	var idList = deviceCriteriaIds.split(',')

	// iterate over array of deviceCriteria ids
	idList.forEach(function (id) {
		// read the DeviceCriterion
		var deviceCriterion = new DeviceCriterionService(this.http).getDeviceCriterion(id);
		// add the DeviceCriterion if not already assigned
		if ( this.targetingProfile.deviceCriteria.indexOf(deviceCriterion) == -1 )
		this.targetingProfile.deviceCriteria.push(deviceCriterion);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more deviceCriteriaIds as a DeviceCriteria
	// from a TargetingProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDeviceCriteria( targetingProfileId, deviceCriteriaIds ): Observable<any> {

		// get the TargetingProfile
		this.loadHelper( targetingProfileId );


	// split on a comma with no spaces
	var idList 					= deviceCriteriaIds.split(',');
	var deviceCriteria 	= this.targetingProfile.deviceCriteria;

	if ( deviceCriteria != null && deviceCriteriaIds != null ) {

		// iterate over array of deviceCriteria ids
		deviceCriteria.forEach(function (obj) {
			if ( deviceCriteriaIds.indexOf(obj._id) > -1 ) {
				// remove the DeviceCriterion
				this.targetingProfile.deviceCriteria.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a TargetingProfile
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/TargetingProfile/update/' + this.targetingProfile;

	return  this.http.post(uri_, this.targetingProfile );
}

	//********************************************************************
	// loadHelper - internal helper to load a TargetingProfile
	//********************************************************************	
	loadHelper( id ) {
		this.getTargetingProfile(id)
			.subscribe((res : TargetingProfile) => {
				this.targetingProfile = res;
			});
	}
}