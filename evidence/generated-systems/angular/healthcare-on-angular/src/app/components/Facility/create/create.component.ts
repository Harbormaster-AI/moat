import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { FacilityService } from '../../../services/Facility.service';
import { Facility } from '../../../models/Facility';
import { SubBaseComponent } from '../../Facility/sub.base.component';

@Component({
    selector: 'app-create-facility',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateFacilityComponent extends SubBaseComponent implements OnInit {

    title = 'Add Facility';

    facilityForm: FormGroup;
    facility: Facility;

    constructor( http: HttpClient,
        private facilityService: FacilityService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.facilityForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      facilityCode: ['', Validators.required],
      address: ['', Validators.required],
      HealthSystem: ['', ],
      Departments: ['', ],
      CareTeams: ['', ],
      Laboratories: ['', ],
      ImagingCenters: ['', ],
      Pharmacies: ['', ],
      InventoryItems: ['', ],
      FacilityType: ['', ]
        });
    }

    
    addFacility(name, facilityCode, address, HealthSystem, Departments, CareTeams, Laboratories, ImagingCenters, Pharmacies, InventoryItems, FacilityType): void {
        this.facilityService
        .addFacility(name, facilityCode, address, HealthSystem, Departments, CareTeams, Laboratories, ImagingCenters, Pharmacies, InventoryItems, FacilityType)
            .subscribe(() => {
                this.router.navigate(['/indexFacility']);
            });
    }

    ngOnInit(): void {
    }
}