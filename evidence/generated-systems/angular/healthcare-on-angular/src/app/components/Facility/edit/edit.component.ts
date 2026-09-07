import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { FacilityService } from '../../../services/Facility.service';
import { SubBaseComponent } from '../../Facility/sub.base.component';


@Component({
    selector: 'app-edit-facility',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditFacilityComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Facility';

    facilityForm: FormGroup;
    facility: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: FacilityService,
        private fb: FormBuilder
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

    
    updateFacility(name, facilityCode, address, HealthSystem, Departments, CareTeams, Laboratories, ImagingCenters, Pharmacies, InventoryItems, FacilityType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateFacility(name, facilityCode, address, HealthSystem, Departments, CareTeams, Laboratories, ImagingCenters, Pharmacies, InventoryItems, FacilityType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexFacility']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getFacility(params['id']).subscribe(res => {
                this.facility = res;
            });
        });
    }
}