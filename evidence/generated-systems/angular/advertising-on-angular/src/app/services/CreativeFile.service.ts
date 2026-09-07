import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CreativeFile} from '../models/CreativeFile';
import {CreativeAssetService} from '../services/CreativeAsset.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CreativeFileService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	creativeFile : CreativeFile;

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
	// add a CreativeFile
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCreativeFile(uri, fileSizeKB, mimeType, checksum, CreativeAsset) : Observable<any> {
		const uri_ = this.apiUrl + '/CreativeFile/create';
		const obj = {
			      		uri: uri,
      		fileSizeKB: fileSizeKB,
      		mimeType: mimeType,
      		checksum: checksum,
			CreativeAsset: CreativeAsset != null && CreativeAsset.length > 0 ? CreativeAsset : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CreativeFile
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCreativeFile(uri, fileSizeKB, mimeType, checksum, CreativeAsset, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CreativeFile/update/' + id;
		const obj = {
				      		uri: uri,
      		fileSizeKB: fileSizeKB,
      		mimeType: mimeType,
      		checksum: checksum,
			CreativeAsset: CreativeAsset != null && CreativeAsset.length > 0 ? CreativeAsset : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CreativeFile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCreativeFile(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CreativeFile/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CreativeFile
	// returns the results untouched as an Observable CreativeFile
	// CreativeFile model
	// delegates via URI
	//********************************************************************
	getCreativeFile(id) : Observable<CreativeFile> {
		const uri_ = this.apiUrl + '/CreativeFile/load/' + id;

		return this.http.get<CreativeFile>(uri_);
	}
	
	//********************************************************************
	// gets all CreativeFile
	// returns the results untouched as JSON representation of an
	// Observable array of CreativeFile models
	// delegates via URI
	//********************************************************************
	getCreativeFiles() : Observable<CreativeFile[]> {
		const uri_ = this.apiUrl + '/CreativeFile/';

		return this
			.http.get<CreativeFile[]>(uri_);
	}
	
			//********************************************************************
	// assigns a CreativeAsset on a CreativeFile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCreativeAsset( creativeFileId, _creativeAssetId ): Observable<any> {

		// get the CreativeFile from storage
		this.loadHelper( creativeFileId );

	// get the CreativeAsset from storage
	var tmp 	= new CreativeAssetService(this.http).getCreativeAsset(_creativeAssetId);

	// assign the CreativeAsset
	this.creativeFile.creativeAsset = tmp;

	// save the CreativeFile
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CreativeAsset on a CreativeFile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCreativeAsset( creativeFileId ): Observable<any> {

		// get the CreativeFile from storage
		this.loadHelper( creativeFileId );

	// assign CreativeAsset to null
	this.creativeFile.creativeAsset = null;

	// save the CreativeFile
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a CreativeFile
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CreativeFile/update/' + this.creativeFile;

	return  this.http.post(uri_, this.creativeFile );
}

	//********************************************************************
	// loadHelper - internal helper to load a CreativeFile
	//********************************************************************	
	loadHelper( id ) {
		this.getCreativeFile(id)
			.subscribe((res : CreativeFile) => {
				this.creativeFile = res;
			});
	}
}