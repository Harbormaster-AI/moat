import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LocationService } from '../../../services/Location.service';
import { Location } from '../../../models/Location';
import { SubBaseComponent } from '../../Location/sub.base.component';

@Component({
    selector: 'app-create-location',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLocationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Location';

    locationForm: FormGroup;
    location: Location;

    constructor( http: HttpClient,
        private locationService: LocationService,
        private fb: FormBuilder,
        private router: Router
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

    
    addLocation(locationCode, description, Warehouse, InventoryItems, LocationType): void {
        this.locationService
        .addLocation(locationCode, description, Warehouse, InventoryItems, LocationType)
            .subscribe(() => {
                this.router.navigate(['/indexLocation']);
            });
    }

    ngOnInit(): void {
    }
}