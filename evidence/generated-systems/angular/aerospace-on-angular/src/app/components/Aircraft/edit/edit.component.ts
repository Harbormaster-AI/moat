import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AircraftService } from '../../../services/Aircraft.service';
import { SubBaseComponent } from '../../Aircraft/sub.base.component';


@Component({
    selector: 'app-edit-aircraft',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAircraftComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Aircraft';

    aircraftForm: FormGroup;
    aircraft: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AircraftService,
        private fb: FormBuilder
) {
        super(http);
        this.aircraftForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  msn: ['', Validators.required],
      deliveryDate: ['', Validators.required],
      Variant: ['', ],
      Operator: ['', ],
      Registration: ['', ],
      Warranty: ['', ],
      MaintenanceRecords: ['', ],
      ConnectedAircraft: ['', ],
      CabinLayout: ['', ]
        });
    }

    
    updateAircraft(msn, deliveryDate, Variant, Operator, Registration, Warranty, MaintenanceRecords, ConnectedAircraft, CabinLayout): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAircraft(msn, deliveryDate, Variant, Operator, Registration, Warranty, MaintenanceRecords, ConnectedAircraft, CabinLayout, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAircraft']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAircraft(params['id']).subscribe(res => {
                this.aircraft = res;
            });
        });
    }
}