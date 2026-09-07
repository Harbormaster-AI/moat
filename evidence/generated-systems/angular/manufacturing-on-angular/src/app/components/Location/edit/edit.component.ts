import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LocationService } from '../../../services/Location.service';
import { SubBaseComponent } from '../../Location/sub.base.component';


@Component({
    selector: 'app-edit-location',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLocationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Location';

    locationForm: FormGroup;
    location: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LocationService,
        private fb: FormBuilder
) {
        super(http);
        this.locationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  locationCode: ['', Validators.required],
      description: ['', Validators.required],
      Warehouse: ['', ],
      InventoryItems: ['', ],
      LocationType: ['', ]
        });
    }

    
    updateLocation(locationCode, description, Warehouse, InventoryItems, LocationType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLocation(locationCode, description, Warehouse, InventoryItems, LocationType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLocation']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLocation(params['id']).subscribe(res => {
                this.location = res;
            });
        });
    }
}