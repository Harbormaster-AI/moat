import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AircraftService } from '../../../services/Aircraft.service';
import { Aircraft } from '../../../models/Aircraft';
import { SubBaseComponent } from '../../Aircraft/sub.base.component';

@Component({
    selector: 'app-create-aircraft',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAircraftComponent extends SubBaseComponent implements OnInit {

    title = 'Add Aircraft';

    aircraftForm: FormGroup;
    aircraft: Aircraft;

    constructor( http: HttpClient,
        private aircraftService: AircraftService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAircraft(msn, deliveryDate, Variant, Operator, Registration, Warranty, MaintenanceRecords, ConnectedAircraft, CabinLayout): void {
        this.aircraftService
        .addAircraft(msn, deliveryDate, Variant, Operator, Registration, Warranty, MaintenanceRecords, ConnectedAircraft, CabinLayout)
            .subscribe(() => {
                this.router.navigate(['/indexAircraft']);
            });
    }

    ngOnInit(): void {
    }
}