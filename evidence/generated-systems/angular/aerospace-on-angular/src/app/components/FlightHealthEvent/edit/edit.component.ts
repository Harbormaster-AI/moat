import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { FlightHealthEventService } from '../../../services/FlightHealthEvent.service';
import { SubBaseComponent } from '../../FlightHealthEvent/sub.base.component';


@Component({
    selector: 'app-edit-flightHealthEvent',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditFlightHealthEventComponent extends SubBaseComponent implements OnInit {

    title = 'Edit FlightHealthEvent';

    flightHealthEventForm: FormGroup;
    flightHealthEvent: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: FlightHealthEventService,
        private fb: FormBuilder
) {
        super(http);
        this.flightHealthEventForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  eventCode: ['', Validators.required],
      ConnectedAircraft: ['', ],
      Severity: ['', ]
        });
    }

    
    updateFlightHealthEvent(eventCode, ConnectedAircraft, Severity): void {
        this.route.params.subscribe((params) => {

                        this.service.updateFlightHealthEvent(eventCode, ConnectedAircraft, Severity, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexFlightHealthEvent']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getFlightHealthEvent(params['id']).subscribe(res => {
                this.flightHealthEvent = res;
            });
        });
    }
}